package constraint

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
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

const INSERT_CONSTRAINT_STAT = "INSERT INTO \"constraint\"(\"id\", privilage_id, \"name\", \"desc\", " +
	"start_date, end_date, update_date, limit_mx, limit_mn, constraint_body) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

func (im *impl) Create(ctx context.Context, constraint *constraintM.Constraint) error {

	fmt.Println(constraint)
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()
	limMx := 0
	limMn := 0
	if constraint.Limit != nil {
		limMx = constraint.Limit.Max
		limMn = constraint.Limit.Min
	}

	updater := []interface{}{
		constraint.ID,
		constraint.PrivilageID,
		constraint.Name,
		constraint.Desc,
		constraint.StartDate,
		constraint.EndDate,
		constraint.UpdateDate,
		limMx,
		limMn,
		constraint.ConstraintBody,
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

const SELECT_STAT = "SELECT \"id\", privilage_id, \"name\", \"desc\", " +
	"start_date, end_date, update_date, limit_mx, limit_mn, constraint_body FROM \"constraint\" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*constraintM.Constraint, error) {

	constraint := &constraintM.Constraint{}

	limMx := 0
	limMn := 0
	if constraint.Limit != nil {
		limMx = constraint.Limit.Max
		limMn = constraint.Limit.Min
	}

	updater := []interface{}{
		&constraint.ID,
		&constraint.PrivilageID,
		&constraint.Name,
		&constraint.Desc,
		&constraint.StartDate,
		&constraint.EndDate,
		&constraint.UpdateDate,
		limMx,
		limMn,
		&constraint.ConstraintBody,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return constraint, nil
}

const SELECT_ALL_STAT = "SELECT \"id\", privilage_id, \"name\", \"desc\", " +
	"start_date, end_date, update_date, limit_mx, limit_mn, constraint_body FROM \"constraint\""

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
		limit := &constraintM.Limit{}
		selector := []interface{}{
			&constraint.ID,
			&constraint.PrivilageID,
			&constraint.Name,
			&constraint.Desc,
			&constraint.StartDate,
			&constraint.EndDate,
			&constraint.UpdateDate,
			&limit.Max,
			&limit.Min,
			&constraint.ConstraintBody,
		}
		constraint.Limit = limit

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

const SELECT_BY_PRIVILAGEID_STAT = "SELECT \"id\", privilage_id, \"name\", \"desc\", " +
	"start_date, end_date, update_date, limit_mx, limit_mn, constraint_body FROM \"constraint\" WHERE privilage_id = $1"

func (im *impl) GetByPrivilageID(ctx context.Context, privilageID string) ([]*constraintM.Constraint, error) {

	constraints := []*constraintM.Constraint{}

	conditions := []interface{}{
		privilageID,
	}

	rows, err := im.psql.Query(SELECT_BY_PRIVILAGEID_STAT, conditions...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)

		return nil, err
	}

	for rows.Next() {
		constraint := &constraintM.Constraint{}
		limit := &constraintM.Limit{}
		selector := []interface{}{
			&constraint.ID,
			&constraint.PrivilageID,
			&constraint.Name,
			&constraint.Desc,
			&constraint.StartDate,
			&constraint.EndDate,
			&constraint.UpdateDate,
			&limit.Max,
			&limit.Min,
			&constraint.ConstraintBody,
		}
		constraint.Limit = limit

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
