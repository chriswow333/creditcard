package payload

import (
	"context"
	"runtime/debug"

	payloadM "example.com/creditcard/models/payload"
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

const UPDATE_BY_ID_STAT = "UPDATE reward SET " +
	" payload = $1 WHERE \"id\" = $2"

func (im *impl) UpdateByID(ctx context.Context, rewardID string, payloads []*payloadM.Payload) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		payloads,
		rewardID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}
