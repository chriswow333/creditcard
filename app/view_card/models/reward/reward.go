package reward

import (
	"example.com/creditcard/app/view_card/models/common"
	"example.com/creditcard/app/view_card/models/task"
)

/*
	信用卡各項回饋
*/

type Reward struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	CardID string `json:"cardID"`
	Desc   string `json:"desc"`

	RewardType common.RewardType `json:"rewardType"`

	OperatorType common.OperatorType `json:"operatorType"`

	ValidateTime *common.ValidateTime `json:"validateDate"`

	TotalPoint float64 `json:"totalPoint"`

	UpdateDate int64 `json:"updateDate"`
}

type Repr struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	CardID string `json:"cardID"`
	Desc   string `json:"desc"`

	RewardType common.RewardType `json:"rewardType"`

	OperatorType common.OperatorType `json:"operatorType"`

	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTIme"`

	TotalPoint float64 `json:"totalPoint"`

	TaskReprs []*task.Repr `json:"tasks"`
}
