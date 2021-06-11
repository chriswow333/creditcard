package constraint

import (
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
)

type OperatorType int32

const (
	AndOperator OperatorType = iota
	OrOperator
)

type Constraint struct {
	ID          string `json:"id"`
	PrivilageID string `json:"privilageID"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`

	StartDate  int64 `json:"startDate"`
	EndDate    int64 `json:"endDate"`
	UpdateDate int64 `json:"updateDate"`

	Limit *Limit `json:"limit,omitempty"`

	ConstraintBody *ConstraintBody `json:"constraintBody,omitempty"`
}

type Limit struct {
	Max int `json:"max"`
	Min int `json:"min"`
}

type ConstraintBody struct {
	ConstraintPayloads []*ConstraintPayload `json:"constraintPayloads,omitempty"`
}

type ConstraintPayload struct {
	Operator OperatorType `json:"operator"`

	ConstraintPayloads []*ConstraintPayload `json:"constraintPayloads,omitempty"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilepays,omitempty"`
	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`

	TimeBases    []*TimeBase    `json:"timeBases,omitempty"`
	AccountBases []*AccountBase `json:"accountBases,omitempty"`
	MoneyBases   []*MoneyBase   `json:"moneyBases,omitempty"`
}

type Base struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
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

	BankAccount string `json:"bankAccount"`
}

type MoneyBase struct {
	Base

	Currency string `json:"currency"`
	AtLeast  int64  `json:"atLeast"`
	AtMost   int64  `json:"atMost"`
}
