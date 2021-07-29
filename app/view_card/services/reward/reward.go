package reward

import (
	"context"

	rewardM "example.com/creditcard/app/view_card/models/reward"
)

type Service interface {
	Create(ctx context.Context, rewardRepr *rewardM.Repr) error
	GetByID(ctx context.Context, ID string) (*rewardM.Repr, error)
	UpdateByID(ctx context.Context, rewardRepr *rewardM.Repr) error
	GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Repr, error)
}
