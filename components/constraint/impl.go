package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	constraintComps []*Component
	constraintResp  *constraintM.ConstraintResp
}

func New(
	constraintComps []*Component,
	constraintResp *constraintM.ConstraintResp,
) Component {

	impl := &impl{
		constraintComps: constraintComps,
		constraintResp:  constraintResp,
	}
	return impl
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*constraintM.ConstraintEventResp, error) {

	constraintEventResp := &constraintM.ConstraintEventResp{}

	constraintEventResps := []*constraintM.ConstraintEventResp{}

	for _, co := range im.constraintComps {
		constraintEventResp, err := (*co).Judge(ctx, e)
		if err != nil {
			return nil, err
		}
		constraintEventResps = append(constraintEventResps, constraintEventResp)
	}

	switch im.constraintResp.ConstraintOperatorType {
	case constraintM.OR:
		for _, resp := range constraintEventResps {
			if resp.Pass {
				constraintEventResp.Pass = true
				break
			} else {
				constraintEventResp.Pass = false
			}
		}
	case constraintM.AND:
		for _, resp := range constraintEventResps {
			if resp.Pass {
				constraintEventResp.Pass = true
			} else {
				constraintEventResp.Pass = false
				break
			}
		}
	}

	constraintEventResp.ConstraintEventResps = constraintEventResps

	return constraintEventResp, nil
}
