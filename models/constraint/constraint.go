package constraint

import (
	"example.com/creditcard/models/customization"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/feedback"
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

type Constraint struct {
	ID       string `json:"id"`
	RewardID string `json:"rewardID,omitempty"`

	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`

	ConstraintPayload *ConstraintPayload `json:"constraintPayload,omitempty"`
}

type ConstraintType int32

const (
	ConstraintPayloadType ConstraintType = iota //  abstract layer, there are several nested layers.
	CustomizationType                           // setting layer, ex. 綁定數位帳號
	TimeIntervalType                            // time range
	MobilepayType                               // ex. line pay, google pay.
	EcommerceType                               // ex. shopee, momo
	SupermarketType                             // ex. px mart
	OnlinegameType                              // ex.
	StreamingType                               // ex. netflix
)

type ConstraintPayload struct {
	Name string `json:"name"`
	Desc string `json:"desc"`

	ConstraintOperator OperatorType `json:"constraintOperator"`

	Feedback *feedback.Feedback `json:"feedback"`

	ConstraintType ConstraintType `json:"constraintType"`

	ConstraintPayloads []*ConstraintPayload           `json:"constraintPayloads"`
	TimeIntervals      []*timeinterval.TimeInterval   `json:"timeIntervals"`
	Customizations     []*customization.Customization `json:"customizations"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames"`
	Streamings   []*streaming.Streaming     `json:"streamings"`
}
