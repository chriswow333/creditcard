package card

import (
	"example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
)

type CardEventResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	BankID   string `json:"bankID"`
	BankName string `json:"bankName"`

	UpdateDate string `json:"updateDate"`

	ImagePath string `json:"imagePath,omitempty"`

	CardRewardEventResps []*CardRewardEventResp `json:"cardRewardEventResps"`
}

type CardRewardEventResp struct {
	ID string `json:"id"`

	Title string   `json:"title"`
	Descs []string `json:"descs"`

	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`

	CardRewardOperator CardRewardOperator `json:"cardRewardOperator,omitempty"`
	RewardType         rewardM.RewardType `json:"rewardType,omitempty"`

	CardRewardBonus *CardRewardBonus `json:"cardRewardBonus,omitempty"`

	CardRewardLimitTypes []CardRewardLimitType `json:"cardRewardLimitTypes"`

	FeedReturn *feedback.FeedReturn `json:"feedReturn,omitempty"`

	// RewardEventResps []*rewardM.RewardEventResp `json:"rewardEventResps,omitempty"`
}
