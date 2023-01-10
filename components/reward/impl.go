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
	case rewardM.CASH:
		cashReturn, err := im.calculateCashReturn(ctx, e, im.reward.PayloadOperator, payloadEventResps)

		if err != nil {
			return nil, err
		}

		feedReturn := &feedbackM.FeedReturn{
			CashReturn: cashReturn,
		}

		rewardEventResp.FeedReturn = feedReturn

		break

	case rewardM.POINT:

		pointReturn, err := im.calculatePointReturn(ctx, e, im.reward.PayloadOperator, payloadEventResps)

		if err != nil {
			return nil, err
		}

		feedReturn := &feedbackM.FeedReturn{
			PointReturn: pointReturn,
		}

		rewardEventResp.FeedReturn = feedReturn

		break

	// case rewardM.RED:
	// 	redReturn, err := im.calculateRedPointReturn(ctx, im.reward.PayloadOperator, payloadEventResps)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	redReturn.CurrentCash = int64(e.Cash)
	// 	redReturn.TotalCash = e.Cash
	// 	feedReturn := &feedbackM.FeedReturn{}

	// 	if redReturn.ActualUseCash == redReturn.CurrentCash {
	// 		rewardEventResp.RewardEventJudgeType = rewardM.ALL
	// 		feedReturn.FeedReturnStatus = feedbackM.ALL
	// 	} else if redReturn.ActualUseCash == 0 {
	// 		rewardEventResp.RewardEventJudgeType = rewardM.NONE
	// 		feedReturn.FeedReturnStatus = feedbackM.NONE
	// 	} else {
	// 		rewardEventResp.RewardEventJudgeType = rewardM.SOME
	// 		feedReturn.FeedReturnStatus = feedbackM.SOME
	// 	}
	// 	rewardEventResp.FeedReturn = feedReturn

	// 	break
	default:
		logrus.Error("not found reward type in reward component")
		return nil, errors.New("not found reward type in reward component")
	}

	return rewardEventResp, nil
}

// func (im *impl) calculateRedPointReturn(ctx context.Context, operator rewardM.PayloadOperator, payloadEventResps []*payloadM.PayloadEventResp) (*feedbackM.RedReturn, error) {

// 	redReturn := &feedbackM.RedReturn{}

// 	switch operator {
// 	case rewardM.ADD:

// 		var totalCash float64 = 0.0
// 		var currentCash int64 = 0

// 		var isRedbackGet bool = false
// 		var redbackTimes int64 = 0.0

// 		var actualUseCash int64 = 0
// 		var actualRedBack float64 = 0.0

// 		for _, p := range payloadEventResps {

// 			totalCash = p.FeedReturn.RedReturn.TotalCash
// 			currentCash = p.FeedReturn.RedReturn.CurrentCash

// 			if p.PayloadEventJudgeType == payloadM.ALL || p.PayloadEventJudgeType == payloadM.SOME {
// 				isRedbackGet = true

// 				if actualUseCash < p.FeedReturn.RedReturn.ActualUseCash {
// 					// get max actual use cash
// 					actualUseCash = p.FeedReturn.RedReturn.ActualUseCash
// 				}

// 				actualRedBack += p.FeedReturn.RedReturn.ActualRedback
// 				redbackTimes += p.FeedReturn.RedReturn.RedbackTimes
// 			}

// 		}

// 		redReturn.IsRedGet = isRedbackGet
// 		redReturn.ActualRedback = actualRedBack
// 		redReturn.ActualUseCash = actualUseCash
// 		redReturn.CurrentCash = currentCash
// 		redReturn.TotalCash = totalCash
// 		redReturn.RedbackTimes = redbackTimes

// 		break
// 	case rewardM.MAXONE:

// 		var maxBonus int64 = 0

// 		var atLeastOnePass = false

// 		finalPayload := &payloadM.PayloadEventResp{}
// 		for _, p := range payloadEventResps {
// 			if p.PayloadEventJudgeType == payloadM.ALL || p.PayloadEventJudgeType == payloadM.SOME {
// 				if maxBonus < p.FeedReturn.RedReturn.RedbackTimes {
// 					finalPayload = p
// 					maxBonus = p.FeedReturn.RedReturn.RedbackTimes
// 				}
// 				atLeastOnePass = true
// 			}
// 		}

