package card

import (
	"example.com/creditcard/app/view_card/models/common"
)

type Card struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`

	BankID string `json:"bankID"`

	MaxPoint float64 `json:"maxPoint"`

	FeatureDesc string `json:"featureDesc"`

	ValidateTime            *common.ValidateTime `json:"validateDate"`
	ApplicantQualifications []string             `json:"applicantQualifications"`

	UpdateDate int64 `json:"updateDate"`
}

type Feature struct {
	FeatureTypes []common.FeatureType `json:"featureTypes"`
}

type Repr struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Icon     string  `json:"icon"`
	BankID   string  `json:"bankID"`
	MaxPoint float64 `json:"maxPoint"`

	Features    []common.FeatureType `json:"features"`
	FeatureDesc string               `json:"featureDesc"`

	StartTime int64 `json:"startDate"`
	EndTime   int64 `json:"endDate"`

	ApplicantQualifications []string `json:"applicantQualifications"`
}
