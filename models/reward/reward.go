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
	ID     string `json:"id"`
	CardID string `json:"cardID"`
	Name   string `json:"name,omitempty"`
	Desc   string `json:"desc,omitempty"`

	StartDate  int64 `json:"startDate,omitempty"`
	EndDate    int64 `json:"endDate,omitempty"`
	UpdateDate int64 `json:"updateDate,omitempty"`

	Cost *cost.Cost `json:"cost"`

	Operator    constraint.OperatorType  `json:"operator"`
	Constraints []*constraint.Constraint `json:"constraints,omitempty"`
}
