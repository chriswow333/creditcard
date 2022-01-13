package event

import (
	"example.com/creditcard/models/action"
	"example.com/creditcard/models/constraint"
	"example.com/creditcard/models/cost"
	"example.com/creditcard/models/customization"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
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

	CardIDs []string `json:"cards,omitempty"` // 已定義要跑哪幾張卡

	EffictiveTime int64 `json:"effictiveTime"`

	ActionType action.ActionType `json:"actionType"`

	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`

	Mobilepays []*mobilepay.Mobilepay `json:"mobilpays,omitempty"`

	Customizations []*customization.Customization `json:"customizations,omitempty"`
	// BankAccounts   []*bankaccount.BankAccount     `json:"bankAccounts,omitempty"`
}

type Response struct {
	EventID string      `json:"eventID"`
	Cards   []*CardResp `json:"cards,omitempty"`
}

type CardResp struct {
	Name      string `json:"name,omitempty"`
	Desc      string `json:"desc,omitempty"`
	StartDate int64  `json:"startDate,omitempty"`
	EndDate   int64  `json:"endDate,omitempty"`
	LinkURL   string `json:"linkURL,omitempty"`

	Rewards []*RewardResp `json:"rewards"`
}

type RewardResp struct {
	Pass bool `json:"pass"`

	Cost *cost.Cost `json:"cost"`

	Name     string                  `json:"name,omitempty"`
	Desc     string                  `json:"desc,omitempty"`
	Operator constraint.OperatorType `json:"operator"`

	Constraints []*ConstraintResp `json:"constraints,omitempty"`
}

type ConstraintResp struct {
	Pass bool `json:"pass"`

	ConstraintType constraint.ConstraintType `json:"constraintType"`

	Matches []string `json:"matches,omitempty"` // 符合限制的id, ex. supermarket
	Misses  []string `json:"misses,omitempty"`  // 符合限制的id, ex. supermarket

	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`

	Constraints []*ConstraintResp `json:"constraints,omitempty"`
}
