package constraint

import (
	"context"
	"fmt"

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

	for _, c := range constraintComps {

		fmt.Println("===== ", (*c))
	}
	impl := &impl{
		constraints:        constraintComps,
		constraintOperator: constraint.ConstraintOperator,
		constraintType:     constraint.ConstraintType,
	}
	return impl
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*constraintM.ConstraintResp, error) {

	constraint := &constraintM.ConstraintResp{
		ConstraintType: im.constraintType,
	}

	fmt.Println("bug")
	fmt.Println(im.constraintType)

	constraintResps := []*constraintM.ConstraintResp{}

	for _, co := range im.constraints {
		fmt.Println("--- ", (*co))
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
