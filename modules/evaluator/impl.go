package evaluator

import (
	"context"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

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

	if e.ID == "" {
		id, err := uuid.NewV4()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": "",
			}).Fatal(err)

			return nil, err
		}
		e.ID = id.String()
	}

	resp := &eventM.Response{
		EventID: e.ID,
	}

	cards := []*eventM.CardResp{}

	for _, c := range im.cards {
		card, err := im.evaluateCard(ctx, e, *c.cardCompnent)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
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
