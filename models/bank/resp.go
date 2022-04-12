package bank

import cardM "example.com/creditcard/models/card"

type BankResp struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	UpdateDate string `json:"updateDate"`
	ImagePath  string `json:"imagePath"`
	LinkURL    string `json:"linkURL"`

	// BankAcconts []*bankaccount.BankAccount `json:"bankAccounts"`
	CardResps []*cardM.CardResp `json:"cardResps"`
}
