package feedback_desc

import (
	"context"

	"example.com/creditcard/models/feedback"
)

type Store interface {
	Create(ctx context.Context, feedbackTypeDesc *feedback.FeedbackDesc) error
	GetAll(ctx context.Context) ([]*feedback.FeedbackDesc, error)
	GetByID(ctx context.Context, ID string) (*feedback.FeedbackDesc, error)
}
