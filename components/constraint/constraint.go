package constraint

import (
	"example.com/creditcard/models/event"
)

type Constraint interface {
	Validate(event.Event)
}
