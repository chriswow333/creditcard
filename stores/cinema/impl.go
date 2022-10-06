package cinema

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

const INSERT_STAT = "INSERT INTO cinema " +
	"(\"id\", \"name\") VALUES ($1, $2)"

func (im *impl) Create(ctx context.Context, cinema *channel.Cinema) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		cinema.ID,
		cinema.Name,
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

const UPDATE_BY_ID_STAT = "UPDATE cinema SET " +
	" \"name\" = $1 " +
	" where \"id\" = $2"

func (im *impl) UpdateByID(ctx context.Context, cinema *channel.Cinema) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		cinema.Name,
		cinema.ID,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\" " +
	" FROM cinema "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Cinema, error) {
	cinemas := []*channel.Cinema{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		cinema := &channel.Cinema{}
		selector := []interface{}{
			&cinema.ID,
			&cinema.Name,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		cinemas = append(cinemas, cinema)
	}

	return cinemas, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\" " +
	" FROM cinema WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Cinema, error) {
	cinema := &channel.Cinema{}

	selector := []interface{}{
		&cinema.ID,
		&cinema.Name,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return cinema, nil
}
