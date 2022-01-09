package reward

import (
	"context"

	eventM "example.com/creditcard/models/event"
)

type Component interface {
	Satisfy(ctx context.Context, e *eventM.Event) (*eventM.RewardResp, error)
}
