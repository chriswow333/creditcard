package reward

import (
	"example.com/creditcard/models/bonus"
	"example.com/creditcard/models/constraint"
)

type Reward struct {
	ID         string       `json:"id"`
	CardID     string       `json:"cardID"`
	Name       string       `json:"name"`
	Desc       string       `json:"desc"`
	StartDate  int64        `json:"startDate"`
	EndDate    int64        `json:"endDate"`
	UpdateDate int64        `json:"updateDate"`
	Bonus      *bonus.Bonus `json:"bonus,omitempty"`

	Operator    constraint.OperatorType  `json:"operator,omitempty"`
	Constraints []*constraint.Constraint `json:"constraints,omitempty"`
}
