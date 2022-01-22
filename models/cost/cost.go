package cost

import (
	"example.com/creditcard/models/bonus"
	"example.com/creditcard/models/dollar"
)

type CostType int32

const (
	Dollar = iota
	Bonus
)

type Cost struct {
	CostType CostType `json:"costType"`

	RewardID string `json:"rewardID"`

	IsRewardGet bool           `json:"isRewardGet"`
	Dollar      *dollar.Dollar `json:"dollar,omitempty"` // 現金回饋
	Bonus       *bonus.Bonus   `json:"bonus,omitempty"`  // 點數回饋

}
