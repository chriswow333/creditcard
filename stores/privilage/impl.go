package privilage

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	privilageM "example.com/creditcard/models/privilage"
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

const INSERT_PRIVILAGE_STAT = "INSERT INTO privilage(\"id\", card_id, \"name\", \"desc\", start_date, end_date, update_date, score) VALUES($1,$2,$3,$4,$5,$6,$7,$8)"

func (im *impl) Create(ctx context.Context, privilage *privilageM.Privilage) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		privilage.ID,
		privilage.CardID,
		privilage.Name,
		privilage.Desc,
		privilage.StartDate,
		privilage.EndDate,
		privilage.UpdateDate,
		privilage.Score,
	}
	if _, err := tx.Exec(INSERT_PRIVILAGE_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_STAT = "SELECT \"id\", card_id, \"name\", \"desc\", start_date, end_date, update_date, score FROM privilage WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*privilageM.Privilage, error) {

	privilage := &privilageM.Privilage{}

	selector := []interface{}{
		&privilage.ID,
		&privilage.CardID,
		&privilage.Name,
		&privilage.Desc,
		&privilage.StartDate,
		&privilage.EndDate,
		&privilage.UpdateDate,
		&privilage.Score,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return privilage, nil
}

const SELECT_ALL_STAT = "SELECT \"id\", card_id, \"name\", \"desc\", start_date, end_date, update_date, score FROM privilage"

func (im *impl) GetAll(ctx context.Context) ([]*privilageM.Privilage, error) {

	privilages := []*privilageM.Privilage{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		privilage := &privilageM.Privilage{}

		selector := []interface{}{
			&privilage.ID,
			&privilage.CardID,
			&privilage.Name,
			&privilage.Desc,
			&privilage.StartDate,
			&privilage.EndDate,
			&privilage.UpdateDate,
			&privilage.Score,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}

		privilages = append(privilages, privilage)

	}

	return privilages, nil

}

const SELECT_BY_CARDID_STAT = "SELECT \"id\", card_id, \"name\", \"desc\", start_date, end_date, update_date, score FROM privilage WHERE card_id = $1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*privilageM.Privilage, error) {

	privilages := []*privilageM.Privilage{}

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

		privilage := &privilageM.Privilage{}

		selector := []interface{}{
			&privilage.ID,
			&privilage.CardID,
			&privilage.Name,
			&privilage.Desc,
			&privilage.StartDate,
			&privilage.EndDate,
			&privilage.UpdateDate,
			&privilage.Score,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}

		privilages = append(privilages, privilage)

	}

	return privilages, nil
}
