package customization

import (
	"context"

	customizationM "example.com/creditcard/models/customization"
	customizationStore "example.com/creditcard/stores/customization"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	uuid "github.com/nu7hatch/gouuid"
)

type impl struct {
	dig.In

	customizationStore customizationStore.Store
}

func New(
	customizationStore customizationStore.Store,
) Service {
	return &impl{
		customizationStore: customizationStore,
	}
}

func (im *impl) Create(ctx context.Context, customization *customizationM.Customization) error {

	ID, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	customization.ID = ID.String()

	if err := im.customizationStore.Create(ctx, customization); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*customizationM.Customization, error) {
	customiztion, err := im.customizationStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return customiztion, nil
}

func (im *impl) UpdateByID(ctx context.Context, customization *customizationM.Customization) error {
	if err := im.customizationStore.UpdateByID(ctx, customization); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetByRewardID(ctx context.Context, rewardID string) ([]*customizationM.Customization, error) {
	customizations, err := im.GetByRewardID(ctx, rewardID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return customizations, nil
}
