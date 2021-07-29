package card

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	cardM "example.com/creditcard/app/view_card/models/card"
	"example.com/creditcard/app/view_card/utils/conn"
)

type impl struct {
	dig.In

	psql *pgx.ConnPool

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

const INSERT_STAT = "INSERT INTO card " +
	" (\"id\", \"name\", icon, bank_id, start_time, end_time, " +
	" max_point, feature_desc, applicant_qualifications, update_date) " +
	" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

func (im *impl) Create(ctx context.Context, conn *conn.Connection, card *cardM.Card) error {

	updater := []interface{}{
		card.ID,
		card.Name,
		card.Icon,
		card.BankID,
		card.ValidateTime.StartTime,
		card.ValidateTime.EndTime,
		card.MaxPoint,
		card.FeatureDesc,
		card.ApplicantQualifications,
		card.UpdateDate,
	}

	if err := im.connService.Exec(conn, INSERT_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}
	return nil

}

const SELECT_STAT = "SELECT \"id\", \"name\", icon, bank_id, " +
	" start_time, end_time, max_point, feature_desc, applicant_qualifications, update_date " +
	" FROM card " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.Card, error) {

	card := &cardM.Card{}

	selector := []interface{}{
		&card.ID,
		&card.Name,
		&card.Icon,
		&card.BankID,
		&card.ValidateTime.StartTime,
		&card.ValidateTime.EndTime,
		&card.MaxPoint,
		&card.FeatureDesc,
		&card.ApplicantQualifications,
		&card.ApplicantQualifications,
		&card.UpdateDate,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err

	}
	return card, nil
}

const UPDATE_BY_ID_STAT = "UPDATE card SET " +
	" \"name\" = $1, icon = $2, bank_id = $3, " +
	" start_time = $4, end_time = $5, max_point = $6, " +
	" feature_desc = $7, applicant_qualifications = $8, update_date = $9 " +
	" where \"id\" = $10"

func (im *impl) UpdateByID(ctx context.Context, conn *conn.Connection, card *cardM.Card) error {

	updater := []interface{}{
		card.Name,
		card.Icon,
		card.BankID,
		card.ValidateTime.StartTime,
		card.ValidateTime.EndTime,
		card.MaxPoint,
		card.FeatureDesc,
		card.ApplicantQualifications,
		card.UpdateDate,
		card.ID,
	}

	if err := im.connService.Exec(conn, UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", icon, bank_id, " +
	" start_time, end_time, max_point, feature_desc, applicant_qualifications, update_date " +
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
			&card.Name,
			&card.Icon,
			&card.BankID,
			&card.ValidateTime.StartTime,
			&card.ValidateTime.EndTime,
			&card.MaxPoint,
			&card.FeatureDesc,
			&card.ApplicantQualifications,
			&card.UpdateDate,
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

const SELECT_BANKID_STAT = "SELECT \"id\", \"name\", icon, bank_id, " +
	" start_time, end_time, max_point, feature_desc, applicant_qualifications, update_date " +
	" FROM card " +
	" WHERE bank_id = $1"

func (im *impl) GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error) {

	cards := []*cardM.Card{}

	conditions := []interface{}{
		bankID,
	}
	rows, err := im.psql.Query(SELECT_BANKID_STAT, conditions...)

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
			&card.Name,
			&card.Icon,
			&card.BankID,
			&card.ValidateTime.StartTime,
			&card.ValidateTime.EndTime,
			&card.MaxPoint,
			&card.FeatureDesc,
			&card.ApplicantQualifications,
			&card.ApplicantQualifications,
			&card.UpdateDate,
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
