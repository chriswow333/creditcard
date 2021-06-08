package gogoconstraint

import (
	eventM "example.com/creditcard/models/event"
)

type impl struct {
}

func (im *impl) Validate(event eventM.Event) (bool, error) {
	return false, nil
}
