package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
	"example.com/creditcard/stores/constraint"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	uuid "github.com/nu7hatch/gouuid"
)

type impl struct {
	dig.In

	constraintStore constraint.Store
}

func New(
	constraintStore constraint.Store,
) Service {

	im := &impl{
		constraintStore: constraintStore,
	}
	return im
}

func (im *impl) Create(ctx context.Context, constraint *constraintM.Constraint) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}

	constraint.ID = id.String()

	if err := im.constraintStore.Create(ctx, constraint); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*constraintM.Constraint, error) {

	constraint, err := im.constraintStore.GetByID(ctx, ID)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return constraint, nil
}

func (im *impl) GetAll(ctx context.Context) ([]*constraintM.Constraint, error) {

	constraints, err := im.constraintStore.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}
	return constraints, nil
}

func (im *impl) GetByPrivilageID(ctx context.Context, privilageID string) ([]*constraintM.Constraint, error) {

	constraints, err := im.constraintStore.GetByPrivilageID(ctx, privilageID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}
	return constraints, nil
}
