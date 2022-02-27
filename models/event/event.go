package event

import (
	"example.com/creditcard/models/action"
	"example.com/creditcard/models/constraint"
	"example.com/creditcard/models/customization"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/feedback"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/reward"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
)

type CashType int32

const (
	NTD CashType = iota
	USD
	BONUS
)

type Event struct {
	ID string `json:"id"`

	Cash     float64  `json:"cash"`
	CashType CashType `json:"cashType"`

	RewardType reward.RewardType `json:"rewardType"`

	EffictiveTime int64 `json:"effictiveTime"`

	ActionType action.ActionType `json:"actionType"`

	DefaultCustomization bool `json:"defaultCustomization"`

	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames"`
	Streamings   []*streaming.Streaming     `json:"streamings"`

	Mobilepays []*mobilepay.Mobilepay `json:"mobilpays"`

	Customizations []*customization.Customization `json:"customizations"`
}

type Response struct {
	EventID string      `json:"eventID"`
	Cards   []*CardResp `json:"cards"`
}

type CardResp struct {
	ID     string `json:"id,omitempty"`
	BankID string `json:"bankID,omitempty"`

	Name string `json:"name,omitempty"`

	StartDate  string `json:"startDate,omitempty"`
	EndDate    string `json:"endDate,omitempty"`
	UpdateDate string `json:"updateDate,omitempty"`

	ImagePath string `json:"imagePath"`
	LinkURL   string `json:"linkURL,omitempty"`

	InCashRewardResp *InCashRewardResp `json:"inCashRewardResp"`
}

type InCashRewardResp struct {
	FeedReturn *feedback.FeedReturn `json:"feedReturn"`

	RewardResps []*RewardResp `json:"rewardResps"`
}

type RewardResp struct {
	ID string `json:"id"`

	Order int32 `json:"order"`

	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`

	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	UpdateDate string `json:"updateDate"`

	FeedReturn      *feedback.FeedReturn   `json:"feedReturn"`
	PayloadOperator reward.PayloadOperator `json:"payloadOperator"`
	PayloadResps    []*PayloadResp         `json:"payloadResps"`
}

type PayloadResp struct {
	Pass bool `json:"pass"`

	Feedback   *feedback.Feedback   `json:"feedback"`
	FeedReturn *feedback.FeedReturn `json:"feedReturn"`

	ConstraintResp *ConstraintResp `json:"constraintResp"`
}

type ConstraintResp struct {
	Pass bool `json:"pass"`

	Matches []string `json:"matches"` // 符合限制的id, ex. supermarket
	Misses  []string `json:"misses"`  // 符合限制的id, ex. supermarket

	ConstraintType constraint.ConstraintType `json:"constraintType"`
	Constraints    []*ConstraintResp         `json:"constraints,omitempty"`
}
