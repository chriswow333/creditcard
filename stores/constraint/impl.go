package constraint

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	constraintM "example.com/creditcard/models/constraint"
)

type impl struct {
	dig.In

	psql *pgx.ConnPool
}

func New(
	psql *pgx.ConnPool,
) Store {

	return &impl{
		psql: psql,
	}
}

const INSERT_CONSTRAINT_STAT = "INSERT INTO \"constraint\"(\"id\", privilage_id, \"name\", \"desc\", \"operator\", start_date, end_date, update_date) VALUES($1,$2,$3,$4,$5,$6,$7,$8)"

func (im *impl) Create(ctx context.Context, constraint *constraintM.Constraint) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}

	updater := []interface{}{
		id.String(),
		constraint.PrivilageID,
		constraint.Name,
		constraint.Desc,
		constraint.Operator,
		constraint.StartDate,
		constraint.EndDate,
		constraint.UpdateDate,
	}

	if _, err := tx.Exec(INSERT_CONSTRAINT_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_STAT = "SELECT \"id\", privilage_id, \"name\", \"desc\", \"operator\", start_date, end_date, update_date FROM \"constraint\" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*constraintM.Constraint, error) {

	constraint := &constraintM.Constraint{}
	constraint.Limit = &constraintM.Limit{}
	updater := []interface{}{
		&constraint.ID,
		&constraint.PrivilageID,
		&constraint.Name,
		&constraint.Desc,
		&constraint.Operator,
		&constraint.StartDate,
		&constraint.EndDate,
		&constraint.UpdateDate,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return constraint, nil
}

const SELECT_ALL_STAT = "SELECT \"id\", privilage_id, \"name\", \"desc\", \"operator\", start_date, end_date, update_date FROM \"constraint\""

func (im *impl) GetAll(ctx context.Context) ([]*constraintM.Constraint, error) {

	constraints := []*constraintM.Constraint{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}
	for rows.Next() {
		constraint := &constraintM.Constraint{}
		selector := []interface{}{
			&constraint.ID,
			&constraint.PrivilageID,
			&constraint.Name,
			&constraint.Desc,
			&constraint.Operator,
			&constraint.StartDate,
			&constraint.EndDate,
			&constraint.UpdateDate,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}
		constraints = append(constraints, constraint)
	}
	return constraints, nil
}

const SELECT_BY_PRIVILAGEID_STAT = "SELECT \"id\", privilage_id, \"name\", \"desc\", \"operator\", start_date, end_date, update_date FROM \"constraint\" WHERE privilage_id = $1"

func (im *impl) GetByPrivilageID(ctx context.Context, privilageID string) ([]*constraintM.Constraint, error) {
	constraints := []*constraintM.Constraint{}

	condition := []interface{}{
		privilageID,
	}
	fmt.Println(privilageID)
	rows, err := im.psql.Query(SELECT_BY_PRIVILAGEID_STAT, condition...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)

		return nil, err
	}

	for rows.Next() {
		constraint := &constraintM.Constraint{}
		selector := []interface{}{
			&constraint.ID,
			&constraint.PrivilageID,
			&constraint.Name,
			&constraint.Desc,
			&constraint.Operator,
			&constraint.StartDate,
			&constraint.EndDate,
			&constraint.UpdateDate,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}
		constraints = append(constraints, constraint)
	}
	return constraints, nil

}