// 		if atLeastOnePass {
// 			redReturn.IsRedGet = finalPayload.FeedReturn.RedReturn.IsRedGet
// 			redReturn.ActualRedback = finalPayload.FeedReturn.RedReturn.ActualRedback
// 			redReturn.ActualUseCash = finalPayload.FeedReturn.RedReturn.ActualUseCash
// 			redReturn.CurrentCash = finalPayload.FeedReturn.RedReturn.CurrentCash
// 			redReturn.TotalCash = finalPayload.FeedReturn.RedReturn.TotalCash
// 			redReturn.RedbackTimes = maxBonus
// 		}

// 		break
// 	default:

// 	}

// 	return redReturn, nil
// }

func (im *impl) calculatePointReturn(ctx context.Context, e *eventM.Event, operator rewardM.PayloadOperator, payloadEventResps []*payloadM.PayloadEventResp) (*feedbackM.PointReturn, error) {

	pointReturn := &feedbackM.PointReturn{
		Cash: e.Cash,
	}

	switch operator {
	case rewardM.ADD:

		var pointReturnBonus float64 = 0.0
		var actualUseCash int64 = 0
		var actualPointReturn float64 = 0.0
		var pointReturnStatus feedbackM.PointReturnStatus = feedbackM.NONE_RETURN_POINT
		hasReturn := false

		for _, p := range payloadEventResps {

			switch p.FeedReturn.PointReturn.PointReturnStatus {
			case feedbackM.ALL_RETURN_POINT:
				if actualUseCash < p.FeedReturn.PointReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = p.FeedReturn.PointReturn.ActualUseCash
				}

				actualPointReturn += p.FeedReturn.PointReturn.ActualPointReturn
				pointReturnBonus += p.FeedReturn.PointReturn.PointReturnBonus

				if !hasReturn {
					pointReturnStatus = feedbackM.ALL_RETURN_POINT
				} else {
					pointReturnStatus = feedbackM.SOME_RETURN_POINT
				}

				hasReturn = true
				break
			case feedbackM.SOME_RETURN_POINT:
				hasReturn = true
				if actualUseCash < p.FeedReturn.PointReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = p.FeedReturn.PointReturn.ActualUseCash
				}

				if actualUseCash < p.FeedReturn.PointReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = p.FeedReturn.PointReturn.ActualUseCash
				}
				actualPointReturn += p.FeedReturn.PointReturn.ActualPointReturn
				pointReturnBonus += p.FeedReturn.PointReturn.PointReturnBonus

				pointReturnStatus = feedbackM.SOME_RETURN_CASH

				break
			case feedbackM.NONE_RETURN_POINT:
				hasReturn = true
				if pointReturnStatus == feedbackM.ALL_RETURN_CASH {
					pointReturnStatus = feedbackM.SOME_RETURN_CASH
				}
				break
			}
		}

		pointReturn.ActualPointReturn = actualPointReturn
		pointReturn.ActualUseCash = actualUseCash
		pointReturn.PointReturnBonus = pointReturnBonus
		pointReturn.PointReturnStatus = pointReturnStatus

		break
	case rewardM.MAXONE:

		var maxBonus float64 = 0.0

		var atLeastOnePass = false

		finalPayload := &payloadM.PayloadEventResp{}

		for _, p := range payloadEventResps {

			if p.FeedReturn.PointReturn.PointReturnStatus == feedbackM.ALL_RETURN_POINT ||
				p.FeedReturn.PointReturn.PointReturnStatus == feedbackM.SOME_RETURN_POINT {

				if maxBonus <= p.FeedReturn.PointReturn.PointReturnBonus {
					finalPayload = p
					maxBonus = p.FeedReturn.PointReturn.PointReturnBonus
				}

				atLeastOnePass = true

			}
		}

		if atLeastOnePass {
			pointReturn.ActualPointReturn = finalPayload.FeedReturn.PointReturn.ActualPointReturn
			pointReturn.ActualUseCash = finalPayload.FeedReturn.PointReturn.ActualUseCash
			pointReturn.PointReturnBonus = maxBonus
			pointReturn.PointReturnStatus = finalPayload.FeedReturn.PointReturn.PointReturnStatus
		}

		break
	default:
		logrus.Error("rewardComponeont.calculatePointReturn no case")

	}

	logrus.Info("rewardComponent.calculatePointReturn ", pointReturn)

	return pointReturn, nil
}

