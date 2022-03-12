package event

import (
	"example.com/creditcard/models/action"
	"example.com/creditcard/models/card"
	"example.com/creditcard/models/customization"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/reward"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
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

	RewardType reward.RewardType `json:"rewardType"`

	EffictiveTime int64 `json:"effictiveTime"`

	ActionType action.ActionType `json:"actionType"`

	DefaultCustomization bool `json:"defaultCustomization"`

	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames"`
	Streamings   []*streaming.Streaming     `json:"streamings"`

	Mobilepays []*mobilepay.Mobilepay `json:"mobilpays"`

	Customizations []*customization.Customization `json:"customizations"`
}

type Response struct {
	EventID string           `json:"eventID"`
	Cards   []*card.CardResp `json:"cards"`
}
