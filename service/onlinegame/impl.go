package onlinegame

import (
	"context"

	onlinegameM "example.com/creditcard/models/onlinegame"
	onlinegameStore "example.com/creditcard/stores/onlinegame"
	"github.com/sirupsen/logrus"

	uuid "github.com/nu7hatch/gouuid"
)

type impl struct {
	onlinegameStore onlinegameStore.Store
}

func New(
	onlinegameStore onlinegameStore.Store,
) Service {
	return &impl{
		onlinegameStore: onlinegameStore,
	}
}

func (im *impl) Create(ctx context.Context, onlinegame *onlinegameM.Onlinegame) error {

	ID, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}
	onlinegame.ID = ID.String()

	if err := im.onlinegameStore.Create(ctx, onlinegame); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) UpdateByID(ctx context.Context, onlinegame *onlinegameM.Onlinegame) error {

	if err := im.onlinegameStore.UpdateByID(ctx, onlinegame); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*onlinegameM.Onlinegame, error) {
	onlinegames, err := im.onlinegameStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return onlinegames, nil
}
