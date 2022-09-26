package reward

import (
	"example.com/creditcard/models/feedback"
	"example.com/creditcard/models/payload"
)

type RewardEventJudgeType int32

const (
	ALL RewardEventJudgeType = iota + 1
	SOME
	NONE
)

type RewardEventResp struct {
	ID           string `json:"id"`
	CardRewardID string `json:"cardRewardID"`

	Order int32 `json:"order"`

	RewardEventJudgeType RewardEventJudgeType `json:"rewardEventJudgeType,omitempty"`

	PayloadOperator PayloadOperator `json:"payloadOperator,omitempty"`

	FeedReturn *feedback.FeedReturn `json:"feedReturn,omitempty"`

	PayloadEventResps []*payload.PayloadEventResp `json:"payloadEventResps,omitempty"`
}
