package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	constraints        []*Component
	constraintOperator constraintM.OperatorType
	constraintType     constraintM.ConstraintType
}

func New(
	constraintComps []*Component,
	constraint *constraintM.Constraint,
) Component {
	impl := &impl{
		constraints:        constraintComps,
		constraintOperator: constraint.ConstraintOperator,
		constraintType:     constraint.ConstraintType,
	}
	return impl
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error) {

	constraint := &eventM.ConstraintResp{
		ConstraintType: im.constraintType,
	}

	constraintResps := []*eventM.ConstraintResp{}

	for _, co := range im.constraints {
		constraintResp, err := (*co).Judge(ctx, e)
		if err != nil {
			return nil, err
		}
		constraintResps = append(constraintResps, constraintResp)
	}

	switch im.constraintOperator {
	case constraintM.OrOperator:
		for _, resp := range constraintResps {
			if resp.Pass {
				constraint.Pass = true
				break
			} else {
				constraint.Pass = false
			}
		}
	case constraintM.AndOperator:
		for _, resp := range constraintResps {
			if resp.Pass {
				constraint.Pass = true
			} else {
				constraint.Pass = false
				break
			}
		}
	}

	constraint.Constraints = constraintResps

	return constraint, nil
}
