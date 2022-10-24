package card

import (
	"context"
	"errors"
	"fmt"
	"time"

	"example.com/creditcard/components/reward"
	cardM "example.com/creditcard/models/card"
	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
	"example.com/creditcard/service/bank"
	"github.com/sirupsen/logrus"

	feedbackDescStore "example.com/creditcard/stores/feedback_desc"
)

type impl struct {
	card                     *cardM.Card
	rewardMapper             map[string][]*reward.Component
	cardRewardOperatorMapper map[rewardM.RewardType]cardM.CardRewardOperator
	bankService              bank.Service
	feedbackDescStore        feedbackDescStore.Store
}

func New(
	card *cardM.Card,
	rewardMapper map[string][]*reward.Component,
	cardRewardOperatorMapper map[rewardM.RewardType]cardM.CardRewardOperator,
	bankService bank.Service,
	feedbackDescStore feedbackDescStore.Store,
) Component {

	impl := &impl{
		card:                     card,
		rewardMapper:             rewardMapper,
		cardRewardOperatorMapper: cardRewardOperatorMapper,
		bankService:              bankService,
		feedbackDescStore:        feedbackDescStore,
	}

	return impl
}

const DATE_FORMAT = "2006/01/02"

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*cardM.CardEventResp, error) {

	bankVo, err := im.bankService.GetByID(ctx, im.card.BankID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"not bankVo ": err,
		}).Error(err)
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

		feedbackDesc, err := im.feedbackDescStore.GetByID(ctx, cr.FeedbackDescID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"not found feedback ": err,
			}).Error(err)
			return nil, err
		}

		cardRewardEventResp := &cardM.CardRewardEventResp{
			ID:                 cr.ID,
			CardRewardOperator: cr.CardRewardOperator,
			RewardType:         cr.RewardType,
			CardRewardBonus:    cr.CardRewardBonus,

			Title:                cr.Title,
			Descs:                cr.Descs,
			StartDate:            startDate,
			EndDate:              endDate,
			CardRewardLimitTypes: cr.CardRewardLimitTypes,
			FeedbackDesc:         feedbackDesc,
		}

		if rewardComps, ok := im.rewardMapper[cr.ID]; ok {

			rewardEventResps := []*rewardM.RewardEventResp{}

			for _, rc := range rewardComps {
				rewardEventResp, err := (*rc).Satisfy(ctx, e)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"card component": err,
					}).Error(err)
					return nil, err
				}
				rewardEventResps = append(rewardEventResps, rewardEventResp)
			}

			im.calculateReturn(ctx, cr, rewardEventResps, cardRewardEventResp)

			cardEventResp.CardRewardEventResps = append(cardEventResp.CardRewardEventResps, cardRewardEventResp)

		} else {

			logrus.WithFields(logrus.Fields{
				"not found card reward ": err,
			}).Error(err)

		}
	}

	return cardEventResp, nil
}

func (im *impl) calculateReturn(ctx context.Context,
	cr *cardM.CardReward, rewardEventResps []*rewardM.RewardEventResp,
	cardRewardEventResp *cardM.CardRewardEventResp) error {

	switch cr.RewardType {
	case rewardM.CASH:

		cashReturn := im.calculateCashFeedReturn(ctx, cr.CardRewardOperator, rewardEventResps)

		cardRewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			FeedReturnStatus: feedbackM.ALL,
			CashReturn:       cashReturn,
		}
		break
	case rewardM.POINT:

		pointReturn := im.calculatePointFeedReturn(ctx, cr.CardRewardOperator, rewardEventResps)

		cardRewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			FeedReturnStatus: feedbackM.ALL,
			PointReturn:      pointReturn,
		}
		break

	case rewardM.RED:

		redReturn := im.calculateRedFeedReturn(ctx, cr.CardRewardOperator, rewardEventResps)

		cardRewardEventResp.FeedReturn = &feedbackM.FeedReturn{
			FeedReturnStatus: feedbackM.ALL,
			RedReturn:        redReturn,
		}

		break

	default:
		return errors.New("no suitable reward type.")

	}

	return nil
}

