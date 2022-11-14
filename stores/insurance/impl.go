package insurance

import (
	"context"

	"example.com/creditcard/models/channel"
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

const INSERT_STAT = "INSERT INTO insurance " +
	"(\"id\", \"name\", \"channel_label\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, insurance *channel.Insurance) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		insurance.ID,
		insurance.Name,
		insurance.ChannelLabels,
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

const UPDATE_BY_ID_STAT = "UPDATE insurance SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, insurance *channel.Insurance) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		insurance.Name,
		insurance.ChannelLabels,
		insurance.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM insurance "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Insurance, error) {

	insurances := []*channel.Insurance{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		insurance := &channel.Insurance{}
		selector := []interface{}{
			&insurance.ID,
			&insurance.Name,
			&insurance.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		insurances = append(insurances, insurance)
	}

	return insurances, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM insurance WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Insurance, error) {
	insurance := &channel.Insurance{}

	selector := []interface{}{
		&insurance.ID,
		&insurance.Name,
		&insurance.ChannelLabels,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return insurance, nil
}
