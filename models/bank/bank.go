package bank

import (
	cardM "example.com/creditcard/models/card"
)

type Bank struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	UpdateDate int64  `json:"updateDate"`
	ImagePath  string `json:"imagePath"`
	LinkURL    string `json:"linkURL"`

	// BankAcconts []*bankaccount.BankAccount `json:"bankAccounts"`
	Cards []*cardM.Card `json:"cards"`
}

const DATE_FORMAT = "2006/01/02"
