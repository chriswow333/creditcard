package bank

import (
	"example.com/creditcard/models/bankaccount"
	"example.com/creditcard/models/card"
)

type Bank struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	UpdateDate int64  `json:"updateDate"`
	LinkURL    string `json:"linkURL"`

	BankAcconts []*bankaccount.BankAccount `json:"bankAccounts"`
	Cards       []*card.Card               `json:"cards"`
}
