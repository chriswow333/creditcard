package cardreward

import (
	"context"

	cardComp "example.com/creditcard/components/card"
	cardM "example.com/creditcard/models/card"
)

type Builder interface {
	BuildCardComponent(ctx context.Context, setting *cardM.Card) (cardComp.Component, error)
}
