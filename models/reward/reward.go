package reward

import (
	"example.com/creditcard/models/constraint"
)

type RewardType int32

const (
	Bonus RewardType = iota
	Dollar
)

type Reward struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`

	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`

	StartDate  int64 `json:"startDate,omitempty"`
	EndDate    int64 `json:"endDate,omitempty"`
	UpdateDate int64 `json:"updateDate,omitempty"`

	ConstraintPayload *constraint.ConstraintPayload `json:"constraintPayload,omitempty"`
}
