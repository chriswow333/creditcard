package payload

import (
	"context"

	payloadM "example.com/creditcard/models/payload"
)

type Service interface {
	UpdateByID(ctx context.Context, ID string, payloads []*payloadM.Payload) error
}
