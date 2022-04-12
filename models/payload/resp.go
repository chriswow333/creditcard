package payload

import (
	"example.com/creditcard/models/constraint"
	"example.com/creditcard/models/feedback"
)

type PayloadResp struct {
	ID    string   `json:"id"`
	Descs []string `json:"descs"`

	Feedback       *feedback.Feedback         `json:"feedback,omitempty"`
	ConstraintResp *constraint.ConstraintResp `json:"constraintResp,omitempty"`
}
