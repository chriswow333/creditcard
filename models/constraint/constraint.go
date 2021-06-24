package constraint

import (
	"example.com/creditcard/models/bonus"
	"example.com/creditcard/models/cost"
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
	RewardID string `json:"rewardID"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	ConstraintPayload *ConstraintPayload `json:"constraintPayload,omitempty"`
}

type ConstraintType int32

const (
	ConstraintPayloadType ConstraintType = iota
	CustomizationType
	TimeIntervalType
	CostLimitType
	BonusLimitType
	MobilepayType
	EcommerceType
	SupermarketType
	OnlinegameType
	StreamingType
)

type ConstraintPayload struct {
	Name           string         `json:"name"`
	Operator       OperatorType   `json:"operator"`
	Descs          []string       `json:"descs"`
	ConstraintType ConstraintType `json:"constraintType"`

	ConstraintPayloads []*ConstraintPayload `json:"constraintPayloads,omitempty"`

	CostLimit  *cost.CostLimit   `json:"costLimit,omitempty"`
	BonusLimit *bonus.BonusLimit `json:"bonusLimit,omitempty"`

	TimeIntervals []*timeinterval.TimeInterval `json:"timeIntervals,omitempty"`

	Customizations []*customization.Customization `json:"customizations,omitempty"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays,omitempty"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`
}
