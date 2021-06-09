package bank

import (
	"example.com/creditcard/models/bankaccount"
	"example.com/creditcard/models/card"
)

type Bank struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	BankAcconts []*bankaccount.BankAccount `json:"bankAccounts,omitempty"`
	Cards       []*card.Card               `json:"cards,omitempty"`
}