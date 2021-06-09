package privilage

import (
	"context"
	"time"

	privilageM "example.com/creditcard/models/privilage"
	"example.com/creditcard/stores/privilage"
	"github.com/sirupsen/logrus"

	"go.uber.org/dig"

	uuid "github.com/nu7hatch/gouuid"
)

var (
	timeNow = time.Now
)

type impl struct {
	dig.In

	privilageStore privilage.Store
}

func New(
	privilageStore privilage.Store,
) Service {
	return &impl{
		privilageStore: privilageStore,
	}
}

func (im *impl) Create(ctx context.Context, privilage *privilageM.Privilage) error {

	privilage.UpdateDate = timeNow().Unix()

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	privilage.ID = id.String()

	if err := im.privilageStore.Create(ctx, privilage); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*privilageM.Privilage, error) {
	privilage, err := im.privilageStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return privilage, nil
}

func (im *impl) GetAll(ctx context.Context) ([]*privilageM.Privilage, error) {
	privilages, err := im.privilageStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return privilages, nil
}

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*privilageM.Privilage, error) {
	privilages, err := im.privilageStore.GetByCardID(ctx, cardID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return privilages, nil
}
