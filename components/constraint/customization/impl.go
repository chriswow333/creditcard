package customization

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	"example.com/creditcard/models/customization"
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
		ConstraintType:         constraintM.CustomizationType,
		ConstraintOperatorType: im.constraintResp.ConstraintOperatorType,
		ConstraintMappingType:  im.constraintResp.ConstraintMappingType,
	}

	matches := []string{}
	misses := []string{}

	customizationMap := make(map[string]bool)

	for _, c := range e.Customizations {
		customizationMap[c] = true
	}

	for _, c := range im.constraintResp.Customizations {
		switch c.CustomizationType {
		case customization.NONE:
			if im.processNoneType(c, customizationMap) {
				matches = append(matches, c.ID)
			} else {
				misses = append(misses, c.ID)
			}
		case customization.CASH:
			if im.processCashType(c, e) {
				matches = append(matches, c.ID)
			} else {
				misses = append(misses, c.ID)
			}
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

func (im *impl) processNoneType(c *customization.Customization, customizationMap map[string]bool) bool {
	if _, ok := customizationMap[c.ID]; ok {
		return true
	} else if c.DefaultPass {
		return true
	} else {
		return false
	}
}

func (im *impl) processCashType(c *customization.Customization, e *eventM.Event) bool {
	cashLimit := c.CustomizationTypeModel.CashLimit
	min := cashLimit.Min
	max := cashLimit.Max

	if min != 0 && int64(e.Cash) < min {
		return false
	}

	if max != 0 && int64(e.Cash) > max {
		return false
	}

	return true

}
