package reward

import (
	"context"
	"time"

	payloadComp "example.com/creditcard/components/payload"
	"github.com/sirupsen/logrus"

	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"

	payloadM "example.com/creditcard/models/payload"
)

type impl struct {
	rewardType        rewardM.RewardType
	rewardResp        *rewardM.RewardResp
	payloadComponents []*payloadComp.Component
}

func New(
	rewardType rewardM.RewardType,
	rewardResp *rewardM.RewardResp,
	payloadComponents []*payloadComp.Component,
) Component {

	return &impl{
		rewardType:        rewardType,
		rewardResp:        rewardResp,
		payloadComponents: payloadComponents,
	}
}

const DATE_FORMAT = "2006/01/02"

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*rewardM.RewardEventResp, error) {

	rewardEventResp := &rewardM.RewardEventResp{

		ID: im.rewardResp.ID,

		Order: im.rewardResp.Order,

		PayloadOperator: im.rewardResp.PayloadOperator,
	}

	startDate, err := time.Parse(DATE_FORMAT, im.rewardResp.StartDate)
	if err != nil {
		logrus.Error("Error parseInt startDate")
		return nil, err
	}

	endDate, err := time.Parse(DATE_FORMAT, im.rewardResp.EndDate)
	if err != nil {
		logrus.Error("Error parseInt endDate")
		return nil, err
	}

	if !(startDate.Unix() <= e.EffectiveTime && e.EffectiveTime <= endDate.Unix()) {
		logrus.Error("Outdated")

		rewardEventResp.RewardEventJudgeType = rewardM.NONE
		// TODO FeedReturn initialized
		return rewardEventResp, nil
	}

	payloadEventResps := []*payloadM.PayloadEventResp{}

	for _, p := range im.payloadComponents {
		payloadEventResp, err := (*p).Satisfy(ctx, e)
		if err != nil {
			return nil, err
		}

		payloadEventResps = append(payloadEventResps, payloadEventResp)
	}

	rewardEventResp.PayloadEventResps = payloadEventResps

	switch im.rewardType {
	case rewardM.CASH_TWD:
		cashReturn, err := im.calculateCashReturn(ctx, im.rewardResp.PayloadOperator, payloadEventResps)

		if err != nil {
			return nil, err
		}

		if cashReturn.ActualUseCash == cashReturn.CurrentCash {
			rewardEventResp.RewardEventJudgeType = rewardM.ALL
		} else if cashReturn.ActualUseCash == 0 {
			rewardEventResp.RewardEventJudgeType = rewardM.NONE
		} else {
			rewardEventResp.RewardEventJudgeType = rewardM.SOME
		}

		rewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			CashReturn: cashReturn,
		}
		break

	case rewardM.POINT:
		break
	default:

	}

	return rewardEventResp, nil
}

func (im *impl) calculateCashReturn(ctx context.Context, operator rewardM.PayloadOperator, payloadEventResps []*payloadM.PayloadEventResp) (*feedbackM.CashReturn, error) {

	cashReturn := &feedbackM.CashReturn{}

	switch operator {
	case rewardM.ADD:

		var totalCash float64 = 0.0
		var currentCash int64 = 0

		var isCashbackGet bool = false
		var cashbackBonus float64 = 0.0

		var actualUseCash int64 = 0
		var actualCashReturn float64 = 0.0

		for _, p := range payloadEventResps {

			totalCash = p.FeedReturn.CashReturn.TotalCash
			currentCash = p.FeedReturn.CashReturn.CurrentCash

			if p.PayloadEventJudgeType == payloadM.ALL || p.PayloadEventJudgeType == payloadM.SOME {
				isCashbackGet = true

				if actualUseCash < p.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = p.FeedReturn.CashReturn.ActualUseCash
				}

				actualCashReturn += p.FeedReturn.CashReturn.ActualCashReturn
				cashbackBonus += p.FeedReturn.CashReturn.CashbackBonus
			}

		}

		cashReturn.IsCashbackGet = isCashbackGet
		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CurrentCash = currentCash
		cashReturn.TotalCash = totalCash
		cashReturn.CashbackBonus = cashbackBonus

		break
	case rewardM.MAXONE:

		var maxBonus float64 = 0.0

		var atLeastOnePass = false

		finalPayload := &payloadM.PayloadEventResp{}

		for _, p := range payloadEventResps {

			if p.PayloadEventJudgeType == payloadM.ALL || p.PayloadEventJudgeType == payloadM.SOME {
				if maxBonus < p.FeedReturn.CashReturn.CashbackBonus {
					finalPayload = p
					maxBonus = p.FeedReturn.CashReturn.CashbackBonus
				}
				atLeastOnePass = true
			}
		}

		if atLeastOnePass {
			cashReturn.IsCashbackGet = finalPayload.FeedReturn.CashReturn.IsCashbackGet
			cashReturn.ActualCashReturn = finalPayload.FeedReturn.CashReturn.ActualCashReturn
			cashReturn.ActualUseCash = finalPayload.FeedReturn.CashReturn.ActualUseCash
			cashReturn.CurrentCash = finalPayload.FeedReturn.CashReturn.CurrentCash
			cashReturn.TotalCash = finalPayload.FeedReturn.CashReturn.TotalCash
			cashReturn.CashbackBonus = maxBonus
		}

		break
	default:

	}

	return cashReturn, nil
}
