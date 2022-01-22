package customization

import (
	"context"

	customizationM "example.com/creditcard/models/customization"
)

type Store interface {
	Create(ctx context.Context, customization *customizationM.Customization) error
	GetByID(ctx context.Context, ID string) (*customizationM.Customization, error)
	UpdateByID(ctx context.Context, customization *customizationM.Customization) error
}
