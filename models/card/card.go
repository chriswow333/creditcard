package card

import (
	"fmt"
	"time"

	"example.com/creditcard/models/feedback"
	"example.com/creditcard/models/reward"
	rewardM "example.com/creditcard/models/reward"
)

type RewardOperator int32

const (
	AddRewardOperator RewardOperator = iota + 1
	XORHighRewardOperator
)

type Card struct {
	ID     string `json:"id"`
	BankID string `json:"bankID"`
	Name   string `json:"name"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	ImagePath string `json:"imagePath"`
	LinkURL   string `json:"linkURL"`

	CardRewards []*CardReward `json:"cardReward"`
}

type CardReward struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`

	RewardOperator RewardOperator    `json:"rewardOperator"`
	RewardType     reward.RewardType `json:"rewardType"`
	Rewards        []*rewardM.Reward `json:"rewards"`
}

type CardResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	BankID   string `json:"bankID"`
	BankName string `json:"bankName"`

	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	UpdateDate string `json:"updateDate"`

	ImagePath string `json:"imagePath"`
	LinkURL   string `json:"linkURL"`

	CardRewardResps []*CardRewardResp `json:"cardRewardResps"`
}

type CardRewardResp struct {
	RewardOperator RewardOperator    `json:"rewardOperator"`
	RewardType     reward.RewardType `json:"rewardType"`

	InCashRewardResp  *reward.InCashRewardResp  `json:"inCashRewardResp"`
	OutCashRewardResp *reward.OutCashRewardResp `json:"outCashRewardResp"`
}

const DATE_FORMAT = "2006/01/02"

func TransferCardResp(card *Card) *CardResp {

	cardResp := &CardResp{
		ID:         card.ID,
		BankID:     card.BankID,
		Name:       card.Name,
		StartDate:  time.Unix(card.StartDate, 0).Format(DATE_FORMAT),
		EndDate:    time.Unix(card.EndDate, 0).Format(DATE_FORMAT),
		UpdateDate: time.Unix(card.UpdateDate, 0).Format(DATE_FORMAT),
		LinkURL:    card.LinkURL,
		ImagePath:  card.ImagePath,
	}

	cardRewardResps := []*CardRewardResp{}

	for _, c := range card.CardRewards {

		switch c.RewardType {
		case reward.InCash:

			inCashRewardResps := []*reward.RewardResp{}
			// rewards
			for _, r := range c.Rewards {
				rewardResp := rewardM.TransferRewardResp(c.RewardType, r)
				inCashRewardResps = append(inCashRewardResps, rewardResp)
			}

			// cashback
			inCashCashbackResp := getOptimalCashbackResp(c.RewardOperator, inCashRewardResps)

			cardRewardResp := &CardRewardResp{
				InCashRewardResp: &rewardM.InCashRewardResp{
					RewardResps: inCashRewardResps,
					FeedbackResp: &feedback.FeedbackResp{
						CashbackResp: inCashCashbackResp,
					},
				},
				RewardOperator: c.RewardOperator,
				RewardType:     c.RewardType,
			}

			cardRewardResps = append(cardRewardResps, cardRewardResp)

		case reward.OutCash:

		case reward.Point:

		}

	}

	cardResp.CardRewardResps = cardRewardResps

	return cardResp
}

func getOptimalCashbackResp(rewardOperator RewardOperator, rewardResps []*reward.RewardResp) *feedback.CashbackResp {

	cashbackResp := &feedback.CashbackResp{}
	var min int64 = 0
	var max int64 = 99999999
	var bonus float64 = 0.0
	var cashbackType feedback.CashbackType

	fmt.Println("check  ", rewardOperator)

	for _, r := range rewardResps {

		cashbackType = r.FeedbackResp.CashbackResp.CashbackType

		switch rewardOperator {
		case XORHighRewardOperator:
			if bonus < r.FeedbackResp.CashbackResp.Bonus {
				min = r.FeedbackResp.CashbackResp.Min
				max = r.FeedbackResp.CashbackResp.Max
				bonus = r.FeedbackResp.CashbackResp.Bonus
			}

		case AddRewardOperator:

			bonus += r.FeedbackResp.CashbackResp.Bonus
			if min < r.FeedbackResp.CashbackResp.Min {
				min = r.FeedbackResp.CashbackResp.Min
			}

			if max > r.FeedbackResp.CashbackResp.Max {
				max = r.FeedbackResp.CashbackResp.Max
			}

		}
	}

	cashbackResp.CashbackType = cashbackType
	cashbackResp.Min = min
	cashbackResp.Max = max
	cashbackResp.Bonus = bonus

	return cashbackResp
}
