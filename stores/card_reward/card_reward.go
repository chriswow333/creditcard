package card_reward

import (
	"context"

	cardM "example.com/creditcard/models/card"
)

type Store interface {
	Create(ctx context.Context, cardReward *cardM.CardReward) error
	GetByCardID(ctx context.Context, cardID string) ([]*cardM.CardReward, error)
	GetByID(ctx context.Context, ID string) (*cardM.CardReward, error)
}
