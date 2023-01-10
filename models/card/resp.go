package card

import (
	"example.com/creditcard/models/channel"
	feedbackM "example.com/creditcard/models/feedback"
	"example.com/creditcard/models/reward"
)

type CardResp struct {
	ID     string   `json:"id"`
	BankID string   `json:"bankID"`
	Name   string   `json:"name"`
	Descs  []string `json:"descs"`

	BankName string `json:"bankName"`

	UpdateDate string `json:"updateDate"`

	CardStatus CardStatus `json:"cardStatus"`

	ImagePath string `json:"imagePath"`
	LinkURL   string `json:"linkURL"`

	CardRewardResps []*CardRewardResp `json:"cardRewardResps"`

	OtherRewardResps []*OtherReward `json:"otherRewardResps"`
}

type CardRewardResp struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`

	Title string   `json:"title"`
	Descs []string `json:"descs"`

	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`

	CardRewardOperator CardRewardOperator `json:"cardRewardOperator,omitempty"` // (R0+(R1&(R2|R3)))

	RewardType reward.RewardType `json:"rewardType,omitempty"`
	// CardRewardBonus *CardRewardBonus  `json:"cardRewardBonus"`
	FeedbackBonus *feedbackM.FeedbackBonus `json:"feedbackBonus"` // show 9折優惠 or 10%回饋 等等

	CardRewardLimitTypes []CardRewardLimitType `json:"cardRewardLimitTypes"`

	ConstraintPassLogics []*ConstraintPassLogic `json:"constraintPassLogics"`

	ChannelResps []*channel.ChannelResp `json:"channelResps"`

	// FeedbackDesc *feedback.FeedbackDesc `json:"feedbackDesc,omitempty"`
}

type OtherRewardResp struct {
	Title string   `json:"title"`
	Descs []string `json:"descs"`
}
