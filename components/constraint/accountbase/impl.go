package accountbase

import (
	"context"

	"example.com/creditcard/components/constraint"
	baseM "example.com/creditcard/models/base"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	accountBase *baseM.AccountBase
}

func New(
	accountBase *baseM.AccountBase,
) constraint.Component {

	return &impl{
		accountBase: accountBase,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: true,
	}

	for _, account := range e.BankAccounts {
		if account.ID == im.accountBase.ID {

			resp.Pass = true
			return resp, nil
		}
	}
	return resp, nil
}
