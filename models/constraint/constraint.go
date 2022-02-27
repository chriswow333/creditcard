package constraint

import (
	"example.com/creditcard/models/customization"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
	"example.com/creditcard/models/timeinterval"
)

type OperatorType int32

const (
	AndOperator OperatorType = iota
	OrOperator
)

type ConstraintType int32

const (
	InnerConstraintType ConstraintType = iota //  abstract layer, there are several nested layers.
	CustomizationType                         // setting layer, ex. 綁定數位帳號
	TimeIntervalType                          // time range
	MobilepayType                             // ex. line pay, google pay.
	EcommerceType                             // ex. shopee, momo
	SupermarketType                           // ex. px mart
	OnlinegameType                            // ex.
	StreamingType                             // ex. netflix
)

type Constraint struct {
	Descs []string `json:"descs"`

	ConstraintOperator OperatorType `json:"constraintOperator"`

	ConstraintType ConstraintType `json:"constraintType"`

	InnerConstraints []*Constraint `json:"constraints"`

	TimeIntervals  []*timeinterval.TimeInterval   `json:"timeIntervals"`
	Customizations []*customization.Customization `json:"customizations"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames"`
	Streamings   []*streaming.Streaming     `json:"streamings"`
}
