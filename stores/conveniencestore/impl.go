package conveniencestore

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

const INSERT_STAT = "INSERT INTO conveniencestore " +
	"(\"id\", \"name\", \"label_types\", \"image_path\") VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, convenienceStore *channel.ConvenienceStore) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		convenienceStore.ID,
		convenienceStore.Name,
		convenienceStore.LabelTypes,
		convenienceStore.ImagePath,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE conveniencestore SET " +
	" \"name\" = $1 " +
	" \"label_types\" = $2 " +
	" \"image_path\" = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, convenienceStore *channel.ConvenienceStore) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		convenienceStore.Name,
		convenienceStore.LabelTypes,
		convenienceStore.ImagePath,
		convenienceStore.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"label_types\", \"image_path\" " +
	" FROM conveniencestore limit $1 offset $2 "

func (im *impl) GetAll(ctx context.Context, offset, limit int) ([]*channel.ConvenienceStore, error) {

	convenienceStores := []*channel.ConvenienceStore{}

	rows, err := im.psql.Query(SELECT_ALL_STAT, limit, offset)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		convenienceStore := &channel.ConvenienceStore{}
		selector := []interface{}{
			&convenienceStore.ID,
			&convenienceStore.Name,
			&convenienceStore.LabelTypes,
			&convenienceStore.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		convenienceStores = append(convenienceStores, convenienceStore)
	}

	return convenienceStores, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"label_types\", \"image_path\" " +
	" FROM conveniencestore WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.ConvenienceStore, error) {
	convenienceStore := &channel.ConvenienceStore{}

	selector := []interface{}{
		&convenienceStore.ID,
		&convenienceStore.Name,
		&convenienceStore.LabelTypes,
		&convenienceStore.ImagePath,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return convenienceStore, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"label_types\", \"image_path\" " +
	" FROM conveniencestore WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.ConvenienceStore, error) {
	conveniencestores := []*channel.ConvenienceStore{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		conveniencestore := &channel.ConvenienceStore{}
		selector := []interface{}{
			&conveniencestore.ID,
			&conveniencestore.Name,
			&conveniencestore.LabelTypes,
			&conveniencestore.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		conveniencestores = append(conveniencestores, conveniencestore)
	}

	return conveniencestores, nil
}
