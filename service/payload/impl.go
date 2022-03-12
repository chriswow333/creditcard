package payload

import (
	"context"
	"time"

	payloadM "example.com/creditcard/models/payload"
	"example.com/creditcard/stores/payload"
	"github.com/sirupsen/logrus"

	"go.uber.org/dig"
)

var (
	timeNow = time.Now
)

type impl struct {
	dig.In

	payloadStore payload.Store
}

func New(
	payloadStore payload.Store,
) Service {
	return &impl{
		payloadStore: payloadStore,
	}
}

func (im *impl) UpdateByID(ctx context.Context, ID string, paylaods []*payloadM.Payload) error {
	if err := im.payloadStore.UpdateByID(ctx, ID, paylaods); err != nil {
		logrus.Error(err)
		return err
	}
	return nil

}
