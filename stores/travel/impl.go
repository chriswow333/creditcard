package travel

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

const INSERT_STAT = "INSERT INTO travel " +
	"(\"id\", \"name\", \"channel_label\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, travel *channel.Travel) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		travel.ID,
		travel.Name,
		travel.ChannelLabels,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE travel SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, travel *channel.Travel) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		travel.Name,
		travel.ID,
		travel.ChannelLabels,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM travel "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Travel, error) {

	travels := []*channel.Travel{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		travel := &channel.Travel{}
		selector := []interface{}{
			&travel.ID,
			&travel.Name,
			&travel.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		travels = append(travels, travel)
	}

	return travels, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM Travel WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Travel, error) {

	travel := &channel.Travel{}

	selector := []interface{}{
		&travel.ID,
		&travel.Name,
		&travel.ChannelLabels,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return travel, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM travel WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Travel, error) {
	travels := []*channel.Travel{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		travel := &channel.Travel{}
		selector := []interface{}{
			&travel.ID,
			&travel.Name,
			&travel.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		travels = append(travels, travel)
	}

	return travels, nil
}
