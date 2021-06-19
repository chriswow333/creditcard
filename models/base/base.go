package base

type Base struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Desc []string `json:"desc"`
}

type AccountBase struct {
	Base

	BankAccount *BankAccount `json:"bankAccount"`
}

type MoneyBase struct {
	Base

	Currency string `json:"currency"`
	AtLeast  int64  `json:"atLeast"`
	AtMost   int64  `json:"atMost"`
}

type BankAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
