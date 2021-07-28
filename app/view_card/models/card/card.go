package card

import (
	"example.com/creditcard/app/view_card/models/common"
	"example.com/creditcard/app/view_card/models/reward"
)

type FeatureType int32

const (
	ECommerce FeatureType = iota
	Supremarket
	Delivery
	Fee
	Transport
)

type Card struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`

	BankID string `json:"bankID"`

	MaxPoint float64 `json:"maxPoint"`

	Feature     *Feature `json:"feature"`
	FeatureDesc string   `json:"featureDesc"`

	ValidateTime            common.ValidateTime `json:"validateDate"`
	ApplicantQualifications []string            `json:"applicantQualifications"`

	UpdateDate int64 `json:"updateDate"`

	Rewards []*reward.Reward `json:"rewards"`
}

type Feature struct {
	FeatureTypes []FeatureType `json:"featureTypes"`
}
