package event

import (
	"example.com/creditcard/models/action"
	"example.com/creditcard/models/bankaccount"
	"example.com/creditcard/models/bonus"
	"example.com/creditcard/models/card"
	"example.com/creditcard/models/constraint"
	"example.com/creditcard/models/cost"
	"example.com/creditcard/models/customization"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
)

type Event struct {
	ID string

	Cost          *cost.CurrentCost  `json:"cost"`
	EffictiveTime int64              `json:"effictiveTime"`
	ActionType    *action.ActionType `json:"actionType"`

	Customizations []*customization.Customization `json:"customizations,omitempty"`

	Ecommerces   []*ecommerce.Ecommerce     `json:"ecommerces,omitempty"`
	Supermarkets []*supermarket.Supermarket `json:"supermarkets,omitempty"`
	Onlinegames  []*onlinegame.Onlinegame   `json:"onlinegames,omitempty"`
	Streamings   []*streaming.Streaming     `json:"streamings,omitempty"`

	Mobilepays   []*mobilepay.Mobilepay     `json:"mobilpays,omitempty"`
	Cards        []*card.Card               `json:"cards,omitempty"`
	BankAccounts []*bankaccount.BankAccount `json:"bankAccounts,omitempty"`
}

type Response struct {
	EventID   string   `json:"eventID"`
	Name      string   `json:"name"`
	Descs     []string `json:"descs"`
	StartDate int64    `json:"startDate"`
	EndDate   int64    `json:"endDate"`

	Rewards    []*Reward    `json:"rewards"`
	TotalBonus *bonus.Bonus `json:"totalBonus"`
	CountBonus *bonus.Bonus `json:"countBonus"`
	LinkURL    string       `json:"linkURL"`
}

type Reward struct {
	Pass        bool         `json:"pass"`
	Bonus       *bonus.Bonus `json:"bonus"`
	Name        string
	Descs       []string                `json:"descs"`
	Operator    constraint.OperatorType `json:"operator,omitempty"`
	Constraints []*Constraint           `json:"constraints,omitempty"`
}

type Constraint struct {
	Pass           bool                      `json:"pass"`
	ConstraintType constraint.ConstraintType `json:"constraintType"`
	Name           string                    `json:"name"`
	Descs          []string                  `json:"descs"`
	Matches        []string                  `json:"matches"`
	Misses         []string                  `json:"misses"`
	Constraints    []*Constraint             `json:"constraints"`
}
