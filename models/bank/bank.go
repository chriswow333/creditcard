package bank

import (
	"example.com/creditcard/models/bankaccount"
	"example.com/creditcard/models/card"
)

type Bank struct {
	ID   string
	Name string
	Desc string

	StartDate  int64
	EndDate    int64
	UpdateDate int64

	BankAcconts []*bankaccount.BankAccount
	Cards       []*card.Card
}
