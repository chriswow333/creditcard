package feedback_desc

import (
	"context"

	feedbackM "example.com/creditcard/models/feedback"
	feedbackDescStore "example.com/creditcard/stores/feedback_desc"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
)

type impl struct {
	feedbackDescStore feedbackDescStore.Store
}

func New(
	feedbackDescStore feedbackDescStore.Store,
) Service {
	return &impl{
		feedbackDescStore: feedbackDescStore,
	}
}

func (im *impl) Create(ctx context.Context, feedbackDesc *feedbackM.FeedbackDesc) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)

		return err
	}
	feedbackDesc.ID = id.String()

	if err := im.feedbackDescStore.Create(ctx, feedbackDesc); err != nil {
		logrus.WithFields(logrus.Fields{
			"feedbackStore create": "failed ",
		}).Error(err)
		return err
	}

	return nil
}
func (im *impl) GetAll(ctx context.Context) ([]*feedbackM.FeedbackDesc, error) {

	feedbackDescs, err := im.feedbackDescStore.GetAll(ctx)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"feedbackStore get all": "failed ",
		}).Error(err)
		return nil, err
	}

	return feedbackDescs, nil
}
func (im *impl) GetByID(ctx context.Context, ID string) (*feedbackM.FeedbackDesc, error) {

	feedbackDesc, err := im.feedbackDescStore.GetByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"feedbackStore get by id ": "failed ",
		}).Error(err)
		return nil, err
	}
	return feedbackDesc, nil
}
