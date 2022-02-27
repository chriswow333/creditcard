package reward

import (
	"context"
	"time"

	payloadComp "example.com/creditcard/components/payload"

	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
)

type impl struct {
	reward            *rewardM.Reward
	payloadComponents []*payloadComp.Component
}

func New(
	reward *rewardM.Reward,
	payloadComponents []*payloadComp.Component,
) Component {

	return &impl{
		reward:            reward,
		payloadComponents: payloadComponents,
	}
}

const DATE_FORMAT = "2006/01/02"

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.RewardResp, error) {

	rewardResp := &eventM.RewardResp{

		Order:    im.reward.Order,
		Title:    im.reward.Title,
		SubTitle: im.reward.SubTitle,

		StartDate:  time.Unix(im.reward.StartDate, 0).Format(DATE_FORMAT),
		EndDate:    time.Unix(im.reward.EndDate, 0).Format(DATE_FORMAT),
		UpdateDate: time.Unix(im.reward.UpdateDate, 0).Format(DATE_FORMAT),

		PayloadOperator: im.reward.PayloadOperator,
	}

	if !(im.reward.StartDate <= e.EffictiveTime && e.EffictiveTime <= im.reward.EndDate) {
		return rewardResp, nil
	}

	payloadResps := []*eventM.PayloadResp{}

	for _, p := range im.payloadComponents {
		payloadResp, err := (*p).Satisfy(ctx, e)
		if err != nil {
			return nil, err
		}

		payloadResps = append(payloadResps, payloadResp)
	}

	rewardResp.PayloadResps = payloadResps
	cashReturn, err := im.calculateCashReturn(ctx, im.reward.PayloadOperator, payloadResps)
	if err != nil {
		return nil, err
	}

	rewardResp.FeedReturn = &feedbackM.FeedReturn{
		CashReturn: cashReturn,
	}

	return rewardResp, nil
}

func (im *impl) calculateCashReturn(ctx context.Context, operator rewardM.PayloadOperator, payloadResps []*eventM.PayloadResp) (*feedbackM.CashReturn, error) {

	cashReturn := &feedbackM.CashReturn{}

	switch operator {
	case rewardM.AddPayloadOperator:

		var totalCash float64 = 0.0
		var currentCash int64 = 0

		var isCashbackGet bool = false
		var cashbackBonus float64 = 0.0

		var actualUseCash int64 = 0
		var actualCashReturn float64 = 0.0

		for _, f := range payloadResps {

			totalCash = f.FeedReturn.CashReturn.TotalCash
			currentCash = f.FeedReturn.CashReturn.CurrentCash

			if f.Pass {
				isCashbackGet = true

				if actualUseCash < f.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = f.FeedReturn.CashReturn.ActualUseCash
				}

				actualCashReturn += f.FeedReturn.CashReturn.ActualCashReturn
				cashbackBonus += f.FeedReturn.CashReturn.CashbackBonus
			}
		}

		cashReturn.IsCashbackGet = isCashbackGet
		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CurrentCash = currentCash
		cashReturn.TotalCash = totalCash
		cashReturn.CashbackBonus = cashbackBonus

		break
	case rewardM.XORHighPayloadOperator:
		var maxBonus float64 = 0.0
		finalPayload := &eventM.PayloadResp{}
		for _, f := range payloadResps {
			if f.Pass {
				if maxBonus < f.Feedback.Cashback.Bonus {
					finalPayload = f
					maxBonus = f.Feedback.Cashback.Bonus
				}
			}
		}
		cashReturn.IsCashbackGet = finalPayload.FeedReturn.CashReturn.IsCashbackGet
		cashReturn.ActualCashReturn = finalPayload.FeedReturn.CashReturn.ActualCashReturn
		cashReturn.ActualUseCash = finalPayload.FeedReturn.CashReturn.ActualUseCash
		cashReturn.CurrentCash = finalPayload.FeedReturn.CashReturn.CurrentCash
		cashReturn.TotalCash = finalPayload.FeedReturn.CashReturn.TotalCash
		cashReturn.CashbackBonus = maxBonus
		break
	default:

	}

	return cashReturn, nil
}
