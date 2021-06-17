package base

type Base struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Desc []string `json:"desc"`
}

type TimeBase struct {
	Base

	DayFrom     string `json:"day"`
	WeekDayFrom string `json:"weekDay"`
	HourFrom    string `json:"hour"`
	MinuteFrom  string `json:"minute"`

	DayTo     string `json:"dayTo"`
	WeekDayTo string `json:"weekDayTo"`
	HourTo    string `json:"hourTo"`
	MinuteTo  string `json:"minuteTo"`
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
