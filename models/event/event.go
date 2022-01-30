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

	CardIDs []string `json:"cards"` // 已定義要跑哪幾張卡

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

	StartDate  int64 `json:"startDate,omitempty"`
	EndDate    int64 `json:"endDate,omitempty"`
	UpdateDate int64 `json:"updateDate,omitempty"`

	LinkURL string `json:"linkURL,omitempty"`

	CardRewards []*CardReward `json:"cardRewards"`
}

type CardReward struct {
	RewardType reward.RewardType `json:"rewardType"`
	TotalCost  float64           `json:"totalCost"`

	TotalGetBonus float64 `json:"totalGetBonus"`
	TotalGetCash  float64 `json:"totalGetCash"`
	TotalGetPoint float64 `json:"totalGetPoint"`

	Rewards []*RewardResp `json:"rewards"`
}

type RewardResp struct {
	Name string `json:"name,omitempty"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	Pass bool `json:"pass"`

	RewardOperator constraint.OperatorType `json:"rewardOperator"`
	Constraint     *ConstraintResp         `json:"constraint"`
}

type ConstraintResp struct {
	Pass bool `json:"pass"`

	Name string `json:"name"`

	Feedback *feedback.Feedback `json:"feedback"`

	Matches []string `json:"matches"` // 符合限制的id, ex. supermarket
	Misses  []string `json:"misses"`  // 符合限制的id, ex. supermarket

	ConstraintType constraint.ConstraintType `json:"constraintType"`
	Constraints    []*ConstraintResp         `json:"constraints"`
}
