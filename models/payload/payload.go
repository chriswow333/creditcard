package payload

import (
	"example.com/creditcard/models/constraint"
	"example.com/creditcard/models/feedback"
)

type Payload struct {
	Descs []string `json:"descs"`

	Feedback   *feedback.Feedback     `json:"feedback"`
	Constraint *constraint.Constraint `json:"constraint"`
}

type PayloadResp struct {
	Descs []string `json:"descs"`

	Pass bool `json:"pass"`

	Feedback *feedback.Feedback `json:"feedback"`

	FeedbackResp *feedback.FeedbackResp `json:"feedbackResp"`
	FeedReturn   *feedback.FeedReturn   `json:"feedReturn"`

	ConstraintResp *constraint.ConstraintResp `json:"constraintResp"`

	Constraint *constraint.Constraint `json:"constraint"`
}

func TransferPayloadResp(payload *Payload) *PayloadResp {

	payloadResp := &PayloadResp{
		Descs:        payload.Descs,
		FeedbackResp: feedback.TransferFeedbackResp(payload.Feedback),
		Constraint:   payload.Constraint,
	}

	return payloadResp
}
