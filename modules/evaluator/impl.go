package evaluator

import (
	"context"

	"example.com/creditcard/builder/cardreward"
	cardComp "example.com/creditcard/components/card"
	eventM "example.com/creditcard/models/event"
	cardService "example.com/creditcard/service/card"
	constraintService "example.com/creditcard/service/constraint"
	rewardService "example.com/creditcard/service/reward"
)

type impl struct {
	cards             map[string]*cardEvaluator
	constraintService constraintService.Service
	cardService       cardService.Service
	rewardService     rewardService.Service
	cardBuilder       cardreward.Builder
}

type cardEvaluator struct {
	cardCompnent *cardComp.Component
}

func New(
	constraintService constraintService.Service,
	cardService cardService.Service,
	rewardService rewardService.Service,
	cardBuilder cardreward.Builder,
) Module {
	im := &impl{
		constraintService: constraintService,
		cardService:       cardService,
		rewardService:     rewardService,
		cardBuilder:       cardBuilder,
	}

	// init the card component
	if err := im.UpdateAllComponents(context.Background()); err != nil {
		panic(err)
	}

	return im
}

func (im *impl) UpdateAllComponents(ctx context.Context) error {
	cards, err := im.cardService.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, card := range cards {
		if err := im.UpdateComponentByCardID(ctx, card.ID); err != nil {
			return err
		}
	}

	return nil
}

func (im *impl) UpdateComponentByCardID(ctx context.Context, cardID string) error {

	card, err := im.cardService.GetByID(ctx, cardID)
	if err != nil {
		return err
	}

	cardCompnent, err := im.cardBuilder.BuildCardComponent(ctx, card)
	if err != nil {
		return nil
	}

	im.cards[cardID] = &cardEvaluator{
		cardCompnent: cardCompnent,
	}

	return nil
}

func (im *impl) Evaluate(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	return nil, nil
}
