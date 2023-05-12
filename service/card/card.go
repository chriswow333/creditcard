package card

import (
	"context"

	cardM "example.com/creditcard/models/card"
)

type Service interface {
	Create(ctx context.Context, card *cardM.Card) error
	GetByID(ctx context.Context, ID string) (*cardM.Card, error)

	GetRespByID(ctx context.Context, ID string) (*cardM.CardResp, error)

	UpdateByID(ctx context.Context, card *cardM.Card) error

	GetAll(ctx context.Context) ([]*cardM.Card, error)

	GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error)

	CreateCardReward(ctx context.Context, cardReward *cardM.CardReward) error
	EvaluateConstraintLogic(ctx context.Context, cardID string, constraintIDs []string) (bool, string, error)

	FindByLike(ctx context.Context, likes []string) ([]*cardM.CardResp, error)
}
