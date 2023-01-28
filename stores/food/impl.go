package food

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

const INSERT_STAT = "INSERT INTO food " +
	"(\"id\", \"name\", \"channel_label\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, food *channel.Food) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		food.ID,
		food.Name,
		food.ChannelLabels,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE food SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2" +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, food *channel.Food) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		food.Name,
		food.ChannelLabels,
		food.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM food "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Food, error) {
	foods := []*channel.Food{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		food := &channel.Food{}
		selector := []interface{}{
			&food.ID,
			&food.Name,
			&food.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM food WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Food, error) {
	food := &channel.Food{}

	selector := []interface{}{
		&food.ID,
		&food.Name,
		&food.ChannelLabels,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return food, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM food WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Food, error) {
	foods := []*channel.Food{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		food := &channel.Food{}
		selector := []interface{}{
			&food.ID,
			&food.Name,
			&food.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}
