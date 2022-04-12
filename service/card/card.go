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
	GetRespAll(ctx context.Context) ([]*cardM.CardResp, error)
	GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error)
	GetRespByBankID(ctx context.Context, bankID string) ([]*cardM.CardResp, error)

	CreateCardReward(ctx context.Context, cardReward *cardM.CardReward) error
	EvaluateConstraintLogic(ctx context.Context, cardID string, constraintIDs []string) (bool, error)
}
