package channel

import (
	"context"

	"example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
)

type Component interface {
	Judge(ctx context.Context, e *eventM.Event) (*channel.ChannelEventResp, error)
}
