package conn

import (
	"errors"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"

	conn_pool "example.com/creditcard/app/view_card/utils/conn_pool"
)

type impl struct {
	psql *pgx.ConnPool
}

func New(psql *pgx.ConnPool) conn_pool.Service {
	return &impl{
		psql: psql,
	}
}

func (im *impl) GetConn() (*conn_pool.Connection, error) {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return nil, err
	}

	conn := &conn_pool.Connection{
		Connection: tx,
	}
	return conn, nil
}

func (im *impl) Commit(conn *conn_pool.Connection) error {

	tx, err := conn.Connection.(*pgx.Tx)
	if err {
		logrus.WithFields(logrus.Fields{
			"msg": "cast failed",
		}).Error(err)
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

func (im *impl) RollBack(conn *conn_pool.Connection) error {

	tx, err := conn.Connection.(*pgx.Tx)
	if err {
		logrus.WithFields(logrus.Fields{
			"msg": "cast failed",
		}).Error(err)
		return errors.New("cast failed")
	}

	if err := tx.Rollback(); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}
	return nil
}

func (im *impl) Exec(conn *conn_pool.Connection, sql string, updater ...interface{}) error {

	tx, err := conn.Connection.(*pgx.Tx)
	if err {
		logrus.WithFields(logrus.Fields{
			"msg": "cast failed",
		}).Error(err)
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
