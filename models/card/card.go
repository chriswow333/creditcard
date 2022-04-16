package card

import (
	"time"

	"example.com/creditcard/models/reward"
	rewardM "example.com/creditcard/models/reward"
)

type Card struct {
	ID     string `json:"id"`
	BankID string `json:"bankID"`
	Name   string `json:"name,omitempty"`

	StartDate  int64 `json:"startDate,omitempty"`
	EndDate    int64 `json:"endDate,omitempty"`
	UpdateDate int64 `json:"updateDate,omitempty"`

	ImagePath string `json:"imagePath,omitempty"`
	LinkURL   string `json:"linkURL,omitempty"`

	CardRewards []*CardReward `json:"cardRewards,omitempty"`
}

type CardRewardOperator int32

const (
	ADD CardRewardOperator = iota + 1
	MAXONE
)

type CardReward struct {
	ID             string `json:"id"`
	CardID         string `json:"cardID"`
	CardRewardDesc string `json:"cardRewardDesc"`

	CardRewardOperator CardRewardOperator `json:"cardRewardOperator,omitempty"` // (R0+(R1&(R2|R3)))
	RewardType         reward.RewardType  `json:"rewardType,omitempty"`

	ConstraintPassLogics []*ConstraintPassLogic `json:"constraintPassLogics"`

	Rewards []*rewardM.Reward `json:"rewards,omitempty"`
}

type ConstraintPassLogic struct {
	Logic   string `json:"logic"`
	Message string `json:"message"`
}

type CardRewardEventJudgeType int32

const (
	ALL CardRewardEventJudgeType = iota + 1
	SOME
	NONE
)

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
		case reward.CASH_TWD:

			rewardResps := []*reward.RewardResp{}
			// rewards
			for _, r := range c.Rewards {
				rewardResp := rewardM.TransferRewardResp(c.RewardType, r)
				rewardResps = append(rewardResps, rewardResp)
			}

			cardRewardResp := &CardRewardResp{
				ID:                   c.ID,
				CardID:               c.CardID,
				CardRewardDesc:       c.CardRewardDesc,
				CardRewardOperator:   c.CardRewardOperator,
				RewardType:           c.RewardType,
				RewardResps:          rewardResps,
				ConstraintPassLogics: c.ConstraintPassLogics,
			}

			cardRewardResps = append(cardRewardResps, cardRewardResp)

		case reward.POINT:

		}

	}

	cardResp.CardRewardResps = cardRewardResps

	return cardResp
}
