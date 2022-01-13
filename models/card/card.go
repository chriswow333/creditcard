package card

import (
	"example.com/creditcard/models/reward"
)

type Card struct {
	ID     string `json:"id"`
	BankID string `json:"bankID,omitempty"`
	Name   string `json:"name,omitempty"`
	Desc   string `json:"desc,omitempty"`

	StartDate  int64 `json:"startDate,omitempty"`
	EndDate    int64 `json:"endDate,omitempty"`
	UpdateDate int64 `json:"updateDate,omitempty"`

	LinkURL string `json:"linkURL,omitempty"`

	Rewards []*reward.Reward `json:"rewards,omitempty"`
}
