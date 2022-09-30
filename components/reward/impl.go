package reward

import (
	"context"
	"errors"

	payloadComp "example.com/creditcard/components/payload"
	"github.com/sirupsen/logrus"

	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"

	payloadM "example.com/creditcard/models/payload"
)

type impl struct {
	rewardType        rewardM.RewardType
	reward            *rewardM.Reward
	payloadComponents []*payloadComp.Component
}

func New(
	rewardType rewardM.RewardType,
	reward *rewardM.Reward,
	payloadComponents []*payloadComp.Component,
) Component {

	return &impl{
		rewardType:        rewardType,
		reward:            reward,
		payloadComponents: payloadComponents,
	}
}

const DATE_FORMAT = "2006/01/02"

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*rewardM.RewardEventResp, error) {

	rewardEventResp := &rewardM.RewardEventResp{

		ID:           im.reward.ID,
		CardRewardID: im.reward.CardRewardID,

		Order: im.reward.Order,

		PayloadOperator: im.reward.PayloadOperator,
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
		cashReturn, err := im.calculateCashReturn(ctx, im.reward.PayloadOperator, payloadEventResps)

		if err != nil {
			return nil, err
		}

		cashReturn.CurrentCash = int64(e.Cash)
		cashReturn.TotalCash = e.Cash

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

	case rewardM.LINE_POINT:

		pointReturn, err := im.calculatePointReturn(ctx, im.reward.PayloadOperator, payloadEventResps)

		if err != nil {
			return nil, err
		}

		pointReturn.CurrentCash = int64(e.Cash)
		pointReturn.TotalCash = e.Cash

		if pointReturn.ActualUseCash == pointReturn.CurrentCash {
			rewardEventResp.RewardEventJudgeType = rewardM.ALL
		} else if pointReturn.ActualUseCash == 0 {
			rewardEventResp.RewardEventJudgeType = rewardM.NONE
		} else {
			rewardEventResp.RewardEventJudgeType = rewardM.SOME
		}

		rewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			PointReturn: pointReturn,
		}
		break

	case rewardM.KUO_BROTHERS_POINT:

		pointReturn, err := im.calculatePointReturn(ctx, im.reward.PayloadOperator, payloadEventResps)

		if err != nil {
			return nil, err
		}

		pointReturn.CurrentCash = int64(e.Cash)
		pointReturn.TotalCash = e.Cash

		if pointReturn.ActualUseCash == pointReturn.CurrentCash {
			rewardEventResp.RewardEventJudgeType = rewardM.ALL
		} else if pointReturn.ActualUseCash == 0 {
			rewardEventResp.RewardEventJudgeType = rewardM.NONE
		} else {
			rewardEventResp.RewardEventJudgeType = rewardM.SOME
		}

		rewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			PointReturn: pointReturn,
		}
		break

	case rewardM.WOWPRIME_POINT:

		pointReturn, err := im.calculatePointReturn(ctx, im.reward.PayloadOperator, payloadEventResps)

		if err != nil {
			return nil, err
		}

		pointReturn.CurrentCash = int64(e.Cash)
		pointReturn.TotalCash = e.Cash

		if pointReturn.ActualUseCash == pointReturn.CurrentCash {
			rewardEventResp.RewardEventJudgeType = rewardM.ALL
		} else if pointReturn.ActualUseCash == 0 {
			rewardEventResp.RewardEventJudgeType = rewardM.NONE
		} else {
			rewardEventResp.RewardEventJudgeType = rewardM.SOME
		}

		rewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			PointReturn: pointReturn,
		}

		break

	case rewardM.OPEN_POINT:
		pointReturn, err := im.calculatePointReturn(ctx, im.reward.PayloadOperator, payloadEventResps)

		if err != nil {
			return nil, err
		}

		pointReturn.CurrentCash = int64(e.Cash)
		pointReturn.TotalCash = e.Cash

		if pointReturn.ActualUseCash == pointReturn.CurrentCash {
			rewardEventResp.RewardEventJudgeType = rewardM.ALL
		} else if pointReturn.ActualUseCash == 0 {
			rewardEventResp.RewardEventJudgeType = rewardM.NONE
		} else {
			rewardEventResp.RewardEventJudgeType = rewardM.SOME
		}

		rewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			PointReturn: pointReturn,
		}

		break
	default:
		logrus.Error("not found reward type in reward component")
		return nil, errors.New("not found reward type in reward component")
	}

	return rewardEventResp, nil
}

func (im *impl) calculatePointReturn(ctx context.Context, operator rewardM.PayloadOperator, payloadEventResps []*payloadM.PayloadEventResp) (*feedbackM.PointReturn, error) {

	pointReturn := &feedbackM.PointReturn{}

	switch operator {
	case rewardM.ADD:

		var totalCash float64 = 0.0
		var currentCash int64 = 0

		var isPointbackGet bool = false
		var pointbackBonus float64 = 0.0

		var actualUseCash int64 = 0
		var actualPointBack float64 = 0.0

		for _, p := range payloadEventResps {

			totalCash = p.FeedReturn.PointReturn.TotalCash
			currentCash = p.FeedReturn.PointReturn.CurrentCash

			if p.PayloadEventJudgeType == payloadM.ALL || p.PayloadEventJudgeType == payloadM.SOME {
				isPointbackGet = true

				if actualUseCash < p.FeedReturn.PointReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = p.FeedReturn.PointReturn.ActualUseCash
				}

				actualPointBack += p.FeedReturn.PointReturn.ActualPointBack
				pointbackBonus += p.FeedReturn.PointReturn.PointbackBonus
			}

		}

		pointReturn.IsPointbackGet = isPointbackGet
		pointReturn.ActualPointBack = actualPointBack
		pointReturn.ActualUseCash = actualUseCash
		pointReturn.CurrentCash = currentCash
		pointReturn.TotalCash = totalCash
		pointReturn.PointbackBonus = pointbackBonus

		break
	case rewardM.MAXONE:

		var maxBonus float64 = 0.0

		var atLeastOnePass = false

		finalPayload := &payloadM.PayloadEventResp{}

		for _, p := range payloadEventResps {

			if p.PayloadEventJudgeType == payloadM.ALL || p.PayloadEventJudgeType == payloadM.SOME {
				if maxBonus < p.FeedReturn.PointReturn.PointbackBonus {
					finalPayload = p
					maxBonus = p.FeedReturn.PointReturn.PointbackBonus
				}
				atLeastOnePass = true
			}
		}

		if atLeastOnePass {
			pointReturn.IsPointbackGet = finalPayload.FeedReturn.PointReturn.IsPointbackGet
			pointReturn.ActualPointBack = finalPayload.FeedReturn.PointReturn.ActualPointBack
			pointReturn.ActualUseCash = finalPayload.FeedReturn.PointReturn.ActualUseCash
			pointReturn.CurrentCash = finalPayload.FeedReturn.PointReturn.CurrentCash
			pointReturn.TotalCash = finalPayload.FeedReturn.PointReturn.TotalCash
			pointReturn.PointbackBonus = maxBonus
		}

		break
	default:

	}

	return pointReturn, nil
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
