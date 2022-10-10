package reward

import (
	"example.com/creditcard/models/payload"
)

type RewardType int32

const (
	CASH RewardType = iota + 1
	POINT
	RED
)

type PayloadOperator int32

const (
	ADD PayloadOperator = iota + 1
	MAXONE
)

type Reward struct {
	ID           string `json:"id"`
	CardRewardID string `json:"cardRewardID"`
	Order        int32  `json:"order"`

	PayloadOperator PayloadOperator    `json:"payloadOperator,omitempty"`
	Payloads        []*payload.Payload `json:"payloads,omitempty"`
}
