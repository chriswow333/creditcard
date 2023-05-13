package card

import (
	"context"
	"errors"
	"runtime/debug"
	"time"

	"example.com/creditcard/components/reward"
	cardM "example.com/creditcard/models/card"
	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
	"example.com/creditcard/service/bank"
	"github.com/sirupsen/logrus"
)

type impl struct {
	card                     *cardM.Card
	rewardMapper             map[string][]*reward.Component
	cardRewardOperatorMapper map[rewardM.RewardType]cardM.CardRewardOperator
	bankService              bank.Service
}

func New(
	card *cardM.Card,
	rewardMapper map[string][]*reward.Component,
	cardRewardOperatorMapper map[rewardM.RewardType]cardM.CardRewardOperator,
	bankService bank.Service,
) Component {

	impl := &impl{
		card:                     card,
		rewardMapper:             rewardMapper,
		cardRewardOperatorMapper: cardRewardOperatorMapper,
		bankService:              bankService,
	}

	return impl
}

const DATE_FORMAT = "2006/01/02"

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*cardM.CardEventResp, error) {

	bankVo, err := im.bankService.GetByID(ctx, im.card.BankID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	updateDate := time.Unix(im.card.UpdateDate, 0).Format(DATE_FORMAT)

	cardEventResp := &cardM.CardEventResp{

		ID:    im.card.ID,
		Name:  im.card.Name,
		Descs: im.card.Descs,

		ImagePath: im.card.ImagePath,

		BankID:   im.card.BankID,
		BankName: bankVo.Name,

		UpdateDate: updateDate,
	}

	for _, cr := range im.card.CardRewards {

		if len(e.CardRewardIDs) != 0 {

			matchedCardRewardID := false

			for _, cardRewardID := range e.CardRewardIDs {
				if cr.ID == cardRewardID {
					matchedCardRewardID = true
					break
				}
			}

			if !matchedCardRewardID {
				continue
			}
		}

		// mismatch reward type
		if cr.RewardType != e.RewardType && e.RewardType != 0 {
			continue
		}

		startDate := time.Unix(cr.StartDate, 0).Format(DATE_FORMAT)
		endDate := time.Unix(cr.EndDate, 0).Format(DATE_FORMAT)

		cardRewardEventResp := &cardM.CardRewardEventResp{
			ID:                 cr.ID,
			CardRewardOperator: cr.CardRewardOperator,
			RewardType:         cr.RewardType,

			Title:                cr.Title,
			Descs:                cr.Descs,
			StartDate:            startDate,
			EndDate:              endDate,
			CardRewardLimitTypes: cr.CardRewardLimitTypes,
			FeedbackBonus:        cr.FeedbackBonus,
		}

		if rewardComps, ok := im.rewardMapper[cr.ID]; ok {

			rewardEventResps := []*rewardM.RewardEventResp{}

			for _, rc := range rewardComps {
				rewardEventResp, err := (*rc).Satisfy(ctx, e)
				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					return nil, err
				}
				rewardEventResps = append(rewardEventResps, rewardEventResp)
			}

			im.calculateReturn(ctx, e, cr, rewardEventResps, cardRewardEventResp)

			cardEventResp.CardRewardEventResps = append(cardEventResp.CardRewardEventResps, cardRewardEventResp)

		} else {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		}
	}

	return cardEventResp, nil
}

func (im *impl) calculateReturn(ctx context.Context, e *eventM.Event,
	cr *cardM.CardReward, rewardEventResps []*rewardM.RewardEventResp,
	cardRewardEventResp *cardM.CardRewardEventResp) error {

	logrus.Info("card.calculateReturn")

	switch cr.RewardType {
	case rewardM.CASH:

		cashReturn := im.calculateCashFeedReturn(ctx, e, cr.CardRewardOperator, rewardEventResps)

		logrus.Info("card reward cash ", cashReturn)
		cardRewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			CashReturn: cashReturn,
		}

		break
	case rewardM.POINT:

		pointReturn := im.calculatePointFeedReturn(ctx, e, cr.CardRewardOperator, rewardEventResps)

		cardRewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			PointReturn: pointReturn,
		}
		break

	default:
		return errors.New("no suitable reward type.")
	}

	return nil
}

