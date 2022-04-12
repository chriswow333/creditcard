package card

import (
	"example.com/creditcard/models/feedback"
	"example.com/creditcard/models/reward"
	rewardM "example.com/creditcard/models/reward"
)

type CardResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	BankID   string `json:"bankID"`
	BankName string `json:"bankName"`

	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	UpdateDate string `json:"updateDate"`

	ImagePath string `json:"imagePath"`
	LinkURL   string `json:"linkURL"`

	ConstraintPassLogic string `json:"constraintPassLogic"`

	CardRewardResps []*CardRewardResp `json:"cardRewardResps"`
}

type CardRewardResp struct {
	ID             string `json:"id"`
	CardID         string `json:"cardID"`
	CardRewardDesc string `json:"cardRewardDesc"`

	CardRewardOperator CardRewardOperator `json:"cardRewardOperator,omitempty"`
	RewardType         reward.RewardType  `json:"rewardType,omitempty"`

	ConstraintPassLogic string `json:"constraintPassLogic"`

	Feedback *feedback.Feedback `json:"feedback,omitempty"`

	RewardResps []*rewardM.RewardResp `json:"rewardResps,omitempty"`
}