func (im *impl) calculateCashReturn(ctx context.Context, e *eventM.Event, operator rewardM.PayloadOperator, payloadEventResps []*payloadM.PayloadEventResp) (*feedbackM.CashReturn, error) {

	logrus.Info("rewardComponent.calculateCashReturn")
	cashReturn := &feedbackM.CashReturn{
		Cash: e.Cash,
	}

	switch operator {
	case rewardM.ADD:

		var cashReturnBonus float64 = 0.0

		var actualUseCash int64 = 0
		var actualCashReturn float64 = 0.0
		var cashReturnStatus feedbackM.CashReturnStatus = feedbackM.NONE_RETURN_CASH

		hasReturn := false

		for _, p := range payloadEventResps {

			switch p.FeedReturn.CashReturn.CashReturnStatus {
			case feedbackM.ALL_RETURN_CASH:
				if actualUseCash < p.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = p.FeedReturn.CashReturn.ActualUseCash
				}

				actualCashReturn += p.FeedReturn.CashReturn.ActualCashReturn
				cashReturnBonus += p.FeedReturn.CashReturn.CashReturnBonus

				if !hasReturn {
					cashReturnStatus = feedbackM.ALL_RETURN_CASH
				} else {
					cashReturnStatus = feedbackM.SOME_RETURN_CASH
				}

				hasReturn = true

				break
			case feedbackM.SOME_RETURN_CASH:
				hasReturn = true
				if actualUseCash < p.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = p.FeedReturn.CashReturn.ActualUseCash
				}

				if actualUseCash < p.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = p.FeedReturn.CashReturn.ActualUseCash
				}
				actualCashReturn += p.FeedReturn.CashReturn.ActualCashReturn
				cashReturnBonus += p.FeedReturn.CashReturn.CashReturnBonus

				cashReturnStatus = feedbackM.SOME_RETURN_CASH

				break
			case feedbackM.NONE_RETURN_CASH:

				hasReturn = true
				if cashReturnStatus == feedbackM.ALL_RETURN_CASH {
					cashReturnStatus = feedbackM.SOME_RETURN_CASH
				}

				break
			}

		}

		logrus.Info("cashReturnStatus, ", cashReturnStatus)
		cashReturn.CashReturnStatus = cashReturnStatus
		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CashReturnBonus = cashReturnBonus

		break
	case rewardM.MAXONE:

		var maxBonus float64 = 0.0

		var atLeastOnePass = false

		finalPayload := &payloadM.PayloadEventResp{}
		for _, p := range payloadEventResps {

			logrus.Info(p.FeedReturn.CashReturn.CashReturnStatus)
			if p.FeedReturn.CashReturn.CashReturnStatus == feedbackM.ALL_RETURN_CASH ||
				p.FeedReturn.CashReturn.CashReturnStatus == feedbackM.SOME_RETURN_CASH {

				logrus.Info(p.FeedReturn.CashReturn.CashReturnBonus)
				if maxBonus <= p.FeedReturn.CashReturn.CashReturnBonus {
					finalPayload = p
					maxBonus = p.FeedReturn.CashReturn.CashReturnBonus
				}

				atLeastOnePass = true

			}
		}

		if atLeastOnePass {
			cashReturn.ActualCashReturn = finalPayload.FeedReturn.CashReturn.ActualCashReturn
			cashReturn.ActualUseCash = finalPayload.FeedReturn.CashReturn.ActualUseCash
			cashReturn.CashReturnBonus = maxBonus
			cashReturn.CashReturnStatus = finalPayload.FeedReturn.CashReturn.CashReturnStatus
		}

		break
	default:
		logrus.Error("rewardComponeont.calculateCashReturn no case")
	}

	return cashReturn, nil
}
