package card

import (
	"example.com/creditcard/models/feedback"
	"example.com/creditcard/models/reward"
	rewardM "example.com/creditcard/models/reward"
)

type CardEventResp struct {
	ID string `json:"id"`

	BankID string `json:"bankID"`

	CardRewardEventResp *CardRewardEventResp `json:"cardRewardEventResp"`
}

type CardRewardEventResp struct {
	ID string `json:"id"`

	CardID string `json:"cardID"`

	CardRewardOperator CardRewardOperator `json:"cardRewardOperator,omitempty"`
	RewardType         reward.RewardType  `json:"rewardType,omitempty"`

	CardRewardEventJudgeType CardRewardEventJudgeType `json:"cardRewardEventJudgeType,omitempty"`

	FeedReturn *feedback.FeedReturn `json:"feedReturn,omitempty"`

	RewardEventResps []*rewardM.RewardEventResp `json:"rewardEventResps,omitempty"`
}
