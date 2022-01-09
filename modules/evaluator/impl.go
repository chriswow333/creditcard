package evaluator

import (
	"context"

	"go.uber.org/dig"

	"github.com/sirupsen/logrus"

	"example.com/creditcard/builder/cardreward"
	cardComp "example.com/creditcard/components/card"
	eventM "example.com/creditcard/models/event"
	cardService "example.com/creditcard/service/card"
	rewardService "example.com/creditcard/service/reward"
)

type impl struct {
	dig.In

	cards         map[string]*cardEvaluator
	cardService   cardService.Service
	rewardService rewardService.Service
	cardBuilder   cardreward.Builder
}

type cardEvaluator struct {
	ID           string
	cardCompnent *cardComp.Component
}

func New(
	cardService cardService.Service,
	rewardService rewardService.Service,
	cardBuilder cardreward.Builder,
) Module {
	im := &impl{
		cards:         make(map[string]*cardEvaluator),
		cardService:   cardService,
		rewardService: rewardService,
		cardBuilder:   cardBuilder,
	}

	// init the card component
	if err := im.UpdateAllComponents(context.Background()); err != nil {
		logrus.Error(err)
	}

	return im
}

func (im *impl) UpdateAllComponents(ctx context.Context) error {
	cards, err := im.cardService.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, card := range cards {
		if err := im.UpdateComponentByCardID(ctx, card.ID); err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}

func (im *impl) UpdateComponentByCardID(ctx context.Context, cardID string) error {

	card, err := im.cardService.GetByID(ctx, cardID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	rewards, err := im.rewardService.GetByCardID(ctx, cardID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	card.Rewards = rewards

	cardCompnent, err := im.cardBuilder.BuildCardComponent(ctx, card)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	im.cards[cardID] = &cardEvaluator{
		cardCompnent: cardCompnent,
	}
	return nil
}

func (im *impl) Evaluate(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		EventID: e.ID,
	}

	specificedCardID := make(map[string]bool)
	for _, c := range e.CardIDs {
		specificedCardID[c] = true
	}

	cards := []*eventM.CardResp{}

	for _, c := range im.cards {
		if len(e.CardIDs) != 0 {
			if _, ok := specificedCardID[c.ID]; ok {
				card, err := im.evaluateCard(ctx, e, *c.cardCompnent)
				if err != nil {
					return nil, err
				}
				cards = append(cards, card)
			}
		} else {
			card, err := im.evaluateCard(ctx, e, *c.cardCompnent)
			if err != nil {
				return nil, err
			}
			cards = append(cards, card)
		}
	}

	resp.Cards = cards
	return resp, nil
}

func (im *impl) evaluateCard(ctx context.Context, e *eventM.Event, cardComp cardComp.Component) (*eventM.CardResp, error) {

	card, err := cardComp.Satisfy(ctx, e)
	if err != nil {
		return nil, err
	}
	return card, nil
}
