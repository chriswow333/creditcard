package conn

import (
	"errors"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	conn "example.com/creditcard/app/view_card/utils/conn"
)

type impl struct {
	dig.In

	psql *pgx.ConnPool
}

func New(psql *pgx.ConnPool) conn.Service {
	return &impl{
		psql: psql,
	}
}

func (im *impl) GetConn() (*conn.Connection, error) {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return nil, err
	}

	conn := &conn.Connection{
		Connection: tx,
	}
	return conn, nil
}

func (im *impl) Commit(conn *conn.Connection) error {

	tx, ok := conn.Connection.(*pgx.Tx)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"msg": "cast failed",
		})
		return errors.New("cast failed")
	}

	if err := tx.Commit(); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}
	return nil
}

func (im *impl) RollBack(conn *conn.Connection) error {

	tx, ok := conn.Connection.(*pgx.Tx)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"msg": "cast failed",
		})
		return errors.New("cast failed")
	}

	tx.Rollback()
	return nil
}

func (im *impl) Exec(conn *conn.Connection, sql string, updater ...interface{}) error {

	tx, ok := conn.Connection.(*pgx.Tx)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"msg": "cast failed",
		})
		return errors.New("cast failed")
	}

	if _, err := tx.Exec(sql, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)
		return err
	}
	return nil
}
