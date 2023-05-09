package cinema

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

const INSERT_STAT = "INSERT INTO cinema " +
	"(\"id\", \"name\", \"channel_label\", \"image_path\") VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, cinema *channel.Cinema) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		cinema.ID,
		cinema.Name,
		cinema.ChannelLabels,
		cinema.ImagePath,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE cinema SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" \"image_path\" = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, cinema *channel.Cinema) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		cinema.Name,
		cinema.ChannelLabels,
		cinema.ImagePath,
		cinema.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM cinema "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Cinema, error) {

	cinemas := []*channel.Cinema{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		cinema := &channel.Cinema{}
		selector := []interface{}{
			&cinema.ID,
			&cinema.Name,
			&cinema.ChannelLabels,
			&cinema.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		cinemas = append(cinemas, cinema)
	}

	return cinemas, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM cinema WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Cinema, error) {

	cinema := &channel.Cinema{}

	selector := []interface{}{
		&cinema.ID,
		&cinema.Name,
		&cinema.ChannelLabels,
		&cinema.ImagePath,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return cinema, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM cinema WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Cinema, error) {
	cinemas := []*channel.Cinema{}

	name := strings.Join(names, "|")
	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)

	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		cinema := &channel.Cinema{}
		selector := []interface{}{
			&cinema.ID,
			&cinema.Name,
			&cinema.ChannelLabels,
			&cinema.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		cinemas = append(cinemas, cinema)
	}

	return cinemas, nil
}
