package customization

import (
	"context"

	customizationM "example.com/creditcard/models/customization"
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

const SELECT_BY_CARDID_STAT = "SELECT \"id\", \"name\", card_id, default_pass " +
	" FROM customization " +
	" WHERE \"card_id\"=$1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*customizationM.Customization, error) {
	customizations := []*customizationM.Customization{}

	conditions := []interface{}{
		cardID,
	}
	rows, err := im.psql.Query(SELECT_BY_CARDID_STAT, conditions...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		customization := &customizationM.Customization{}
		selector := []interface{}{
			&customization.ID,
			&customization.Name,
			&customization.CardID,
			&customization.DefaultPass,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		customizations = append(customizations, customization)
	}

	return customizations, nil

}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", card_id, default_pass " +
	" FROM customization " +
	" WHERE \"id\"=$1"

func (im *impl) GetByID(ctx context.Context, ID string) (*customizationM.Customization, error) {
	customization := &customizationM.Customization{}

	selector := []interface{}{
		&customization.ID,
		&customization.Name,
		&customization.CardID,
		&customization.DefaultPass,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return customization, nil

}
