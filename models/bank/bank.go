package bank

import (
	"time"

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

func TransferBankResp(bank *Bank) *BankResp {

	bankResp := &BankResp{
		ID:         bank.ID,
		Name:       bank.Name,
		UpdateDate: time.Unix(bank.UpdateDate, 0).Format(DATE_FORMAT),
		ImagePath:  bank.ImagePath,
		LinkURL:    bank.LinkURL,
	}

	cardResps := []*cardM.CardResp{}

	if len(bank.Cards) != 0 {
		for _, c := range bank.Cards {
			cardResp := cardM.TransferCardResp(c)
			cardResps = append(cardResps, cardResp)
		}
	}

	bankResp.CardResps = cardResps

	return bankResp
}
