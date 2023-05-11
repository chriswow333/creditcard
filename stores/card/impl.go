package card

import (
	"context"
	"runtime/debug"
	"strings"

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
	" (\"id\", bank_id, \"name\", \"descs\", update_date, image_path, link_url, card_status, other_reward) " +
	" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

func (im *impl) Create(ctx context.Context, card *cardM.Card) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		card.ID,
		card.BankID,
		card.Name,
		card.Descs,
		card.UpdateDate,
		card.ImagePath,
		card.LinkURL,
		card.CardStatus,
		card.OtherRewards,
	}

	if _, err := tx.Exec(INSERT_CARD_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE card SET " +
	" bank_id = $1, \"name\" = $2, \"descs\" = $3, update_date = $4, " +
	" image_path = $5, link_url = $6, card_status = $67 " +
	" other_reward = $8 " +
	" where \"id\" = $9"

func (im *impl) UpdateByID(ctx context.Context, card *cardM.Card) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		card.BankID,
		card.Name,
		card.Descs,
		card.UpdateDate,
		card.ImagePath,
		card.LinkURL,
		card.CardStatus,
		card.CardRewards,
		card.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	tx.Commit()
	return nil
}

const SELECT_STAT = "SELECT \"id\", bank_id, \"name\", \"descs\", update_date, " +
	" image_path, link_url, card_status, other_reward " +
	" FROM card " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.Card, error) {
	card := &cardM.Card{}

	selector := []interface{}{
		&card.ID,
		&card.BankID,
		&card.Name,
		&card.Descs,
		&card.UpdateDate,
		&card.ImagePath,
		&card.LinkURL,
		&card.CardStatus,
		&card.OtherRewards,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err

	}
	return card, nil
}

const SELECT_ALL_STAT = "SELECT \"id\", bank_id, \"name\", \"descs\", update_date, " +
	" image_path, link_url, card_status, other_reward " +
	" FROM card"

func (im *impl) GetAll(ctx context.Context) ([]*cardM.Card, error) {

	cards := []*cardM.Card{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		card := &cardM.Card{}
		selector := []interface{}{
			&card.ID,
			&card.BankID,
			&card.Name,
			&card.Descs,
			&card.UpdateDate,
			&card.ImagePath,
			&card.LinkURL,
			&card.CardStatus,
			&card.OtherRewards,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}

const SELECT_BY_BANKID_STAT = "SELECT \"id\", bank_id, \"name\", \"descs\", update_date, " +
	" iamge_path, link_url, card_status, other_reward " +
	" FROM card " +
	" WHERE \"bank_id\"=$1"

func (im *impl) GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error) {

	cards := []*cardM.Card{}

	conditions := []interface{}{
		bankID,
	}

	rows, err := im.psql.Query(SELECT_BY_BANKID_STAT, conditions...)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		card := &cardM.Card{}
		selector := []interface{}{
			&card.ID,
			&card.BankID,
			&card.Name,
			&card.Descs,
			&card.UpdateDate,
			&card.ImagePath,
			&card.LinkURL,
			&card.CardStatus,
			&card.OtherRewards,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil

}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", bank_id, \"name\", \"descs\", update_date, " +
	" image_path, link_url, card_status, other_reward " +
	" FROM card WHERE \"name\" ~* $1 limit 20"

func (im *impl) FindByLike(ctx context.Context, likes []string) ([]*cardM.Card, error) {

	cards := []*cardM.Card{}
	like := strings.Join(likes, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, like)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		card := &cardM.Card{}

		selector := []interface{}{
			&card.ID,
			&card.BankID,
			&card.Name,
			&card.Descs,
			&card.UpdateDate,
			&card.ImagePath,
			&card.LinkURL,
			&card.CardStatus,
			&card.OtherRewards,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}
