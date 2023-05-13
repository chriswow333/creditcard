package evaluator

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sort"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"example.com/creditcard/builder/cardreward"

	cardComp "example.com/creditcard/components/card"

	cardM "example.com/creditcard/models/card"
	"example.com/creditcard/models/event"
	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"

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
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
	}

	return im
}

func (im *impl) UpdateAllComponents(ctx context.Context) error {

	cards, err := im.cardService.GetAll(ctx)

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	for _, card := range cards {
		logrus.Infof("get card id %s", card.ID)
		if card.CardStatus != 1 {
			continue
		}

		if err := im.UpdateComponentByCardID(ctx, card.ID); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return err
		}
	}

	return nil
}

func (im *impl) UpdateComponentByCardID(ctx context.Context, cardID string) error {

	card, err := im.cardService.GetByID(ctx, cardID)

	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return err
	}

	cardCompnent, err := im.cardBuilder.BuildCardComponent(ctx, card)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
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
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		e.ID = id.String()
	}

	resp := &eventM.Response{
		EventID: e.ID,
	}

	cardEventResps := []*cardM.CardEventResp{}

	if len(e.CardIDs) == 0 {
		// eavalute all cards
		for _, c := range im.cards {

			cardEventResp, err := im.evaluateCard(ctx, e, c.cardCompnent)
			if err != nil {
				logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
				return nil, err
			}

			if len(cardEventResp.CardRewardEventResps) != 0 {
				cardEventResps = append(cardEventResps, cardEventResp)
			}

		}
		resp.CardEventResps = im.sortEvaluatedCardResults(ctx, e, cardEventResps)

	} else {

		for _, cardID := range e.CardIDs {

			if c, ok := im.cards[cardID]; ok {

				cardEventResp, err := im.evaluateCard(ctx, e, c.cardCompnent)

				if err != nil {
					logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
					return nil, err
				}

				cardEventResps = append(cardEventResps, cardEventResp)

			} else {
				return nil, errors.New("Not found card ID: " + cardID)
			}
		}

		resp.CardEventResps = cardEventResps

	}

	return resp, nil
}

func (im *impl) evaluateCard(ctx context.Context, e *eventM.Event, cardComp cardComp.Component) (*cardM.CardEventResp, error) {

	cardEventResp, err := cardComp.Satisfy(ctx, e)

	if err != nil {
		return nil, err
	}
	return cardEventResp, nil
}

func (im *impl) sortEvaluatedCardResults(ctx context.Context, e *eventM.Event, cardEventResps []*cardM.CardEventResp) []*cardM.CardEventResp {

	logrus.Info("evaluator.sortEvaluatedCardResults, sortType: ", e.SortType)

	trimCardEventResps := []*cardM.CardEventResp{}
	fmt.Println(e.SortType)

	switch e.SortType {
	case eventM.NONE:
		trimCardEventResps = cardEventResps
		break
	case eventM.MATCH:
		trimCardEventResps = im.sortMatchEvaluateCardResults(ctx, e, cardEventResps)
		break
	case event.MAX_REWARD_BONUS:
		trimCardEventResps = im.sortMaxRewardBonusEvaluatedCardResults(ctx, e, cardEventResps)
		break
	case event.MAX_REWARD_RETURN:
		trimCardEventResps = im.sortMaxRewardReturnEvaluatedCardResults(ctx, e, cardEventResps)
		break
	case event.MAX_REWARD_EXPECTED_BONUS:
		trimCardEventResps = im.sortMaxRewardReturnExpectedCardResults(ctx, e, cardEventResps)
		break
	default:
		trimCardEventResps = cardEventResps
	}

	maxSize := 10

	if len(trimCardEventResps) < maxSize {
		maxSize = len(trimCardEventResps)
	}

	return trimCardEventResps[0:maxSize]

}

