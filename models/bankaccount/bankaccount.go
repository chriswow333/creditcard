package bankaccount

import "example.com/creditcard/models/reward"

type BankAccount struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`

	Rewards []*reward.Reward `json:"rewards,omitempty"`
}