func (im *impl) calculateRedFeedReturn(ctx context.Context, cardRewardOperator cardM.CardRewardOperator, rewardEventResps []*rewardM.RewardEventResp) *feedbackM.RedReturn {

	redReturn := &feedbackM.RedReturn{}

	var totalCash float64 = 0.0
	var currentCash int64 = 0

	var isRedbackGet bool = false
	var redbackTimes int64 = 0.0

	var actualUseCash int64 = 0
	var actualRedback float64 = 0.0

	switch cardRewardOperator {
	case cardM.ADD:

		for _, r := range rewardEventResps {

			if r.RewardEventJudgeType == rewardM.NONE {
				continue
			}

			totalCash = r.FeedReturn.RedReturn.TotalCash
			currentCash = r.FeedReturn.RedReturn.CurrentCash

			if r.FeedReturn.RedReturn.IsRedGet {

				isRedbackGet = true

				if actualUseCash < r.FeedReturn.RedReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.RedReturn.ActualUseCash
				}
				redbackTimes += r.FeedReturn.RedReturn.RedbackTimes
				actualRedback += r.FeedReturn.RedReturn.ActualRedback

			}

		}

		redReturn.IsRedGet = isRedbackGet
		redReturn.ActualRedback = actualRedback
		redReturn.ActualUseCash = actualUseCash
		redReturn.CurrentCash = currentCash
		redReturn.TotalCash = totalCash
		redReturn.RedbackTimes = redbackTimes

		break

	case cardM.MAXONE:
		for _, r := range rewardEventResps {
			if r.FeedReturn.RedReturn.IsRedGet {
				if redbackTimes <= r.FeedReturn.RedReturn.RedbackTimes {
					isRedbackGet = true
					redbackTimes = r.FeedReturn.RedReturn.RedbackTimes
					actualUseCash = r.FeedReturn.RedReturn.ActualUseCash
					actualRedback = r.FeedReturn.RedReturn.ActualRedback
					currentCash = r.FeedReturn.RedReturn.CurrentCash
					totalCash = r.FeedReturn.RedReturn.TotalCash
				}
			}
		}

		redReturn.IsRedGet = isRedbackGet
		redReturn.ActualRedback = actualRedback
		redReturn.ActualUseCash = actualUseCash
		redReturn.CurrentCash = currentCash
		redReturn.TotalCash = totalCash
		redReturn.RedbackTimes = redbackTimes

		break

	default:
		logrus.Error("calculatePointFeedReturn Error operator")

		redReturn.IsRedGet = isRedbackGet
		redReturn.ActualRedback = actualRedback
		redReturn.ActualUseCash = actualUseCash
		redReturn.CurrentCash = currentCash
		redReturn.TotalCash = totalCash
		redReturn.RedbackTimes = redbackTimes
	}
	return redReturn
}

func (im *impl) calculatePointFeedReturn(ctx context.Context,
	cardRewardOperator cardM.CardRewardOperator, rewardEventResps []*rewardM.RewardEventResp) *feedbackM.PointReturn {

	pointReturn := &feedbackM.PointReturn{}

	var totalCash float64 = 0.0
	var currentCash int64 = 0

	var isPointbackGet bool = false
	var pointbackBonus float64 = 0.0

	var actualUseCash int64 = 0
	var actualPointBack float64 = 0.0

	switch cardRewardOperator {
	case cardM.ADD:

		for _, r := range rewardEventResps {

			if r.RewardEventJudgeType == rewardM.NONE {
				continue
			}

			totalCash = r.FeedReturn.PointReturn.TotalCash
			currentCash = r.FeedReturn.PointReturn.CurrentCash

			if r.FeedReturn.PointReturn.IsPointbackGet {

				isPointbackGet = true

				if actualUseCash < r.FeedReturn.PointReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.PointReturn.ActualUseCash
				}
				pointbackBonus += r.FeedReturn.PointReturn.PointbackBonus
				actualPointBack += r.FeedReturn.PointReturn.ActualPointBack

			}

		}

		pointReturn.IsPointbackGet = isPointbackGet
		pointReturn.ActualPointBack = actualPointBack
		pointReturn.ActualUseCash = actualUseCash
		pointReturn.CurrentCash = currentCash
		pointReturn.TotalCash = totalCash
		pointReturn.PointbackBonus = pointbackBonus

		break

	case cardM.MAXONE:
		for _, r := range rewardEventResps {
			fmt.Println(r.FeedReturn)
			if r.FeedReturn.PointReturn.IsPointbackGet {
				if pointbackBonus <= r.FeedReturn.PointReturn.PointbackBonus {
					isPointbackGet = true
					pointbackBonus = r.FeedReturn.PointReturn.PointbackBonus
					actualUseCash = r.FeedReturn.PointReturn.ActualUseCash
					actualPointBack = r.FeedReturn.PointReturn.ActualPointBack
					currentCash = r.FeedReturn.PointReturn.CurrentCash
					totalCash = r.FeedReturn.PointReturn.TotalCash
				}
			}
		}

		pointReturn.IsPointbackGet = isPointbackGet
		pointReturn.ActualPointBack = actualPointBack
		pointReturn.ActualUseCash = actualUseCash
		pointReturn.CurrentCash = currentCash
		pointReturn.TotalCash = totalCash
		pointReturn.PointbackBonus = pointbackBonus

		break

	default:
		logrus.Error("calculatePointFeedReturn Error operator")

		pointReturn.IsPointbackGet = isPointbackGet
		pointReturn.ActualPointBack = actualPointBack
		pointReturn.ActualUseCash = actualUseCash
		pointReturn.CurrentCash = currentCash
		pointReturn.TotalCash = totalCash
		pointReturn.PointbackBonus = pointbackBonus
	}

	return pointReturn

}

