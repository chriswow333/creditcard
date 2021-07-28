package reward

import (
	"example.com/creditcard/app/view_card/models/common"
)

type RewardType int32

const (
	DomesticCache RewardType = iota // 國內現金回饋
	AbroadCache                     // 國外現金回饋
	Point                           // 點數回饋
	Limited                         // 限量優惠
	FirstGift                       // 首刷禮
)

/*
	信用卡各項回饋
*/

type Reward struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	CardID string `json:"cardID"`
	Desc   string `json:"desc"`

	RewardType RewardType `json:"rewardType"`

	OperatorType common.OperatorType `json:"operatorType"`

	ValidateTime common.ValidateTime `json:"validateDate"`

	TotalPoint int64 `json:"totalPoint"`

	UpdateDate int64 `json:"updateDate"`
}
