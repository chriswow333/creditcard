package bankaccount

import "example.com/creditcard/models/reward"

type BankAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`

	Rewards []*reward.Reward `json:"rewards,omitempty"`
}
