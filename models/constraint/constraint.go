package constraint

import (
	"example.com/creditcard/models/base"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
)

type OperatorType int32

const (
	AndOperator OperatorType = iota
	OrOperator
)

type Constraint struct {
	ID       string `json:"id"`
	RewardID string `json:"rewardID"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	Limit *Limit `json:"limit,omitempty"`

	ConstraintBody *ConstraintBody `json:"constraintBody,omitempty"`
}

type Limit struct {
	Max int `json:"max"`
	Min int `json:"min"`
}

type ConstraintBody struct {
	ConstraintPayloads []*ConstraintPayload `json:"constraintPayloads,omitempty"`
}

type ConstraintPayload struct {
	Operator OperatorType `json:"operator"`

	ConstraintPayloads []*ConstraintPayload `json:"constraintPayloads,omitempty"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays,omitempty"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`

	TimeBases    []*base.TimeBase    `json:"timeBases,omitempty"`
	AccountBases []*base.AccountBase `json:"accountBases,omitempty"`
	MoneyBases   []*base.MoneyBase   `json:"moneyBases,omitempty"`
}
