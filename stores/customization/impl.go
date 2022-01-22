package customization

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	customizationM "example.com/creditcard/models/customization"
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

const INSERT_CUSTOMIZARION_STAT = "INSERT INTO customization " +
	"(\"id\", \"name\", desc, link_url) VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, customization *customizationM.Customization) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		customization.ID,
		customization.Name,
		customization.Desc,
		customization.LinkURL,
	}

	if _, err := tx.Exec(INSERT_CUSTOMIZARION_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	tx.Commit()

	return nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", desc, link_url " +
	"FROM customization " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*customizationM.Customization, error) {

	customization := &customizationM.Customization{}

	selector := []interface{}{
		&customization.ID,
		&customization.Name,
		&customization.LinkURL,
		&customization.Desc,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err

	}
	return customization, nil
}

const UPDATE_BY_ID_STAT = "UPDATE customization SET " +
	" name = $1, \"desc\" = $2, link_url = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, customization *customizationM.Customization) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		customization.Name,
		customization.Desc,
		customization.LinkURL,
		customization.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	tx.Commit()
	return nil
}
