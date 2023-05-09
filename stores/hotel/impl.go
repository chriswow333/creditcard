package hotel

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

const INSERT_STAT = "INSERT INTO hotel " +
	"(\"id\", \"name\", \"channel_label\", \"image_path\") VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, hotel *channel.Hotel) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		hotel.ID,
		hotel.Name,
		hotel.ChannelLabels,
		hotel.ImagePath,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE hotel SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2" +
	" \"image_path\" = $3" +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, hotel *channel.Hotel) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		hotel.Name,
		hotel.ChannelLabels,
		hotel.ImagePath,
		hotel.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\"  " +
	" FROM hotel "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Hotel, error) {

	hotels := []*channel.Hotel{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		hotel := &channel.Hotel{}
		selector := []interface{}{
			&hotel.ID,
			&hotel.Name,
			&hotel.ChannelLabels,
			&hotel.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM hotel WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Hotel, error) {
	hotel := &channel.Hotel{}

	selector := []interface{}{
		&hotel.ID,
		&hotel.Name,
		&hotel.ChannelLabels,
		&hotel.ImagePath,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return hotel, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM hotel WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Hotel, error) {
	hotels := []*channel.Hotel{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		hotel := &channel.Hotel{}
		selector := []interface{}{
			&hotel.ID,
			&hotel.Name,
			&hotel.ChannelLabels,
			&hotel.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