func (im *impl) sortMaxRewardReturnExpectedCardResults(ctx context.Context, e *eventM.Event, cardEventResps []*cardM.CardEventResp) []*cardM.CardEventResp {

	logrus.Info("evaluator.sortMaxRewardReturnExpectedCardResults")

	for _, cardEventResp := range cardEventResps {

		separatedCardRewardEventResps := []*cardM.CardRewardEventResp{}
		multiplyBonusRewardEventResps := []*cardM.CardRewardEventResp{}
		fixedBonusRewardEventResps := []*cardM.CardRewardEventResp{}

		for _, cardRewardEventResp := range cardEventResp.CardRewardEventResps {

			rewardType := cardRewardEventResp.RewardType

			switch rewardType {
			case rewardM.CASH:

				cashCalculateType := cardRewardEventResp.FeedbackBonus.CashFeedbackBonus.CashCalculateType

				switch cashCalculateType {
				case feedbackM.BONUS_MULTIPLY_CASH:
					multiplyBonusRewardEventResps = append(multiplyBonusRewardEventResps, cardRewardEventResp)
					break
				case feedbackM.FIXED_CASH_RETURN:
					fixedBonusRewardEventResps = append(fixedBonusRewardEventResps, cardRewardEventResp)
					break
				}

				break

			case rewardM.POINT:

				pointCalculateType := cardRewardEventResp.FeedbackBonus.PointFeedbackBonus.PointCalculateType

				switch pointCalculateType {
				case feedbackM.BONUS_MULTIPLY_POINT:
					multiplyBonusRewardEventResps = append(multiplyBonusRewardEventResps, cardRewardEventResp)
					break
				case feedbackM.FIXED_POINT_RETURN:
					fixedBonusRewardEventResps = append(fixedBonusRewardEventResps, cardRewardEventResp)
					break
				}
				break
			}
		}

		separatedCardRewardEventResps = append(separatedCardRewardEventResps, multiplyBonusRewardEventResps...)
		separatedCardRewardEventResps = append(separatedCardRewardEventResps, fixedBonusRewardEventResps...)

		cardEventResp.CardRewardEventResps = separatedCardRewardEventResps
	}

	sort.SliceStable(cardEventResps, func(i, j int) bool {

		firstRewardType := cardEventResps[i].CardRewardEventResps[0].RewardType
		firstTotalBonus := 0.0

		switch firstRewardType {
		case rewardM.CASH:
			firstTotalBonus = cardEventResps[i].CardRewardEventResps[0].FeedbackBonus.CashFeedbackBonus.TotalBonus
			break
		case rewardM.POINT:
			firstTotalBonus = cardEventResps[i].CardRewardEventResps[0].FeedbackBonus.PointFeedbackBonus.TotalBonus
			break
		}

		secondRewardType := cardEventResps[j].CardRewardEventResps[0].RewardType
		secondTotalBonus := 0.0
		switch secondRewardType {
		case rewardM.CASH:
			secondTotalBonus = cardEventResps[j].CardRewardEventResps[0].FeedbackBonus.CashFeedbackBonus.TotalBonus
			break
		case rewardM.POINT:
			secondTotalBonus = cardEventResps[j].CardRewardEventResps[0].FeedbackBonus.PointFeedbackBonus.TotalBonus
			break
		}

		return firstTotalBonus > secondTotalBonus
	})

	return cardEventResps

}

// func (im *impl) getTotalBonus(cardRewardEventResp *cardM.CardRewardEventResp) float64 {

// 	rewardType := cardRewardEventResp.RewardType
// 	switch rewardType {
// 	case rewardM.CASH:
// 		return cardRewardEventResp.FeedbackBonus.PointFeedbackBonus.TotalBonus

// 	case rewardM.POINT:
// 		return cardRewardEventResp.FeedbackBonus.CashFeedbackBonus.TotalBonus

// 	default:
// 		return 0.0
// 	}

// }

