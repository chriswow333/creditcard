package reward

import (
	"example.com/creditcard/models/constraint"
	"example.com/creditcard/models/cost"
)

type RewardType int32

const (
	Bonus RewardType = iota
	Dollar
)

type Reward struct {
	ID         string `json:"id"`
	CardID     string `json:"cardID"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	StartDate  int64  `json:"startDate"`
	EndDate    int64  `json:"endDate"`
	UpdateDate int64  `json:"updateDate"`

	RewardType RewardType `json:"rewardType"`

	Cost *cost.Cost `json:"cost"`

	Operator    constraint.OperatorType  `json:"operator,omitempty"`
	Constraints []*constraint.Constraint `json:"constraints,omitempty"`
}
