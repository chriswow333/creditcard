package event

import (
	"example.com/creditcard/models/action"
	"example.com/creditcard/models/card"
	"example.com/creditcard/models/reward"
)

type CashType int32

const (
	NTD CashType = iota
	USD
	BONUS
)

type Event struct {
	ID string `json:"id"`

	Cash     float64  `json:"cash"`
	CashType CashType `json:"cashType"`

	CardIDs []string `json:"cardIDs"`

	RewardType reward.RewardType `json:"rewardType"`

	EffectiveTime int64 `json:"effectiveTime"`

	ActionType action.ActionType `json:"actionType"`

	DefaultCustomization bool `json:"defaultCustomization"`

	ConstraintIDs []string `json:"constraintIDs"`

	Ecommerces   []string `json:"ecommerces"`
	Supermarkets []string `json:"supermarkets"`
	Onlinegames  []string `json:"onlinegames"`
	Streamings   []string `json:"streamings"`

	Mobilepays []string `json:"mobilepays"`

	Customizations []string `json:"customizations"`
}

type Response struct {
	EventID        string                `json:"eventID"`
	CardEventResps []*card.CardEventResp `json:"cardEventResps"`
}