func (im *impl) sortMaxRewardReturnEvaluatedCardResults(ctx context.Context, e *eventM.Event, cardEventResps []*cardM.CardEventResp) []*cardM.CardEventResp {

	logrus.Info("evaluator.sortMaxRewardReturnEvaluatedCardResults")

	// cashCardEventResps := []*cardM.CardEventResp{}

	// separate...
	for _, cardEventResp := range cardEventResps {

		// take the best reward in a card.
		maxCardRewardEventResps := []*cardM.CardRewardEventResp{}
		firstRewardType := rewardM.CASH

		rewardType := cardEventResp.CardRewardEventResps[0].RewardType
		maxCardRewardEventResp := cardEventResp.CardRewardEventResps[0]

		for _, cardRewardEventResp := range cardEventResp.CardRewardEventResps {

			switch rewardType {
			case rewardM.CASH:

				if len(maxCardRewardEventResps) == 0 {
					maxCardRewardEventResp = cardRewardEventResp
					firstRewardType = rewardM.CASH
					break
				}

				switch firstRewardType {
				case rewardM.CASH:

					if maxCardRewardEventResp.FeedReturn.CashReturn.ActualCashReturn < cardRewardEventResp.FeedReturn.CashReturn.ActualCashReturn {
						maxCardRewardEventResp = cardRewardEventResp
						firstRewardType = rewardM.CASH
					}
					break

				case rewardM.POINT:
					// 這裡目前還是先用1:1 point & cash
					if maxCardRewardEventResp.FeedReturn.PointReturn.ActualPointReturn < cardRewardEventResp.FeedReturn.CashReturn.ActualCashReturn {
						maxCardRewardEventResp = cardRewardEventResp
						firstRewardType = rewardM.CASH
					}
					break
				}

				break
			case rewardM.POINT:

				if len(maxCardRewardEventResps) == 0 {
					maxCardRewardEventResp = cardRewardEventResp
					firstRewardType = rewardM.POINT
					break
				}

				switch firstRewardType {
				case rewardM.CASH:
					if maxCardRewardEventResp.FeedReturn.CashReturn.ActualCashReturn < cardRewardEventResp.FeedReturn.PointReturn.ActualPointReturn {
						maxCardRewardEventResp = cardRewardEventResp
						firstRewardType = rewardM.POINT
					}
					break

				case rewardM.POINT:
					// 這裡目前還是先用1:1 point vs cash
					if maxCardRewardEventResp.FeedReturn.PointReturn.ActualPointReturn < cardRewardEventResp.FeedReturn.PointReturn.ActualPointReturn {
						maxCardRewardEventResp = cardRewardEventResp
						firstRewardType = rewardM.POINT
					}
					break
				}
				break
			}

		}
		cardEventResp.CardRewardEventResps = []*cardM.CardRewardEventResp{maxCardRewardEventResp}
	}

	// sorting...

	sort.SliceStable(cardEventResps, func(i, j int) bool {

		firstRewardType := cardEventResps[i].CardRewardEventResps[0].RewardType
		firstActualReturn := 0.0

		switch firstRewardType {
		case rewardM.CASH:
			firstActualReturn = cardEventResps[i].CardRewardEventResps[0].FeedReturn.CashReturn.ActualCashReturn
			break
		case rewardM.POINT:
			firstActualReturn = cardEventResps[i].CardRewardEventResps[0].FeedReturn.PointReturn.ActualPointReturn
			break
		}

		secondRewardType := cardEventResps[j].CardRewardEventResps[0].RewardType
		secondActualReturn := 0.0
		switch secondRewardType {
		case rewardM.CASH:
			secondActualReturn = cardEventResps[j].CardRewardEventResps[0].FeedReturn.CashReturn.ActualCashReturn
			break
		case rewardM.POINT:
			secondActualReturn = cardEventResps[j].CardRewardEventResps[0].FeedReturn.PointReturn.ActualPointReturn
			break
		}

		return firstActualReturn > secondActualReturn

	})

	return cardEventResps
}

