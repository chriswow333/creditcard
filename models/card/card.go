package card

import (
	feedbackM "example.com/creditcard/models/feedback"
	"example.com/creditcard/models/reward"
	rewardM "example.com/creditcard/models/reward"
)

type CardStatus int32

const (
	INACTIVE CardStatus = iota
	ACTIVE
)

type Card struct {
	ID     string   `json:"id"`
	BankID string   `json:"bankID"`
	Name   string   `json:"name,omitempty"`
	Descs  []string `json:"descs"`

	UpdateDate int64 `json:"updateDate,omitempty"`

	CardStatus CardStatus `json:"cardStatus"`

	ImagePath string `json:"imagePath,omitempty"`
	LinkURL   string `json:"linkURL,omitempty"`

	CardRewards  []*CardReward  `json:"cardRewards,omitempty"`
	OtherRewards []*OtherReward `json:"otherRewards,omitempty"`
}

type CardRewardOperator int32

const (
	ADD CardRewardOperator = iota + 1
	MAXONE
)

type CardReward struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`

	Title string   `json:"title"`
	Descs []string `json:"descs"`

	StartDate int64 `json:"startDate,omitempty"`
	EndDate   int64 `json:"endDate,omitempty"`

	CardRewardOperator CardRewardOperator `json:"cardRewardOperator,omitempty"` // (R0+(R1&(R2|R3)))
	RewardType         reward.RewardType  `json:"rewardType,omitempty"`

	FeedbackBonus *feedbackM.FeedbackBonus `json:"feedbackBonus"` // show 9折優惠 or 10%回饋 等等
	// FeedbackDescID string `json:"feedbackDescID"`

	CardRewardLimitTypes []CardRewardLimitType `json:"cardRewardLimitTypes"`

	ConstraintPassLogics []*ConstraintPassLogic `json:"constraintPassLogics"`

	Rewards []*rewardM.Reward `json:"rewards,omitempty"`
}

type CardRewardLimitType int32

const (
	QUANTITY CardRewardLimitType = iota + 1 // 限量
	DURATION                                // 限時
	REGISTER                                // 登入
	NEWONE                                  // 新戶
	TIMES                                   // 次數
	CARD                                    // 卡片限定
)

type OtherReward struct {
	Title string   `json:"title"`
	Descs []string `json:"descs"`
}

type ConstraintPassLogic struct {
	Logic   string `json:"logic"`
	Message string `json:"message"`
}
