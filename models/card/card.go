package card

import (
	"example.com/creditcard/models/privilage"
)

type Card struct {
	ID        string `json:"id"`
	BankID    string `json:"bankID"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	StartDate int64  `json:"startDate"`
	EndDate   int64  `json:"endDate"`

	UpdateDate int64 `json:"updateDate"`

	Privilages []*privilage.Privilage `json:"privilages,omitempty"`
}
