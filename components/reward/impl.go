package reward

import (
	"context"

	constraintComp "example.com/creditcard/components/constraint"
	eventM "example.com/creditcard/models/event"
	rewardM "example.com/creditcard/models/reward"
	"github.com/sirupsen/logrus"
)

type impl struct {
	reward          *rewardM.Reward
	constraintComps []*constraintComp.Component
}

func New(
	reward *rewardM.Reward,
	constraintComps []*constraintComp.Component,
) Component {
	return &impl{
		reward:          reward,
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