func (im *impl) calculateCashFeedReturn(ctx context.Context, cardRewardOperator cardM.CardRewardOperator, rewardEventResps []*rewardM.RewardEventResp) *feedbackM.CashReturn {

	cashReturn := &feedbackM.CashReturn{}

	var totalCash float64 = 0.0
	var currentCash int64 = 0

	var isCashbackGet bool = false
	var cashbackBonus float64 = 0.0

	var actualUseCash int64 = 0
	var actualCashReturn float64 = 0.0

	switch cardRewardOperator {
	case cardM.ADD:

		for _, r := range rewardEventResps {

			if r.RewardEventJudgeType == rewardM.NONE {
				continue
			}

			totalCash = r.FeedReturn.CashReturn.TotalCash
			currentCash = r.FeedReturn.CashReturn.CurrentCash

			if r.FeedReturn.CashReturn.IsCashbackGet {
				isCashbackGet = true

				if actualUseCash < r.FeedReturn.CashReturn.ActualUseCash {
					// get max actual use cash
					actualUseCash = r.FeedReturn.CashReturn.ActualUseCash
				}
				cashbackBonus += r.FeedReturn.CashReturn.CashbackBonus
				actualCashReturn += r.FeedReturn.CashReturn.ActualCashReturn

			}

		}

		cashReturn.IsCashbackGet = isCashbackGet
		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CurrentCash = currentCash
		cashReturn.TotalCash = totalCash
		cashReturn.CashbackBonus = cashbackBonus

		break

	case cardM.MAXONE:

		for _, r := range rewardEventResps {
			if r.FeedReturn.CashReturn.IsCashbackGet {
				if cashbackBonus <= r.FeedReturn.CashReturn.CashbackBonus {
					isCashbackGet = true
					cashbackBonus = r.FeedReturn.CashReturn.CashbackBonus
					actualUseCash = r.FeedReturn.CashReturn.ActualUseCash
					actualCashReturn = r.FeedReturn.CashReturn.ActualCashReturn
					currentCash = r.FeedReturn.CashReturn.CurrentCash
					totalCash = r.FeedReturn.CashReturn.TotalCash
				}
			}
		}

		cashReturn.IsCashbackGet = isCashbackGet
		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CurrentCash = currentCash
		cashReturn.TotalCash = totalCash
		cashReturn.CashbackBonus = cashbackBonus
		break

	default:
		logrus.Error("Error operator")

		cashReturn.IsCashbackGet = isCashbackGet
		cashReturn.ActualCashReturn = actualCashReturn
		cashReturn.ActualUseCash = actualUseCash
		cashReturn.CurrentCash = currentCash
		cashReturn.TotalCash = totalCash
		cashReturn.CashbackBonus = cashbackBonus
	}

	return cashReturn
}
