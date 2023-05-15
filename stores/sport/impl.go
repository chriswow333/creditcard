package sport

import (
	"context"
	"runtime/debug"
	"strings"

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

const INSERT_STAT = "INSERT INTO sport " +
	"(\"id\", \"name\", \"label_types\", \"image_path\") VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, sport *channel.Sport) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		sport.ID,
		sport.Name,
		sport.LabelTypes,
		sport.ImagePath,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE sport SET " +
	" \"name\" = $1 " +
	" \"label_types\" = $2 " +
	" \"image_path\" = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, sport *channel.Sport) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		sport.Name,
		sport.LabelTypes,
		sport.ImagePath,
		sport.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"label_types\", \"image_path\" " +
	" FROM sport limit $1 offset $2 "

func (im *impl) GetAll(ctx context.Context, offset, limit int) ([]*channel.Sport, error) {
	sports := []*channel.Sport{}
	rows, err := im.psql.Query(SELECT_ALL_STAT, limit, offset)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		sport := &channel.Sport{}
		selector := []interface{}{
			&sport.ID,
			&sport.Name,
			&sport.LabelTypes,
			&sport.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		sports = append(sports, sport)
	}

	return sports, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"label_types\", \"image_path\" " +
	" FROM sport WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Sport, error) {
	sport := &channel.Sport{}

	selector := []interface{}{
		&sport.ID,
		&sport.Name,
		&sport.LabelTypes,
		&sport.ImagePath,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return sport, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"label_types\", \"image_path\" " +
	" FROM sport WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Sport, error) {
	sports := []*channel.Sport{}

	name := strings.Join(names, "|")
	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		sport := &channel.Sport{}
		selector := []interface{}{
			&sport.ID,
			&sport.Name,
			&sport.LabelTypes,
			&sport.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		sports = append(sports, sport)
	}

	return sports, nil
}
