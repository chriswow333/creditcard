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
	"github.com/sirupsen/logrus"
)

type impl struct {
	card                  *cardM.Card
	rewardMapper          map[rewardM.RewardType][]*reward.Component
	payloadOperatorMapper map[rewardM.RewardType]cardM.RewardOperator
}

func New(
	card *cardM.Card,
	rewardMapper map[rewardM.RewardType][]*reward.Component,
	payloadOperatorMapper map[rewardM.RewardType]cardM.RewardOperator,
) Component {
	return &impl{
		card:                  card,
		rewardMapper:          rewardMapper,
		payloadOperatorMapper: payloadOperatorMapper,
	}
}

const DATE_FORMAT = "2006/01/02"

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*cardM.CardResp, error) {

	cardResp := &cardM.CardResp{
		ID:         im.card.ID,
		BankID:     im.card.BankID,
		Name:       im.card.Name,
		StartDate:  time.Unix(im.card.StartDate, 0).Format(DATE_FORMAT),
		EndDate:    time.Unix(im.card.EndDate, 0).Format(DATE_FORMAT),
		UpdateDate: time.Unix(im.card.UpdateDate, 0).Format(DATE_FORMAT),
		ImagePath:  im.card.ImagePath,
		LinkURL:    im.card.LinkURL,
	}

	cardRewardResps := []*cardM.CardRewardResp{}

	if e.RewardType != 0 {
		// means using specficed type

	} else {

		for rewardType, rs := range im.rewardMapper {

			rewardResps := []*rewardM.RewardResp{}

			for _, r := range rs {
				reward, err := (*r).Satisfy(ctx, e)

				if err != nil {
					logrus.WithFields(logrus.Fields{
						"card component": "",
					}).Error(err)
					return nil, err
				}
				rewardResps = append(rewardResps, reward)
			}

			switch rewardType {
			case rewardM.InCash:
				cardRewardResp := &cardM.CardRewardResp{
					InCashRewardResp: &rewardM.InCashRewardResp{
						RewardResps: rewardResps,
					},
				}

				// cardResp.InCashRewardResp = &rewardM.InCashRewardResp{
				// 	RewardResps: rewardResps,
				// }

				if payloadOperator, ok := im.payloadOperatorMapper[rewardType]; ok {
					cashReturn, err := im.calculateCashFeedReturn(ctx, payloadOperator, rewardResps)

					if err != nil {
						return nil, err
					}
					cardRewardResp.InCashRewardResp.FeedReturn = &feedbackM.FeedReturn{
						CashReturn: cashReturn,
					}
					cardRewardResp.RewardOperator = payloadOperator

				} else {
					return nil, errors.New("Cannot find reward mapper")
				}

				cardRewardResps = append(cardRewardResps, cardRewardResp)

				continue

			case rewardM.OutCash:
				fmt.Println("card component out cash")
				continue
			case rewardM.Point:
				// card.PointReward = cardReward
				fmt.Println("card component point")
				continue
			default:

			}
		}
	}

	cardResp.CardRewardResps = cardRewardResps
	return cardResp, nil
}

func (im *impl) calculateCashFeedReturn(ctx context.Context, operator cardM.RewardOperator, rewardResps []*rewardM.RewardResp) (*feedbackM.CashReturn, error) {

	cashReturn := &feedbackM.CashReturn{}

	fmt.Println("operator  ", operator)
	switch operator {
	case cardM.AddRewardOperator:

		var totalCash float64 = 0.0
		var currentCash int64 = 0

		var isCashbackGet bool = false
		var cashbackBonus float64 = 0.0

		var actualUseCash int64 = 0
		var actualCashReturn float64 = 0.0

		for _, r := range rewardResps {

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
	case cardM.XORHighRewardOperator:

		var maxBonus float64 = 0.0
		finalReward := &rewardM.RewardResp{}
		for _, r := range rewardResps {
			if r.FeedReturn.CashReturn.IsCashbackGet {
				if maxBonus < r.FeedReturn.CashReturn.CashbackBonus {
					finalReward = r
					maxBonus = r.FeedReturn.CashReturn.CashbackBonus
				}
			}
		}
		cashReturn.IsCashbackGet = finalReward.FeedReturn.CashReturn.IsCashbackGet
		cashReturn.ActualCashReturn = finalReward.FeedReturn.CashReturn.ActualCashReturn
		cashReturn.ActualUseCash = finalReward.FeedReturn.CashReturn.ActualUseCash
		cashReturn.CurrentCash = finalReward.FeedReturn.CashReturn.CurrentCash
		cashReturn.TotalCash = finalReward.FeedReturn.CashReturn.TotalCash
		cashReturn.CashbackBonus = maxBonus
		break
	default:
		fmt.Println("Error operator")
	}

	return cashReturn, nil
}
