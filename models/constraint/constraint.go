package constraint

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"example.com/creditcard/models/customization"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
	"example.com/creditcard/models/timeinterval"
	"example.com/creditcard/service/constraint"
)

type ConstraintOperatorType int32

const (
	AND ConstraintOperatorType = iota + 1
	OR
)

type ConstraintType int32

const (
	InnerConstraintType ConstraintType = iota + 1 //  abstract layer, there are several nested layers.
	CustomizationType                             // setting layer, ex. 綁定數位帳號
	TimeIntervalType                              // time range
	MobilepayType                                 // ex. line pay, google pay.
	EcommerceType                                 // ex. shopee, momo
	SupermarketType                               // ex. px mart
	OnlinegameType                                // ex.
	StreamingType                                 // ex. netflix
)

type Constraint struct {
	ConstraintOperatorType ConstraintOperatorType `json:"constraintOperatorType,omitempty"`

	ConstraintType ConstraintType `json:"constraintType,omitempty"`

	InnerConstraints []*Constraint `json:"innerConstraints,omitempty"`

	TimeIntervals  []string `json:"timeIntervals,omitempty"`
	Customizations []string `json:"customizations,omitempty"`

	Mobilepays   []string `json:"mobilepays,omitempty"`
	Ecommerces   []string `json:"ecommerces,omitempty"`
	Supermarkets []string `json:"supermarkets,omitempty"`
	Onlinegames  []string `json:"onlinegames,omitempty"`
	Streamings   []string `json:"streamings,omitempty"`
}

type ConstraintEventJudgeType int32

func TransferConstraintResp(ctx context.Context, constraint *Constraint, constraintSvc constraint.Service) (*ConstraintResp, error) {

	constraintResp := &ConstraintResp{
		ConstraintOperatorType: constraint.ConstraintOperatorType,
	}

	switch constraint.ConstraintType {
	case InnerConstraintType:
		innerConstraints := []*ConstraintResp{}

		for _, c := range constraint.InnerConstraints {
			innerConstraint, err := TransferConstraintResp(ctx, c, constraintSvc)
			if err != nil {
				logrus.Error("Failed to get InnerConstraint")
				return nil, err
			}
			innerConstraints = append(innerConstraints, innerConstraint)
		}

		constraintResp.ConstraintType = InnerConstraintType
		constraintResp.InnerConstraints = innerConstraints

		break
	case CustomizationType:
		customizations := []*customization.Customization{}
		for _, ID := range constraint.Customizations {
			customization, err := constraintSvc.GetCustomizationByID(ctx, ID)
			if err != nil {
				logrus.Error("Failed to get customization")
				return nil, err
			}
			customizations = append(customizations, customization)
		}
		constraintResp.ConstraintType = CustomizationType
		constraintResp.Customizations = customizations

		break
	case TimeIntervalType:
		timeIntervals := []*timeinterval.TimeInterval{}

		for _, ID := range constraint.TimeIntervals {
			timeInterval, err := constraintSvc.GetTimeInterval(ctx, ID)
			if err != nil {
				logrus.Error("Failed to get GetTimeInterval")
				return nil, err
			}
			timeIntervals = append(timeIntervals, timeInterval)
		}
		constraintResp.ConstraintType = TimeIntervalType
		constraintResp.TimeIntervals = timeIntervals
		break
	case MobilepayType:
		mobilepays := []*mobilepay.Mobilepay{}

		for _, ID := range constraint.Mobilepays {
			mobilepay, err := constraintSvc.GetMobilepay(ctx, ID)
			if err != nil {
				logrus.Error("Failed to get GetMobilepay")
				return nil, err
			}

			mobilepays = append(mobilepays, mobilepay)
		}
		constraintResp.ConstraintType = MobilepayType
		constraintResp.Mobilepays = mobilepays
		break
	case EcommerceType:
		ecommerces := []*ecommerce.Ecommerce{}

		for _, ID := range constraint.Ecommerces {
			ecommerce, err := constraintSvc.GetEcommerce(ctx, ID)
			if err != nil {
				logrus.Error("Failed to get GetMobilepay")
				return nil, err
			}
			ecommerces = append(ecommerces, ecommerce)
		}
		constraintResp.ConstraintType = EcommerceType
		constraintResp.Ecommerces = ecommerces
		break
	case SupermarketType:

		supermarkets := []*supermarket.Supermarket{}

		for _, ID := range constraint.Supermarkets {
			supermarket, err := constraintSvc.GetSupermarket(ctx, ID)
			if err != nil {
				logrus.Error("Failed to get GetSupermarket")
				return nil, err
			}
			supermarkets = append(supermarkets, supermarket)
		}
		constraintResp.ConstraintType = SupermarketType
		constraintResp.Supermarkets = supermarkets
		break
	case OnlinegameType:
		onlinegames := []*onlinegame.Onlinegame{}

		for _, ID := range constraint.Onlinegames {
			onlinegame, err := constraintSvc.GetOnlinegame(ctx, ID)
			if err != nil {
				logrus.Error("Failed to get GetSupermarket")
				return nil, err
			}
			onlinegames = append(onlinegames, onlinegame)
		}

		constraintResp.ConstraintType = OnlinegameType
		constraintResp.Onlinegames = onlinegames
		break
	case StreamingType:
		streamings := []*streaming.Streaming{}
		for _, ID := range constraint.Streamings {
			streaming, err := constraintSvc.GetStreaming(ctx, ID)
			if err != nil {
				logrus.Error("Failed to get GetStreaming")
				return nil, err
			}
			streamings = append(streamings, streaming)
		}
		constraintResp.ConstraintType = StreamingType
		constraintResp.Streamings = streamings
		break
	default:
		logrus.Error("No matching ConstraintType")
		return nil, errors.New("No matching ConstraintType")
	}

	return constraintResp, nil

}
