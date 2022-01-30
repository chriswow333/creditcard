package reward

import (
	"example.com/creditcard/models/constraint"
)

type RewardType int32

const (
	Cash RewardType = iota + 1
	Point
)

type Reward struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`

	Name string `json:"name"`
	Desc string `json:"desc"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	RewardType        RewardType                    `json:"rewardType"`
	ConstraintPayload *constraint.ConstraintPayload `json:"constraintPayload"`
}
