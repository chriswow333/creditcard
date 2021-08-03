package card

import (
	"context"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	cardM "example.com/creditcard/app/view_card/models/card"
	"example.com/creditcard/app/view_card/models/common"
	"example.com/creditcard/app/view_card/stores/card"
	"example.com/creditcard/app/view_card/stores/feature"
	"example.com/creditcard/app/view_card/utils/conn"
)

var (
	timeNow = time.Now
)

type impl struct {
	dig.In

	cardStore    card.Store    `name:"cardStore"`
	featureStore feature.Store `name:"featureStore"`
	connService  conn.Service  `name:"connService"`
}

func New(
	cardStore card.Store,
	featureStore feature.Store,
	connService conn.Service,
) Service {
	return &impl{
		cardStore:    cardStore,
		featureStore: featureStore,
		connService:  connService,
	}
}

func (im *impl) Create(ctx context.Context, cardRepr *cardM.Repr) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	conn, err := im.connService.GetConn()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}
	defer im.connService.RollBack(conn)

	validateTime := &common.ValidateTime{
		StartTime: cardRepr.StartTime,
		EndTime:   cardRepr.EndTime,
	}

	card := &cardM.Card{
		ID:                      id.String(),
		Name:                    cardRepr.Name,
		Icon:                    cardRepr.Icon,
		BankID:                  cardRepr.BankID,
		MaxPoint:                cardRepr.MaxPoint,
		FeatureDesc:             cardRepr.FeatureDesc,
		ValidateTime:            validateTime,
		ApplicantQualifications: cardRepr.ApplicantQualifications,
		UpdateDate:              timeNow().Unix(),
	}

	if err := im.cardStore.Create(ctx, conn, card); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	if cardRepr.Features == nil {
		return nil
	}

	feature := &cardM.Feature{
		FeatureTypes: cardRepr.Features,
	}

	if err := im.featureStore.CreateByCardID(ctx, conn, id.String(), feature); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	im.connService.Commit(conn)

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.Repr, error) {

	card, err := im.cardStore.GetByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	cardRepr := &cardM.Repr{
		ID:                      card.ID,
		Name:                    card.Name,
		Icon:                    card.Icon,
		BankID:                  card.BankID,
		MaxPoint:                card.MaxPoint,
		FeatureDesc:             card.FeatureDesc,
		StartTime:               card.ValidateTime.StartTime,
		EndTime:                 card.ValidateTime.EndTime,
		ApplicantQualifications: card.ApplicantQualifications,
	}

	feature, err := im.getFeaturesByCardID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}
	cardRepr.Features = feature.FeatureTypes

	return cardRepr, nil
}

func (im *impl) UpdateByID(ctx context.Context, cardRepr *cardM.Repr) error {

	conn, err := im.connService.GetConn()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}
	defer im.connService.RollBack(conn)

	validateTime := &common.ValidateTime{
		StartTime: cardRepr.StartTime,
		EndTime:   cardRepr.EndTime,
	}

	card := &cardM.Card{
		ID:                      cardRepr.ID,
		Name:                    cardRepr.Name,
		Icon:                    cardRepr.Icon,
		BankID:                  cardRepr.BankID,
		MaxPoint:                cardRepr.MaxPoint,
		FeatureDesc:             cardRepr.FeatureDesc,
		ValidateTime:            validateTime,
		ApplicantQualifications: cardRepr.ApplicantQualifications,

		UpdateDate: timeNow().Unix(),
	}

	if err := im.cardStore.UpdateByID(ctx, conn, card); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	feature := &cardM.Feature{
		FeatureTypes: cardRepr.Features,
	}

	if err := im.featureStore.DeleteByCardID(ctx, conn, card.ID); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	if err := im.featureStore.CreateByCardID(ctx, conn, card.ID, feature); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	im.connService.Commit(conn)

	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*cardM.Repr, error) {

	cards, err := im.cardStore.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	cardReprs := []*cardM.Repr{}

	for _, card := range cards {
		cardRepr := &cardM.Repr{
			ID:                      card.ID,
			Name:                    card.Name,
			Icon:                    card.Icon,
			BankID:                  card.BankID,
			MaxPoint:                card.MaxPoint,
			FeatureDesc:             card.FeatureDesc,
			StartTime:               card.ValidateTime.StartTime,
			EndTime:                 card.ValidateTime.EndTime,
			ApplicantQualifications: card.ApplicantQualifications,
		}
		cardReprs = append(cardReprs, cardRepr)
	}

	for _, c := range cardReprs {
		feature, err := im.getFeaturesByCardID(ctx, c.ID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": "",
			}).Fatal(err)
			return nil, err
		}
		c.Features = feature.FeatureTypes
	}

	return cardReprs, nil
}

func (im *impl) GetByBankID(ctx context.Context, bankID string) ([]*cardM.Repr, error) {

	cards, err := im.cardStore.GetByBankID(ctx, bankID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	cardReprs := []*cardM.Repr{}
	for _, card := range cards {
		cardRepr := &cardM.Repr{
			ID:                      card.ID,
			Name:                    card.Name,
			Icon:                    card.Icon,
			BankID:                  card.BankID,
			MaxPoint:                card.MaxPoint,
			FeatureDesc:             card.FeatureDesc,
			StartTime:               card.ValidateTime.StartTime,
			EndTime:                 card.ValidateTime.EndTime,
			ApplicantQualifications: card.ApplicantQualifications,
		}
		cardReprs = append(cardReprs, cardRepr)
	}

	for _, c := range cardReprs {
		feature, err := im.getFeaturesByCardID(ctx, c.ID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": "",
			}).Fatal(err)
			return nil, err
		}
		c.Features = feature.FeatureTypes
	}

	return cardReprs, nil
}

func (im *impl) getFeaturesByCardID(ctx context.Context, cardID string) (*cardM.Feature, error) {
	feature, err := im.featureStore.GetByCardID(ctx, cardID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return nil, err
	}

	return feature, nil
}
