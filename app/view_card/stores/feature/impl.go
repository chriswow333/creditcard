package feature

import (
	"context"
	"errors"
	"strconv"

	cardM "example.com/creditcard/app/view_card/models/card"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type impl struct {
	dig.In

	psql *pgx.ConnPool
}

func New(psql *pgx.ConnPool) Store {
	return &impl{
		psql: psql,
	}
}

const INSERT_STAT = "INSERT INTO feature " +
	" (card_id, type, \"desc\") " +
	" VALUES ($1, $2, $3)"

func (im *impl) CreateByCardID(ctx context.Context, cardID string, feature *cardM.Feature) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	for _, f_type := range feature.FeatureTypes {

		updater := []interface{}{
			cardID,
			f_type,
		}

		if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Fatal(err)
			return err
		}

	}

	tx.Commit()

	return nil
}

const SELECT_CARDID_STAT = "SELECT \"type\" " +
	" FROM feature " +
	" WHERE card_id = $1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) (*cardM.Feature, error) {

	feature := &cardM.Feature{}

	conditions := []interface{}{
		cardID,
	}
	rows, err := im.psql.Query(SELECT_CARDID_STAT, conditions...)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		var featureType int
		selector := []interface{}{
			featureType,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}

		switch featureType {
		case int(cardM.ECommerce):
			feature.FeatureTypes = append(feature.FeatureTypes, cardM.ECommerce)
		case int(cardM.Supremarket):
			feature.FeatureTypes = append(feature.FeatureTypes, cardM.ECommerce)
		case int(cardM.Delivery):
			feature.FeatureTypes = append(feature.FeatureTypes, cardM.ECommerce)
		case int(cardM.Fee):
			feature.FeatureTypes = append(feature.FeatureTypes, cardM.ECommerce)
		case int(cardM.Transport):
			feature.FeatureTypes = append(feature.FeatureTypes, cardM.ECommerce)
		default:
			return nil, errors.New("No match featureType: " + strconv.Itoa(featureType))
		}
	}

	return feature, nil
}

const DELETE_STAT = "DELETE FROM feature " +
	" WHERE card_id = $1 "

func (im *impl) DeleteByCardID(ctx context.Context, cardID string) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		cardID,
	}
	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)
		return err
	}

	tx.Commit()

	return nil
}
