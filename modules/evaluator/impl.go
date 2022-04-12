package evaluator

import (
	"context"
	"errors"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"example.com/creditcard/builder/cardreward"

	cardComp "example.com/creditcard/components/card"

	cardM "example.com/creditcard/models/card"
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
	cardCompnent cardComp.Component
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

	cardResp, err := im.cardService.GetRespByID(ctx, cardID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// rewards, err := im.rewardService.GetByCardID(ctx, cardID)
	// if err != nil {
	// 	logrus.Error(err)
	// 	return err
	// }

	// card.Rewards = rewards

	cardCompnent, err := im.cardBuilder.BuildCardComponent(ctx, cardResp)
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

	cardEventResps := []*cardM.CardEventResp{}

	if len(e.CardIDs) == 0 {
		for _, c := range im.cards {

			cardEventResp, err := im.evaluateCard(ctx, e, c.cardCompnent)
			if err != nil {
				return nil, err
			}
			cardEventResps = append(cardEventResps, cardEventResp)
		}
	} else {
		for _, cardID := range e.CardIDs {

			if c, ok := im.cards[cardID]; ok {
				cardEventResp, err := im.evaluateCard(ctx, e, c.cardCompnent)
				if err != nil {
					return nil, err
				}
				cardEventResps = append(cardEventResps, cardEventResp)

			} else {
				return nil, errors.New("Not found card ID: " + cardID)
			}
		}
	}

	resp.CardEventResps = cardEventResps
	return resp, nil
}

func (im *impl) evaluateCard(ctx context.Context, e *eventM.Event, cardComp cardComp.Component) (*cardM.CardEventResp, error) {

	cardEventResp, err := cardComp.Satisfy(ctx, e)

	if err != nil {
		return nil, err
	}
	return cardEventResp, nil
}
