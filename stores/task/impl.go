package task

import (
	"context"

	taskM "example.com/creditcard/models/task"
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

const SELECT_BY_CARDID_STAT = "SELECT \"id\", \"name\", \"descs\", " +
	" card_id, task_type, task_type_model, default_pass " +
	" FROM task " +
	" WHERE \"card_id\"=$1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*taskM.Task, error) {
	tasks := []*taskM.Task{}

	conditions := []interface{}{
		cardID,
	}
	rows, err := im.psql.Query(SELECT_BY_CARDID_STAT, conditions...)
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
			&task.Descs,
			&task.TaskType,
			&task.TaskTypeModel,
			&task.CardID,
			&task.DefaultPass,
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

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"descs\", " +
	" card_id, task_type, task_type_model, default_pass " +
	" FROM task " +
	" WHERE \"id\"=$1"

func (im *impl) GetByID(ctx context.Context, ID string) (*taskM.Task, error) {
	task := &taskM.Task{}

	selector := []interface{}{
		&task.ID,
		&task.Name,
		&task.Descs,
		&task.CardID,
		&task.TaskType,
		&task.TaskTypeModel,
		&task.DefaultPass,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": err,
		}).Error(err)
		return nil, err
	}

	return task, nil

}

const INSERT_TASK_STAT = "INSERT INTO task " +
	"(\"id\", \"name\", \"descs\", card_id, task_type, task_type_model, default_pass) " +
	" VALUES($1, $2, $3, $4, $5, $6, $7)"

func (im *impl) Create(ctx context.Context, task *taskM.Task) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": err,
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		task.ID,
		task.Name,
		task.Descs,
		task.CardID,
		task.TaskType,
		task.TaskTypeModel,
		task.DefaultPass,
	}
	if _, err := tx.Exec(INSERT_TASK_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	tx.Commit()
	return nil
}
