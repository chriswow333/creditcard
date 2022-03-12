package customization

import (
	"context"

	customizationM "example.com/creditcard/models/customization"
	customizationStore "example.com/creditcard/stores/customization"
)

type impl struct {
	customizationStore customizationStore.Store
}

func New(
	customizationStore customizationStore.Store,
) Service {
	return &impl{
		customizationStore: customizationStore,
	}
}

func (im *impl) GetByCardID(ctx context.Context, rewardID string) ([]*customizationM.Customization, error) {

	customizations, err := im.customizationStore.GetByCardID(ctx, rewardID)
	if err != nil {
		return nil, err
	}
	return customizations, nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*customizationM.Customization, error) {
	customization, err := im.customizationStore.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return customization, nil
}
