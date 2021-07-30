package task

import (
	"context"

	taskM "example.com/creditcard/app/view_card/models/task"
	"example.com/creditcard/app/view_card/utils/conn"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type impl struct {
	dig.In

	psql        *pgx.ConnPool
	connService conn.Service
}

func New(
	psql *pgx.ConnPool,
	connService conn.Service,
) Store {
	return &impl{
		psql:        psql,
		connService: connService,
	}
}

const INSERT_STAT = "INSERT INTO task " +
	" (\"id\", \"name\", \"desc\", reward_id, point, update_date) " +
	" VALUES ($1, $2, $3, $4, $5, $6)"

func (im *impl) Create(ctx context.Context, conn *conn.Connection, task *taskM.Task) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		task.ID,
		task.Name,
		task.Desc,
		task.RewardID,
		task.Point,
		task.UpdateDate,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	tx.Commit()

	return nil
}

func (im *impl) CreateTasks(ctx context.Context, conn *conn.Connection, tasks []*taskM.Task) error {

	for _, task := range tasks {
		updater := []interface{}{
			task.ID,
			task.Name,
			task.Desc,
			task.RewardID,
			task.Point,
			task.UpdateDate,
		}

		if err := im.connService.Exec(conn, INSERT_STAT, updater...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Fatal(err)

			return err
		}
	}

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE task SET " +
	" \"name\" = $1, \"desc\" = $2, " +
	" reward_id = $3, point = $4, update_date = $5 " +
	" where \"id\" = $6"

func (im *impl) UpdateByRewardID(ctx context.Context, conn *conn.Connection, tasks []*taskM.Task) error {

	for _, task := range tasks {

		updater := []interface{}{
			task.Name,
			task.Desc,
			task.RewardID,
			task.Point,
			task.UpdateDate,
			task.ID,
		}

		if err := im.connService.Exec(conn, UPDATE_BY_ID_STAT, updater...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			})
			return err
		}

	}

	return nil
}

const SELECT_REWARDID_STAT = "SELECT \"id\", \"name\", \"desc\", " +
	" reward_id, point, update_date " +
	" FROM task " +
	" WHERE reward_id = $1"

func (im *impl) GetByRewardID(ctx context.Context, rewardID string) ([]*taskM.Task, error) {

	tasks := []*taskM.Task{}

	conditions := []interface{}{
		rewardID,
	}
	rows, err := im.psql.Query(SELECT_REWARDID_STAT, conditions...)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		task := &taskM.Task{}

		selector := []interface{}{
			&task.ID,
			&task.Name,
			&task.Desc,
			&task.RewardID,
			&task.Point,
			&task.UpdateDate,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}

		tasks = append(tasks, task)

	}

	return tasks, nil
}

const DELETE_STAT = "DELETE FROM task " +
	" WHERE reward_id = $1 "

func (im *impl) DeleteByRewardID(ctx context.Context, conn *conn.Connection, rewardID string) error {

	updater := []interface{}{
		rewardID,
	}
	if err := im.connService.Exec(conn, DELETE_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)
		return err
	}

	return nil
}
