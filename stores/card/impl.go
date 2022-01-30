package card

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	cardM "example.com/creditcard/models/card"
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

const INSERT_CARD_STAT = "INSERT INTO card " +
	" (\"id\", bank_id, \"name\", \"desc\", start_date, end_date, update_date, link_url) " +
	" VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

func (im *impl) Create(ctx context.Context, card *cardM.Card) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		card.ID,
		card.BankID,
		card.Name,
		card.Desc,
		card.StartDate,
		card.EndDate,
		card.UpdateDate,
		card.LinkURL,
	}

	if _, err := tx.Exec(INSERT_CARD_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE card SET " +
	" bank_id = $1, \"name\" = $2, \"desc\" = $3, start_date = $4, end_date = $5, update_date = $6, link_url = $7 " +
	" where \"id\" = $8"

func (im *impl) UpdateByID(ctx context.Context, card *cardM.Card) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		card.BankID,
		card.Name,
		card.Desc,
		card.StartDate,
		card.EndDate,
		card.UpdateDate,
		card.LinkURL,
		card.ID,
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

const SELECT_STAT = "SELECT \"id\", bank_id, \"name\", \"desc\", start_date, end_date, update_date, link_url " +
	" FROM card " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.Card, error) {

	card := &cardM.Card{}

	selector := []interface{}{
		&card.ID,
		&card.BankID,
		&card.Name,
		&card.Desc,
		&card.StartDate,
		&card.EndDate,
		&card.UpdateDate,
		&card.LinkURL,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err

	}
	return card, nil
}

const SELECT_ALL_STAT = "SELECT \"id\", bank_id, \"name\", \"desc\", start_date, end_date, update_date, link_url " +
	" FROM card"

func (im *impl) GetAll(ctx context.Context) ([]*cardM.Card, error) {

	cards := []*cardM.Card{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		card := &cardM.Card{}
		selector := []interface{}{
			&card.ID,
			&card.BankID,
			&card.Name,
			&card.Desc,
			&card.StartDate,
			&card.EndDate,
			&card.UpdateDate,
			&card.LinkURL,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}

const SELECT_BY_BANKID_STAT = "SELECT \"id\", bank_id, \"name\", \"desc\", start_date, end_date, update_date, link_url " +
	" FROM card " +
	" WHERE \"bank_id\"=$1"

func (im *impl) GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error) {
	cards := []*cardM.Card{}

	conditions := []interface{}{
		bankID,
	}
	rows, err := im.psql.Query(SELECT_BY_BANKID_STAT, conditions...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		card := &cardM.Card{}
		selector := []interface{}{
			&card.ID,
			&card.BankID,
			&card.Name,
			&card.Desc,
			&card.StartDate,
			&card.EndDate,
			&card.UpdateDate,
			&card.LinkURL,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil

}
