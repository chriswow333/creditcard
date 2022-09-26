package payload

import (
	"example.com/creditcard/models/channel"
	"example.com/creditcard/models/feedback"
)

type PayloadEventJudgeType int32

const (
	ALL PayloadEventJudgeType = iota + 1
	SOME
	NONE
)

type PayloadEventResp struct {
	ID string `json:"id"`

	PayloadEventJudgeType PayloadEventJudgeType `json:"payloadEventJudgeType,omitempty"`

	FeedReturn *feedback.FeedReturn `json:"feedReturn,omitempty"`

	ConstraintEventResp *channel.ChannelEventResp `json:"channelEventResp,omitempty"`
}