func (im *impl) calculatePointFeedReturn(ctx context.Context, e *eventM.Event,
	cardRewardOperator cardM.CardRewardOperator, rewardEventResps []*rewardM.RewardEventResp) *feedbackM.PointReturn {

	pointReturn := &feedbackM.PointReturn{
		Cash: e.Cash,
	}

	var pointReturnBonus float64 = 0.0

	var actualUseCash int64 = 0
	var actualPointReturn float64 = 0.0
	var pointReturnStatus feedbackM.PointReturnStatus = feedbackM.NONE_RETURN_POINT
	hasReturn := false

	switch cardRewardOperator {
	case cardM.ADD:

		for _, r := range rewardEventResps {

			switch r.FeedReturn.PointReturn.PointReturnStatus {
			case feedbackM.ALL_RETURN_POINT:
				if actualUseCash < r.FeedReturn.PointReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.PointReturn.ActualUseCash
				}

				actualPointReturn += r.FeedReturn.PointReturn.ActualPointReturn
				pointReturnBonus += r.FeedReturn.PointReturn.PointReturnBonus

				if !hasReturn {
					pointReturnStatus = feedbackM.ALL_RETURN_POINT
				} else {
					pointReturnStatus = feedbackM.SOME_RETURN_POINT
				}

				logrus.Info("cardRewardComponent.calculatePointFeedReturn ", pointReturnStatus)
				hasReturn = true
				break
			case feedbackM.SOME_RETURN_POINT:
				hasReturn = true
				if actualUseCash < r.FeedReturn.PointReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.PointReturn.ActualUseCash
				}

				if actualUseCash < r.FeedReturn.PointReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.PointReturn.ActualUseCash
				}
				actualPointReturn += r.FeedReturn.PointReturn.ActualPointReturn
				pointReturnBonus += r.FeedReturn.PointReturn.PointReturnBonus

				pointReturnStatus = feedbackM.SOME_RETURN_POINT

				break
			case feedbackM.NONE_RETURN_POINT:
				hasReturn = true
				if pointReturnStatus == feedbackM.ALL_RETURN_POINT {
					pointReturnStatus = feedbackM.SOME_RETURN_POINT
				}
				break
			}
		}

		pointReturn.ActualPointReturn = actualPointReturn
		pointReturn.PointReturnBonus = pointReturnBonus * 100
		pointReturn.ActualUseCash = actualUseCash
		pointReturn.PointReturnStatus = pointReturnStatus

		break

	case cardM.MAXONE:
		for _, r := range rewardEventResps {

			if pointReturnBonus <= r.FeedReturn.PointReturn.PointReturnBonus {
				pointReturnBonus = r.FeedReturn.PointReturn.PointReturnBonus
				actualUseCash = r.FeedReturn.PointReturn.ActualUseCash
				actualPointReturn = r.FeedReturn.PointReturn.ActualPointReturn
				pointReturnStatus = r.FeedReturn.PointReturn.PointReturnStatus
			}

		}

		pointReturn.ActualPointReturn = actualPointReturn
		pointReturn.ActualUseCash = actualUseCash
		pointReturn.PointReturnBonus = pointReturnBonus * 100
		pointReturn.PointReturnStatus = pointReturnStatus

		break

	default:
		logrus.Error("calculatePointFeedReturn Error operator")

		pointReturn.ActualPointReturn = actualPointReturn
		pointReturn.ActualUseCash = actualUseCash
		pointReturn.PointReturnBonus = pointReturnBonus
		// pointReturn.PointReturnStatus = pointReturnStatus
	}

	logrus.Info("cardRewardComponent.calculatePointFeedReturn result ", pointReturn)

	return pointReturn

}

func (im *impl) calculateCashFeedReturn(ctx context.Context, e *eventM.Event, cardRewardOperator cardM.CardRewardOperator, rewardEventResps []*rewardM.RewardEventResp) *feedbackM.CashReturn {
	logrus.Info("cardCompontnet.calculateCashFeedReturn")
	logrus.Info("cardRewardOperator ", cardRewardOperator)

	cashReturn := &feedbackM.CashReturn{
		Cash: e.Cash,
	}

	var cashReturnBonus float64 = 0.0

	var actualUseCash int64 = 0
	var actualCashReturn float64 = 0.0

	var cashReturnStatus feedbackM.CashReturnStatus = feedbackM.NONE_RETURN_CASH
	hasReturn := false

	switch cardRewardOperator {
	case cardM.ADD:
		logrus.Info("is added")
		for _, r := range rewardEventResps {
			switch r.FeedReturn.CashReturn.CashReturnStatus {
			case feedbackM.ALL_RETURN_CASH:
				if actualUseCash < r.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.CashReturn.ActualUseCash
				}

				actualCashReturn += r.FeedReturn.CashReturn.ActualCashReturn
				cashReturnBonus += r.FeedReturn.CashReturn.CashReturnBonus

				if !hasReturn {
					cashReturnStatus = feedbackM.ALL_RETURN_CASH
				} else {
					cashReturnStatus = feedbackM.SOME_RETURN_CASH
				}

				hasReturn = true

				break
			case feedbackM.SOME_RETURN_CASH:
				hasReturn = true
				if actualUseCash < r.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.CashReturn.ActualUseCash
				}

				if actualUseCash < r.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.CashReturn.ActualUseCash
				}
				actualCashReturn += r.FeedReturn.CashReturn.ActualCashReturn
				cashReturnBonus += r.FeedReturn.CashReturn.CashReturnBonus

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

		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CashReturnBonus = cashReturnBonus * 100
		cashReturn.CashReturnStatus = cashReturnStatus

		break

	case cardM.MAXONE:
		for _, r := range rewardEventResps {

			if cashReturnBonus <= r.FeedReturn.CashReturn.CashReturnBonus {
				cashReturnBonus = r.FeedReturn.CashReturn.CashReturnBonus
				actualUseCash = r.FeedReturn.CashReturn.ActualUseCash
				actualCashReturn = r.FeedReturn.CashReturn.ActualCashReturn
				cashReturnStatus = r.FeedReturn.CashReturn.CashReturnStatus
			}

		}

		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CashReturnStatus = cashReturnStatus
		cashReturn.CashReturnBonus = cashReturnBonus * 100
		break

	default:
		logrus.Error("Error operator")

		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CashReturnBonus = cashReturnBonus
	}

	logrus.Info("results:", cashReturn)

	return cashReturn
}
