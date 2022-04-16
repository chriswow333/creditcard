package payload

import (
	"example.com/creditcard/models/constraint"
	"example.com/creditcard/models/feedback"
)

type Payload struct {
	ID    string   `json:"id"`
	Descs []string `json:"descs"`

	Feedback   *feedback.Feedback     `json:"feedback,omitempty"`
	Constraint *constraint.Constraint `json:"constraint,omitempty"`
}

type PayloadEventJudgeType int32

const (
	ALL PayloadEventJudgeType = iota + 1
	SOME
	NONE
)

func TransferPayloadResp(payload *Payload) *PayloadResp {

	return &PayloadResp{
		ID:       payload.ID,
		Descs:    payload.Descs,
		Feedback: payload.Feedback,
	}
}
