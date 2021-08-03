package feature

import (
	"context"

	cardM "example.com/creditcard/app/view_card/models/card"
	"example.com/creditcard/app/view_card/models/common"
	"example.com/creditcard/app/view_card/utils/conn"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type impl struct {
	dig.In

	psql        *pgx.ConnPool `name:"psql"`
	connService conn.Service  `name:"connService"`
}

func New(
	psql *pgx.ConnPool,
	connService conn.Service,
) Store {
	return &impl{
		psql:        psql,
		connService: connService,
	}
}

const INSERT_STAT = "INSERT INTO feature " +
	" (card_id, type) " +
	" VALUES ($1, $2)"

func (im *impl) CreateByCardID(ctx context.Context, conn *conn.Connection, cardID string, feature *cardM.Feature) error {

	for _, f_type := range feature.FeatureTypes {
		updater := []interface{}{
			cardID,
			int(f_type),
		}
		if err := im.connService.Exec(conn, INSERT_STAT, updater...); err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": err,
			}).Fatal(err)
			return err
		}
	}

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
		featureType := new(int32)

		selector := []interface{}{
			featureType,
		}

		if err := rows.Scan(selector...); err != nil {

			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}

		f, err := common.ConvertFeature(*featureType)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}
		feature.FeatureTypes = append(feature.FeatureTypes, f)
	}

	return feature, nil
}

const DELETE_STAT = "DELETE FROM feature " +
	" WHERE card_id = $1 "

func (im *impl) DeleteByCardID(ctx context.Context, conn *conn.Connection, cardID string) error {

	updater := []interface{}{
		cardID,
	}
	if err := im.connService.Exec(conn, INSERT_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)
		return err
	}

	return nil
}
