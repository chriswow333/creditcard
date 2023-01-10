package feedback_desc

import (
	"github.com/jackc/pgx"
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

// const INSERT_STAT = "INSERT INTO feedback_desc " +
// 	"(\"id\", \"type\", \"calculate_type\", \"name\", \"descs\") VALUES ($1, $2, $3, $4, $5)"

// func (im *impl) Create(ctx context.Context, feedbackDesc *feedback.FeedbackDesc) error {

// 	tx, err := im.psql.Begin()
// 	if err != nil {
// 		logrus.WithFields(logrus.Fields{
// 			"msg": "",
// 		}).Error(err)
// 		return err
// 	}

// 	defer tx.Rollback()

// 	updater := []interface{}{
// 		feedbackDesc.ID,
// 		feedbackDesc.FeedbackType,
// 		feedbackDesc.FeedbackCalculateType,
// 		feedbackDesc.Name,
// 		feedbackDesc.Descs,
// 	}

// 	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
// 		logrus.WithFields(logrus.Fields{
// 			"": "",
// 		}).Fatal(err)

// 		return err
// 	}

// 	tx.Commit()

// 	return nil
// }

// const SELECT_ALL_STAT = "SELECT \"id\", \"type\",  \"calculate_type\", \"name\", \"descs\" " +
// 	" FROM feedback_desc "

// func (im *impl) GetAll(ctx context.Context) ([]*feedback.FeedbackDesc, error) {

// 	feedbackDescs := []*feedback.FeedbackDesc{}

// 	rows, err := im.psql.Query(SELECT_ALL_STAT)
// 	if err != nil {
// 		logrus.WithFields(logrus.Fields{
// 			"": "",
// 		}).Error(err)
// 		return nil, err
// 	}

// 	for rows.Next() {

// 		feedbackDesc := &feedback.FeedbackDesc{}
// 		selector := []interface{}{
// 			&feedbackDesc.ID,
// 			&feedbackDesc.FeedbackType,
// 			&feedbackDesc.FeedbackCalculateType,
// 			&feedbackDesc.Name,
// 			&feedbackDesc.Descs,
// 		}

// 		if err := rows.Scan(selector...); err != nil {
// 			logrus.WithFields(logrus.Fields{
// 				"": "",
// 			}).Error(err)
// 			return nil, err
// 		}

// 		feedbackDescs = append(feedbackDescs, feedbackDesc)
// 	}

// 	return feedbackDescs, nil
// }

// const SELECT_BY_TYPE_STAT = "SELECT \"id\", \"type\", \"calculate_type\", \"name\", \"descs\" " +
// 	" FROM feedback_desc WHERE \"id\" = $1  "

// func (im *impl) GetByID(ctx context.Context, ID string) (*feedback.FeedbackDesc, error) {

// 	feedbackDesc := &feedback.FeedbackDesc{}

// 	selector := []interface{}{
// 		&feedbackDesc.ID,
// 		&feedbackDesc.FeedbackType,
// 		&feedbackDesc.FeedbackCalculateType,
// 		&feedbackDesc.Name,
// 		&feedbackDesc.Descs,
// 	}

// 	if err := im.psql.QueryRow(SELECT_BY_TYPE_STAT, ID).Scan(selector...); err != nil {
// 		logrus.Error(err)
// 		return nil, err
// 	}

// 	return feedbackDesc, nil
// }
