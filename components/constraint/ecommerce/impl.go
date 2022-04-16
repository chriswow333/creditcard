package ecommerce

import (
	"context"

	"example.com/creditcard/components/constraint"

	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	constraintResp *constraintM.ConstraintResp
}

func New(
	constraintResp *constraintM.ConstraintResp,
) constraint.Component {

	return &impl{
		constraintResp: constraintResp,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*constraintM.ConstraintEventResp, error) {

	constraintEventResp := &constraintM.ConstraintEventResp{
		ConstraintType:         constraintM.EcommerceType,
		ConstraintOperatorType: im.constraintResp.ConstraintOperatorType,
		ConstraintMappingType:  im.constraintResp.ConstraintMappingType,
	}

	matches := []string{}
	misses := []string{}

	ecommerceMap := make(map[string]bool)

	for _, ec := range e.Ecommerces {
		ecommerceMap[ec] = true
	}

	for _, ec := range im.constraintResp.Ecommerces {
		if _, ok := ecommerceMap[ec.ID]; ok {
			matches = append(matches, ec.ID)
		} else {
			misses = append(misses, ec.ID)
		}
	}

	constraintEventResp.Matches = matches
	constraintEventResp.Misses = misses

	switch im.constraintResp.ConstraintOperatorType {
	case constraintM.OR:
		if len(matches) > 0 {
			constraintEventResp.Pass = true
		} else {
			constraintEventResp.Pass = false
		}
	case constraintM.AND:
		if len(misses) > 0 {
			constraintEventResp.Pass = false
		} else {
			constraintEventResp.Pass = true
		}
	}

	if im.constraintResp.ConstraintMappingType == constraintM.MISMATCH {
		constraintEventResp.Pass = !constraintEventResp.Pass
	}

	return constraintEventResp, nil
}
