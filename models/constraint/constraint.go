package constraint

import (
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
	ID          string `json:"id"`
	PrivilageID string `json:"privilageID"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	Limit *Limit `json:"limit,omitempty"`

	ConstraintBody *ConstraintBody `json:"constraintBody,omitempty"`
}

type ConstraintBody struct {
	ConstraintPayloads []*ConstraintPayload `json:"constraintPayloads,omitempty"`
}

type ConstraintPayload struct {
	Operator OperatorType `json:"operator"`

	ConstraintPayloads []*ConstraintPayload `json:"constraintPayloads,omitempty"`

	Base         []*Base                    `json:"base,omitempty"`
	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays,omitempty"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`
}

type ActionType int32

const (
	Shopping ActionType = iota
	Deposit
	Setting
)

type BaseType int32

const (
	TimeBase BaseType = iota
	MoneyBase
	AccountBase
)

type Unit int32

const (
	DayUint Unit = iota
	AccountUint
	NTD
)

type Base struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Desc       string   `json:"desc"`
	TargetFrom string   `json:"targetFrom"`
	TargetTo   string   `json:"targetTo"`
	BaseType   BaseType `json:"baseType"`
	UnitType   string   `json:"unitType"`

	Action ActionType `json:"actionType"`
}

type Limit struct {
	Max int `json:"max"`
	Min int `json:"min"`
}
