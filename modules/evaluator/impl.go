package evaluator

import (
	"context"

	cardComp "example.com/creditcard/components/card"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	cards []*cardComp.Component
}

func New(
	cards []*cardComp.Component,
) Module {
	return &impl{
		cards: cards,
	}
}

func (im *impl) Evaluate(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	eventResp := &eventM.Response{}
	for _, c := range im.cards {
		_, err := (*c).Satisfy(ctx, e)
		if err != nil {
			return nil, err
		}

	}
	return eventResp, nil
}
