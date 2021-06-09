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

	Operator  OperatorType `json:"operator"`
	StartDate int64        `json:"startDate"`
	EndDate   int64        `json:"endDate"`

	UpdateDate int64 `json:"updateDate"`

	Constraint []*Constraint `json:"constraints,omitempty"`

	Limit *Limit `json:"limit,omitempty"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays,omitempty"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegamis,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`
}

type Limit struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
}
