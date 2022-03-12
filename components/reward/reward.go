package reward

import (
	"context"

	eventM "example.com/creditcard/models/event"

	rewardM "example.com/creditcard/models/reward"
)

type Component interface {
	Satisfy(ctx context.Context, e *eventM.Event) (*rewardM.RewardResp, error)
}
