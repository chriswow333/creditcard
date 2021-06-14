package reward

import (
	"context"

	constraintComp "example.com/creditcard/components/constraint"
	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"
)

type impl struct {
	constraintComps []*constraintComp.Component
}

func New(
	constraintComps []*constraintComp.Component,
) Component {
	return &impl{
		constraintComps: constraintComps,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) ([]*eventM.Response, error) {

	resps := []*eventM.Response{}

	for _, c := range im.constraintComps {
		resp, err := (*c).Judge(ctx, e)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		resps = append(resps, resp)
	}

	return resps, nil
}
