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

type ConstraintResp struct {
	ConstraintOperatorType ConstraintOperatorType `json:"constraintOperatorType,omitempty"`

	ConstraintType ConstraintType `json:"constraintType,omitempty"`

	ConstraintMappingType ConstraintMappingType `json:"constraintMappingType,omitempty"`

	InnerConstraints []*ConstraintResp `json:"innerConstraintResps,omitempty"`

	TimeIntervals  []*timeinterval.TimeInterval   `json:"timeIntervals,omitempty"`
	Customizations []*customization.Customization `json:"customizations,omitempty"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays,omitempty"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`
}
