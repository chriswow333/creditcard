package customization

import (
	"context"

	customizationM "example.com/creditcard/models/customization"
)

type Service interface {
	GetByCardID(ctx context.Context, rewardID string) ([]*customizationM.Customization, error)
	GetByID(ctx context.Context, ID string) (*customizationM.Customization, error)
}
