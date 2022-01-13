package bank

import (
	"example.com/creditcard/models/bankaccount"
	"example.com/creditcard/models/card"
)

type Bank struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`

	StartDate  int64 `json:"startDate,omitempty"`
	EndDate    int64 `json:"endDate,omitempty"`
	UpdateDate int64 `json:"updateDate,omitempty"`

	BankAcconts []*bankaccount.BankAccount `json:"bankAccounts,omitempty"`
	Cards       []*card.Card               `json:"cards,omitempty"`
}
