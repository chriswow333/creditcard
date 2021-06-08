package bankaccount

import (
	"example.com/creditcard/models/privilage"
)

type BankAccount struct {
	ID   string
	Name string
	Desc string

	Privilages []*privilage.Privilage
}