func (im *impl) sortMaxRewardBonusEvaluatedCardResults(ctx context.Context, e *eventM.Event, cardEventResps []*cardM.CardEventResp) []*cardM.CardEventResp {

	logrus.Info("evaluator.sortMaxRewardEvaluatedCardResults")

	// separate...
	for _, cardEventResp := range cardEventResps {

		firstRewardType := rewardM.CASH

		firstMaxCardRewardEventResps := []*cardM.CardRewardEventResp{}

		for _, cardRewardEventResp := range cardEventResp.CardRewardEventResps {

			rewardType := cardRewardEventResp.RewardType

			if len(firstMaxCardRewardEventResps) == 0 {
				firstMaxCardRewardEventResps = append(firstMaxCardRewardEventResps, cardRewardEventResp)
				firstRewardType = cardRewardEventResp.RewardType
				continue
			}

			switch rewardType {
			case rewardM.CASH:
				returnBonus := cardRewardEventResp.FeedReturn.CashReturn.CashReturnBonus

				switch firstRewardType {
				case rewardM.CASH:
					if firstMaxCardRewardEventResps[0].FeedReturn.CashReturn.CashReturnBonus < returnBonus {
						firstMaxCardRewardEventResps = append([]*cardM.CardRewardEventResp{cardRewardEventResp}, firstMaxCardRewardEventResps...)
						firstRewardType = rewardType
					} else {
						firstMaxCardRewardEventResps = append(firstMaxCardRewardEventResps, cardRewardEventResp)
					}
					break
				case rewardM.POINT:
					if firstMaxCardRewardEventResps[0].FeedReturn.PointReturn.PointReturnBonus < returnBonus {
						firstMaxCardRewardEventResps = append([]*cardM.CardRewardEventResp{cardRewardEventResp}, firstMaxCardRewardEventResps...)
						firstRewardType = rewardType
					} else {
						firstMaxCardRewardEventResps = append(firstMaxCardRewardEventResps, cardRewardEventResp)
					}
					break
				}

				break
			case rewardM.POINT:
				returnBonus := cardRewardEventResp.FeedReturn.PointReturn.PointReturnBonus

				switch firstRewardType {
				case rewardM.CASH:
					if firstMaxCardRewardEventResps[0].FeedReturn.CashReturn.CashReturnBonus < returnBonus {
						firstMaxCardRewardEventResps = append([]*cardM.CardRewardEventResp{cardRewardEventResp}, firstMaxCardRewardEventResps...)
						firstRewardType = rewardType
					} else {
						firstMaxCardRewardEventResps = append(firstMaxCardRewardEventResps, cardRewardEventResp)
					}
					break
				case rewardM.POINT:
					if firstMaxCardRewardEventResps[0].FeedReturn.PointReturn.PointReturnBonus < returnBonus {
						firstMaxCardRewardEventResps = append([]*cardM.CardRewardEventResp{cardRewardEventResp}, firstMaxCardRewardEventResps...)
						firstRewardType = rewardType
					} else {
						firstMaxCardRewardEventResps = append(firstMaxCardRewardEventResps, cardRewardEventResp)
					}
					break
				}
				break
			}

		}

		cardEventResp.CardRewardEventResps = firstMaxCardRewardEventResps

	}

	// sorting...

	sort.SliceStable(cardEventResps, func(i, j int) bool {

		firstRewardType := cardEventResps[i].CardRewardEventResps[0].RewardType

		firstReturnBonus := 0.0

		switch firstRewardType {
		case rewardM.CASH:
			firstReturnBonus = cardEventResps[i].CardRewardEventResps[0].FeedReturn.CashReturn.CashReturnBonus
			break
		case rewardM.POINT:
			firstReturnBonus = cardEventResps[i].CardRewardEventResps[0].FeedReturn.PointReturn.PointReturnBonus
			break
		}

		secondRewardType := cardEventResps[j].CardRewardEventResps[0].RewardType
		secondReturnBonus := 0.0
		switch secondRewardType {
		case rewardM.CASH:
			secondReturnBonus = cardEventResps[j].CardRewardEventResps[0].FeedReturn.CashReturn.CashReturnBonus
			break
		case rewardM.POINT:
			secondReturnBonus = cardEventResps[j].CardRewardEventResps[0].FeedReturn.PointReturn.PointReturnBonus
			break
		}

		return firstReturnBonus > secondReturnBonus
	})

	return cardEventResps
}

