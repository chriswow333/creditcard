package constraint

import (
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
)

type Constraint struct {
	ID          string `json:"id"`
	PrivilageID string `json:"privilageID"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`

	Operator  int32 `json:"operator"`
	StartDate int64 `json:"startDate"`
	EndDate   int64 `json:"endDate"`

	Constraint []*Constraint `json:"constraints,omitempty"`

	Limit *Limit `json:"limit"`

	Mobilepaies  []*mobilepay.Mobilepay     `json:"mobilaies,omitempty"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"Ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegamis,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`

	UpdateDate int64 `json:"updateDate"`
}

type Limit struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
}
