package streaming

import (
	"context"

	streamingM "example.com/creditcard/models/streaming"
)

type Service interface {
	Create(ctx context.Context, streaming *streamingM.Streaming) error
	UpdateByID(ctx context.Context, streaming *streamingM.Streaming) error
	GetAll(ctx context.Context) ([]*streamingM.Streaming, error)
}
