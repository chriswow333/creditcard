package streaming

import (
	"context"

	streamingM "example.com/creditcard/models/streaming"
	streamingStore "example.com/creditcard/stores/streaming"
	"github.com/sirupsen/logrus"

	uuid "github.com/nu7hatch/gouuid"
)

type impl struct {
	streamingStore streamingStore.Store
}

func New(
	streamingStore streamingStore.Store,
) Service {
	return &impl{
		streamingStore: streamingStore,
	}
}

func (im *impl) Create(ctx context.Context, streaming *streamingM.Streaming) error {

	ID, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}
	streaming.ID = ID.String()

	if err := im.streamingStore.Create(ctx, streaming); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) UpdateByID(ctx context.Context, streaming *streamingM.Streaming) error {
	if err := im.streamingStore.UpdateByID(ctx, streaming); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*streamingM.Streaming, error) {
	streamings, err := im.streamingStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return streamings, nil
}
