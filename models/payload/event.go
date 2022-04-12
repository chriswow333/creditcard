package payload

import (
	"example.com/creditcard/models/constraint"
	"example.com/creditcard/models/feedback"
)

type PayloadEventResp struct {
	ID string `json:"id"`

	PayloadEventJudgeType PayloadEventJudgeType `json:"payloadEventJudgeType,omitempty"`

	FeedReturn *feedback.FeedReturn `json:"feedReturn,omitempty"`

	ConstraintEventResp *constraint.ConstraintEventResp `json:"constraintEventResp,omitempty"`
}
