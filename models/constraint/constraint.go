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

type Constraint struct {
	ID       string `json:"id"`
	RewardID string `json:"rewardID,omitempty"`
	Name     string `json:"name,omitempty"`
	Desc     string `json:"desc,omitempty"`

	StartDate  int64 `json:"startDate,omitempty"`
	EndDate    int64 `json:"endDate,omitempty"`
	UpdateDate int64 `json:"updateDate,omitempty"`

	ConstraintPayload *ConstraintPayload `json:"constraintPayload,omitempty"`
}

type ConstraintType int32

const (
	ConstraintPayloadType ConstraintType = iota
	CustomizationType
	TimeIntervalType
	MobilepayType
	EcommerceType
	SupermarketType
	OnlinegameType
	StreamingType
)

type ConstraintPayload struct {
	Name     string       `json:"name,omitempty"`
	Operator OperatorType `json:"operator"`
	Desc     string       `json:"desc,omitempty"`

	ConstraintType ConstraintType `json:"constraintType"`

	ConstraintPayloads []*ConstraintPayload `json:"constraintPayloads,omitempty"`

	TimeIntervals []*timeinterval.TimeInterval `json:"timeIntervals,omitempty"`

	Customizations []*customization.Customization `json:"customizations,omitempty"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays,omitempty"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`
}
