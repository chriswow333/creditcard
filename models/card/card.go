package card

import (
	"example.com/creditcard/models/reward"
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

	RewardOperator RewardOperator   `json:"rewardOperator"`
	Rewards        []*reward.Reward `json:"rewards"`
}
