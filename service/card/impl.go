package card

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	cardM "example.com/creditcard/models/card"
	constraintM "example.com/creditcard/models/constraint"
	payloadM "example.com/creditcard/models/payload"
	"example.com/creditcard/service/bank"
	"example.com/creditcard/service/constraint"
	"example.com/creditcard/stores/card"
	"example.com/creditcard/stores/card_reward"
	"example.com/creditcard/stores/reward"
	"github.com/sirupsen/logrus"

	uuid "github.com/nu7hatch/gouuid"

	"go.uber.org/dig"
)

var (
	timeNow = time.Now
)

type impl struct {
	dig.In

	cardStore         card.Store
	rewardStore       reward.Store
	cardRewardStore   card_reward.Store
	bankService       bank.Service
	constraintService constraint.Service
}

func New(
	cardStore card.Store,
	rewardStore reward.Store,
	cardRewardStore card_reward.Store,
	bankService bank.Service,
	constraintService constraint.Service,
) Service {
	return &impl{
		cardStore:         cardStore,
		rewardStore:       rewardStore,
		cardRewardStore:   cardRewardStore,
		bankService:       bankService,
		constraintService: constraintService,
	}
}

func (im *impl) Create(ctx context.Context, card *cardM.Card) error {

	card.UpdateDate = timeNow().Unix()

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)

		return err
	}
	card.ID = id.String()

	if err := im.cardStore.Create(ctx, card); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.Card, error) {
	card, err := im.cardStore.GetByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	cardRewards, err := im.cardRewardStore.GetByCardID(ctx, card.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for _, cr := range cardRewards {
		rewards, err := im.rewardStore.GetByCardRewardID(ctx, cr.ID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		cr.Rewards = rewards
	}

	card.CardRewards = cardRewards

	return card, nil
}

func (im *impl) GetRespByID(ctx context.Context, ID string) (*cardM.CardResp, error) {

	card, err := im.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	cardResp := cardM.TransferCardResp(card)

	bankResp, err := im.bankService.GetRespByID(ctx, card.BankID)
	if err != nil {
		return nil, err
	}
	cardResp.BankName = bankResp.Name

	// transfer constraintResp

	payloadRespMap := make(map[string]*payloadM.PayloadResp)
	for _, c := range cardResp.CardRewardResps {
		for _, r := range c.RewardResps {
			for _, p := range r.PayloadResps {
				payloadRespMap[p.ID] = p
			}
		}
	}

	for _, c := range card.CardRewards {
		for _, r := range c.Rewards {
			for _, p := range r.Payloads {

				constraintResp, err := im.getConstraintResp(ctx, p.Constraint)
				if err != nil {
					logrus.Error("im.setConstraintRespToPayloadResp Error")
					return nil, err
				}

				if payloadResp, ok := payloadRespMap[p.ID]; ok {
					payloadResp.ConstraintResp = constraintResp
				} else {
					logrus.Error("im.setConstraintRespToPayloadResp not found payloadID Error")
					return nil, err
				}

			}
		}
	}

	return cardResp, nil
}
func (im *impl) getConstraintResp(ctx context.Context, constraint *constraintM.Constraint) (*constraintM.ConstraintResp, error) {

	constraintResp, err := constraintM.TransferConstraintResp(ctx, constraint, im.constraintService)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return constraintResp, nil
}

func (im *impl) UpdateByID(ctx context.Context, card *cardM.Card) error {
	if err := im.cardStore.UpdateByID(ctx, card); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*cardM.Card, error) {
	cards, err := im.cardStore.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return cards, nil
}

func (im *impl) GetRespAll(ctx context.Context) ([]*cardM.CardResp, error) {

	cards, err := im.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	cardResps := []*cardM.CardResp{}

	for _, c := range cards {
		cardResp := cardM.TransferCardResp(c)
		cardResps = append(cardResps, cardResp)
	}

	return cardResps, nil
}

func (im *impl) GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error) {
	cards, err := im.cardStore.GetByBankID(ctx, bankID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return cards, nil
}

func (im *impl) GetRespByBankID(ctx context.Context, bankID string) ([]*cardM.CardResp, error) {

	cards, err := im.GetByBankID(ctx, bankID)
	if err != nil {
		return nil, err
	}

	cardResps := []*cardM.CardResp{}

	for _, c := range cards {
		cardResp := cardM.TransferCardResp(c)
		cardResps = append(cardResps, cardResp)
	}

	return cardResps, nil
}

func (im *impl) CreateCardReward(ctx context.Context, cardReward *cardM.CardReward) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)

		return err
	}
	cardReward.ID = id.String()

	if err := im.cardRewardStore.Create(ctx, cardReward); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": err,
		}).Fatal(err)
		return err
	}

	for _, r := range cardReward.Rewards {
		r.CardRewardID = cardReward.ID

		id, err := uuid.NewV4()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": "",
			}).Fatal(err)
			return err
		}

		r.ID = id.String()

		for _, p := range r.Payloads {

			pid, err := uuid.NewV4()

			if err != nil {
				logrus.WithFields(logrus.Fields{
					"msg": "",
				}).Fatal(err)
				return err
			}
			p.ID = pid.String()
		}

		im.rewardStore.Create(ctx, r)
	}

	return nil
}

