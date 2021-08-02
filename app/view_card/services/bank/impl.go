package bank

import (
	"context"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	bankM "example.com/creditcard/app/view_card/models/bank"
	"example.com/creditcard/app/view_card/stores/bank"
)

var (
	timeNow = time.Now
)

type impl struct {
	dig.In

	bankStore bank.Store
}

func New(
	bankStore bank.Store,
) Service {
	return &impl{
		bankStore: bankStore,
	}
}

func (im *impl) Create(ctx context.Context, Repr *bankM.Repr) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	bank := &bankM.Bank{
		ID:         id.String(),
		Name:       Repr.Name,
		Icon:       Repr.Icon,
		UpdateDate: timeNow().Unix(),
	}

	if err := im.bankStore.Create(ctx, bank); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*bankM.Repr, error) {
	bank, err := im.bankStore.GetByID(ctx, ID)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	Repr := &bankM.Repr{
		ID:   bank.ID,
		Name: bank.Name,
		Icon: bank.Icon,
	}

	return Repr, nil
}

func (im *impl) UpdateByID(ctx context.Context, Repr *bankM.Repr) error {

	bank := &bankM.Bank{
		ID:         Repr.ID,
		Name:       Repr.Name,
		Icon:       Repr.Icon,
		UpdateDate: timeNow().Unix(),
	}

	if err := im.bankStore.UpdateByID(ctx, bank); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*bankM.Repr, error) {

	banks, err := im.bankStore.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	bankReprs := []*bankM.Repr{}

	for _, b := range banks {
		Repr := &bankM.Repr{
			ID:   b.ID,
			Name: b.Name,
			Icon: b.Icon,
		}
		bankReprs = append(bankReprs, Repr)
	}

	return bankReprs, nil
}
