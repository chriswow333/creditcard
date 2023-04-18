package publicutilities

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

const INSERT_STAT = "INSERT INTO publicutility " +
	"(\"id\", \"name\", \"channel_label\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, publicUtility *channel.PublicUtility) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		publicUtility.ID,
		publicUtility.Name,
		publicUtility.ChannelLabels,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE publicutility SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, publicUtility *channel.PublicUtility) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		publicUtility.Name,
		publicUtility.ChannelLabels,
		publicUtility.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM publicutility "

func (im *impl) GetAll(ctx context.Context) ([]*channel.PublicUtility, error) {

	publicUtilities := []*channel.PublicUtility{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		publicUtility := &channel.PublicUtility{}
		selector := []interface{}{
			&publicUtility.ID,
			&publicUtility.Name,
			&publicUtility.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		publicUtilities = append(publicUtilities, publicUtility)
	}

	return publicUtilities, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM publicutility WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.PublicUtility, error) {

	publicUtility := &channel.PublicUtility{}

	selector := []interface{}{
		&publicUtility.ID,
		&publicUtility.Name,
		&publicUtility.ChannelLabels,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return publicUtility, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM publicutility WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.PublicUtility, error) {
	publicutilities := []*channel.PublicUtility{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		publicutility := &channel.PublicUtility{}
		selector := []interface{}{
			&publicutility.ID,
			&publicutility.Name,
			&publicutility.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		publicutilities = append(publicutilities, publicutility)
	}

	return publicutilities, nil
}