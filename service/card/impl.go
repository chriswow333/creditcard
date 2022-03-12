package card

import (
	"context"
	"fmt"
	"time"

	cardM "example.com/creditcard/models/card"
	"example.com/creditcard/service/bank"
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

	cardStore       card.Store
	rewardStore     reward.Store
	cardRewardStore card_reward.Store
	bankService     bank.Service
}

func New(
	cardStore card.Store,
	rewardStore reward.Store,
	cardRewardStore card_reward.Store,
	bankService bank.Service,
) Service {
	return &impl{
		cardStore:       cardStore,
		rewardStore:     rewardStore,
		cardRewardStore: cardRewardStore,
		bankService:     bankService,
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
		fmt.Println("reward operator ", cr.RewardOperator)
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

	return cardResp, nil
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
			"msg": "",
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

		im.rewardStore.Create(ctx, r)
	}

	return nil
}
