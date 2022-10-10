package feedback_desc

import (
	"context"

	feedbackM "example.com/creditcard/models/feedback"
)

type Service interface {
	Create(ctx context.Context, feedbackDesc *feedbackM.FeedbackDesc) error
	GetAll(ctx context.Context) ([]*feedbackM.FeedbackDesc, error)
	GetByID(ctx context.Context, ID string) (*feedbackM.FeedbackDesc, error)
}
