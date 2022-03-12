package onlinegame

import (
	"context"

	onlinegameM "example.com/creditcard/models/onlinegame"
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

const INSERT_STAT = "INSERT INTO onlinegame " +
	"(\"id\", \"name\", \"image_path\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, onlinegame *onlinegameM.Onlinegame) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		onlinegame.ID,
		onlinegame.Name,
		onlinegame.ImagePath,
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

const UPDATE_BY_ID_STAT = "UPDATE onlinegame SET " +
	" \"name\" = $1, \"image_path\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, onlinegame *onlinegameM.Onlinegame) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		onlinegame.Name,
		onlinegame.ImagePath,
		onlinegame.ID,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"image_path\" " +
	" FROM onlinegame "

func (im *impl) GetAll(ctx context.Context) ([]*onlinegameM.Onlinegame, error) {
	onlinegames := []*onlinegameM.Onlinegame{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		onlinegame := &onlinegameM.Onlinegame{}
		selector := []interface{}{
			&onlinegame.ID,
			&onlinegame.Name,
			&onlinegame.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		onlinegames = append(onlinegames, onlinegame)
	}

	return onlinegames, nil
}