// remove mismatch reward parts except specific card reward.
func (im *impl) sortMatchEvaluateCardResults(ctx context.Context, e *eventM.Event, cardEventResps []*cardM.CardEventResp) []*cardM.CardEventResp {
	logrus.Info("evaluator.sortMatchEvaluateCardResults")

	matchedCardEventResps := make(map[string]*cardM.CardEventResp)

	allMatchedCardRewardEventMapper := make(map[string][]*cardM.CardRewardEventResp)
	someMatchedCardRewardEventMapper := make(map[string][]*cardM.CardRewardEventResp)

	for _, cardEventResp := range cardEventResps {

		for _, cardRewardEventResp := range cardEventResp.CardRewardEventResps {

			switch cardRewardEventResp.RewardType {
			case rewardM.CASH:

				switch cardRewardEventResp.FeedReturn.CashReturn.CashReturnStatus {
				case feedbackM.ALL_RETURN_CASH:
					if _, ok := allMatchedCardRewardEventMapper[cardEventResp.ID]; !ok {
						allMatchedCardRewardEventMapper[cardEventResp.ID] = []*cardM.CardRewardEventResp{}
					}
					allMatchedCardRewardEventMapper[cardEventResp.ID] = append(allMatchedCardRewardEventMapper[cardEventResp.ID], cardRewardEventResp)

					if _, ok := matchedCardEventResps[cardEventResp.ID]; !ok {
						matchedCardEventResps[cardEventResp.ID] = cardEventResp
					}
					break
				case feedbackM.SOME_RETURN_CASH:
					if _, ok := allMatchedCardRewardEventMapper[cardEventResp.ID]; !ok {
						someMatchedCardRewardEventMapper[cardEventResp.ID] = []*cardM.CardRewardEventResp{}
					}
					someMatchedCardRewardEventMapper[cardEventResp.ID] = append(someMatchedCardRewardEventMapper[cardEventResp.ID], cardRewardEventResp)

					if _, ok := matchedCardEventResps[cardEventResp.ID]; !ok {
						matchedCardEventResps[cardEventResp.ID] = cardEventResp
					}
					break
				}

				break

			case rewardM.POINT:
				switch cardRewardEventResp.FeedReturn.PointReturn.PointReturnStatus {
				case feedbackM.ALL_RETURN_POINT:
					if _, ok := allMatchedCardRewardEventMapper[cardEventResp.ID]; !ok {
						allMatchedCardRewardEventMapper[cardEventResp.ID] = []*cardM.CardRewardEventResp{}
					}
					allMatchedCardRewardEventMapper[cardEventResp.ID] = append(allMatchedCardRewardEventMapper[cardEventResp.ID], cardRewardEventResp)

					if _, ok := matchedCardEventResps[cardEventResp.ID]; !ok {
						matchedCardEventResps[cardEventResp.ID] = cardEventResp
					}
					break
				case feedbackM.SOME_RETURN_POINT:
					if _, ok := allMatchedCardRewardEventMapper[cardEventResp.ID]; !ok {
						someMatchedCardRewardEventMapper[cardEventResp.ID] = []*cardM.CardRewardEventResp{}
					}

					someMatchedCardRewardEventMapper[cardEventResp.ID] = append(someMatchedCardRewardEventMapper[cardEventResp.ID], cardRewardEventResp)

					if _, ok := matchedCardEventResps[cardEventResp.ID]; !ok {
						matchedCardEventResps[cardEventResp.ID] = cardEventResp
					}

					break
				}
				break
			default:
				continue
			}
		}
	}

	sortedCardEventReps := []*cardM.CardEventResp{}

	for key, element := range matchedCardEventResps {
		matchedCardRewardEventResps := []*cardM.CardRewardEventResp{}

		if _, ok := allMatchedCardRewardEventMapper[key]; ok {
			matchedCardRewardEventResps = append(matchedCardRewardEventResps, allMatchedCardRewardEventMapper[key]...)
		}

		if _, ok := someMatchedCardRewardEventMapper[key]; ok {
			matchedCardRewardEventResps = append(matchedCardRewardEventResps, someMatchedCardRewardEventMapper[key]...)
		}

		element.CardRewardEventResps = matchedCardRewardEventResps

		sortedCardEventReps = append(sortedCardEventReps, element)
	}

	return sortedCardEventReps

}
