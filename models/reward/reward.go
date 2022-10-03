package reward

import (
	"example.com/creditcard/models/payload"
)

type RewardType int32

const (
	CASH_TWD RewardType = iota + 1
	LINE_POINT
	KUO_BROTHERS_POINT // 生活市集幣
	WOWPRIME_POINT     // 王品瘋點數
	OPEN_POINT         // OPEN POINT

	RED_POINT // 紅利點數
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
