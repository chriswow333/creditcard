package creditcardreward

import (
	"context"

	bankComp "example.com/creditcard/components/bank"
	bankM "example.com/creditcard/models/bank"
)

type Builder interface {
	NewCreditcard(ctx context.Context, settings []*bankM.Bank) ([]*bankComp.Component, error)
}
