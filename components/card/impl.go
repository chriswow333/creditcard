package card

import (
	"context"
	"errors"
	"fmt"

	"example.com/creditcard/components/reward"
	cardM "example.com/creditcard/models/card"
	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
	"github.com/sirupsen/logrus"
)

type impl struct {
	cardResp                 *cardM.CardResp
	rewardMapper             map[rewardM.RewardType][]*reward.Component
	cardRewardOperatorMapper map[rewardM.RewardType]cardM.CardRewardOperator
}

func New(
	cardResp *cardM.CardResp,
	rewardMapper map[rewardM.RewardType][]*reward.Component,
	cardRewardOperatorMapper map[rewardM.RewardType]cardM.CardRewardOperator,
) Component {

	impl := &impl{
		cardResp:                 cardResp,
		rewardMapper:             rewardMapper,
		cardRewardOperatorMapper: cardRewardOperatorMapper,
	}

	return impl
}

const DATE_FORMAT = "2006/01/02"

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*cardM.CardEventResp, error) {

	if e.RewardType == 0 {
		return nil, errors.New("need reward type to evaluate")
	}

	rewardComps, ok := im.rewardMapper[e.RewardType]
	if !ok {
		return nil, errors.New("not found rewardComponent")
	}

	cardRewardResp := &cardM.CardRewardResp{}
	for _, cr := range im.cardResp.CardRewardResps {
		if cr.RewardType == e.RewardType {
			cardRewardResp = cr
			break
		}

	}

	cardEventResp := &cardM.CardEventResp{
		ID:     im.cardResp.ID,
		BankID: im.cardResp.BankID,
	}

	cardRewardEventResp := &cardM.CardRewardEventResp{
		ID:     cardRewardResp.ID,
		CardID: cardRewardResp.CardID,

		CardRewardOperator: cardRewardResp.CardRewardOperator,
		RewardType:         e.RewardType,
	}

	cardEventResp.CardRewardEventResp = cardRewardEventResp

	rewardEventResps := []*rewardM.RewardEventResp{}

	for _, rc := range rewardComps {
		rewardEventResp, err := (*rc).Satisfy(ctx, e)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"card component": "",
			}).Error(err)
			return nil, err
		}

		rewardEventResps = append(rewardEventResps, rewardEventResp)
	}

	cashReturn, err := im.calculateCashFeedReturn(ctx, cardRewardResp.CardRewardOperator, rewardEventResps)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"card component": "",
		}).Error(err)
		return nil, err
	}

	cardRewardEventResp.FeedReturn = &feedbackM.FeedReturn{
		FeedReturnStatus: feedbackM.ALL,
		CashReturn:       cashReturn,
	}

	cardRewardEventResp.RewardEventResps = rewardEventResps

	return cardEventResp, nil
}

func (im *impl) calculateCashFeedReturn(ctx context.Context, cardRewardOperator cardM.CardRewardOperator, rewardEventResps []*rewardM.RewardEventResp) (*feedbackM.CashReturn, error) {

	cashReturn := &feedbackM.CashReturn{}

	switch cardRewardOperator {
	case cardM.ADD:
		var totalCash float64 = 0.0
		var currentCash int64 = 0

		var isCashbackGet bool = false
		var cashbackBonus float64 = 0.0

		var actualUseCash int64 = 0
		var actualCashReturn float64 = 0.0

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

		var isCashbackGet bool = false
		var cashbackBonus float64 = 0.0
		var actualUseCash int64 = 0
		var actualCashReturn float64 = 0.0
		var currentCash int64 = 0
		var totalCash float64 = 0

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
		fmt.Println("Error operator")
	}

	return cashReturn, nil
}
