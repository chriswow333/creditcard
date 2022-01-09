package cost

import (
	"example.com/creditcard/models/bonus"
	"example.com/creditcard/models/dollar"
)

type CostType int32

const (
	Dollar = iota
	Bonus
)

type Cost struct {
	CostType CostType `json:"costType"`

	Dollar *dollar.Dollar `json:"dollar"` // 現金回饋
	Bonus  *bonus.Bonus   `json:"bonus"`  // 點數回饋

}
