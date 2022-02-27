package reward

import (
	"example.com/creditcard/models/payload"
)

type RewardType int32

const (
	InCash RewardType = iota + 1
	OutCash
	Point
)

type PayloadOperator int32

const (
	AddPayloadOperator PayloadOperator = iota + 1
	XORHighPayloadOperator
)

type Reward struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`
	Order  int32  `json:"order"`

	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	RewardType RewardType `json:"rewardType"`

	PayloadOperator PayloadOperator    `json:"payloadOperator"`
	Payloads        []*payload.Payload `json:"payloads"`
}
