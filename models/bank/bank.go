package bank

import (
	"example.com/creditcard/models/card"
)

type Bank struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	UpdateDate int64  `json:"updateDate"`
	ImagePath  string `json:"imagePath"`
	LinkURL    string `json:"linkURL"`

	// BankAcconts []*bankaccount.BankAccount `json:"bankAccounts"`
	Cards []*card.Card `json:"cards"`
}
