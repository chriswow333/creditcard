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
