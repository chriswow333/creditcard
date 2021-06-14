package event

import (
	"example.com/creditcard/models/bankaccount"
	"example.com/creditcard/models/card"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
)

type Event struct {
	ID string

	Amount        *Amount     `json:"amount"`
	EffictiveTime int64       `json:"effictiveTime"`
	Action        *ActionType `json:"actionType"`

	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerce,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarket,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegame,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streaming,omitempty"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilpays,omitempty"`
	Cards        []*card.Card               `json:"cards,omitempty"`
	BankAccounts []*bankaccount.BankAccount `json:"bankAccounts,omitempty"`
}

type CurrencyType int32

const (
	NTD CurrencyType = iota
)

type Amount struct {
	Total    int64         `json:"total"`
	Currency *CurrencyType `json:"currency"`
}

type ActionType int32

const (
	Shopping ActionType = iota
	Deposit
)

type Response struct {
	Pass bool `json:"pass,omitempty"`
}