func (im *impl) EvaluateConstraintLogic(ctx context.Context, cardRewardID string, constraintIDs []string) (bool, string, error) {

	cardReward, err := im.cardRewardStore.GetByID(ctx, cardRewardID)
	if err != nil {
		return false, "internal error", err
	}

	constraintSet := make(map[string]bool)
	for _, constraintID := range constraintIDs {
		constraintSet[constraintID] = true
	}

	fmt.Println(constraintSet)

	for _, logic := range cardReward.ConstraintPassLogics {

		fmt.Println(logic.Logic)

		pass, _, err := checkConstraintLogic(logic.Logic, constraintSet)
		if err != nil {
			return false, "internal error", err
		}

		fmt.Println(pass)

		if !pass {
			return false, logic.Message, nil
		}
	}

	return true, "", nil
}

/**

A, B, C are constraint ID
((A^B)C)

if event has no constraint ID, that means true

*/
func checkConstraintLogic(constraintPassLogic string, constraintIDs map[string]bool) (bool, bool, error) {

	pos := 0

	for pos = 0; pos < len(constraintPassLogic); pos++ {

		ch := constraintPassLogic[pos : pos+1]

		if ch == "(" {
			lastPos := strings.LastIndex(constraintPassLogic, ")")
			if lastPos == -1 {
				return false, false, errors.New("constraintPassLogic is illegal")
			}

			pass, exist, err := checkConstraintLogic(constraintPassLogic[1:lastPos], constraintIDs)
			if err != nil {
				return false, exist, err
			} else {
				return pass, exist, nil
			}

		} else if ch == "&" || ch == "|" || ch == "^" {
			constraintPassLogicPrev := constraintPassLogic[0:pos]
			constraintPassLogicLast := constraintPassLogic[pos+1:]
			passPrev, existPrev, err := checkConstraintLogic(constraintPassLogicPrev, constraintIDs)
			if err != nil {
				return false, false, err
			}
			passLast, existLast, err := checkConstraintLogic(constraintPassLogicLast, constraintIDs)
			if err != nil {
				return false, false, err
			}

			switch ch {
			case "&":
				return (passPrev && passLast) || (!existPrev && !existLast), existPrev, nil // if no one exist, return true
			case "|":
				return (passPrev || passLast) || (!existPrev && !existLast), existLast, nil // if no one exist, return true
			case "^":
				return (passPrev || passLast) && !(passPrev && passLast) || (!existPrev && !existLast), existLast, nil // if no one exist, return true
			}
		}
	}

	if _, ok := constraintIDs[constraintPassLogic]; ok {
		return true, true, nil
	} else {
		return false, false, nil
	}
}
