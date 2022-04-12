package customization

import (
	"context"
	"fmt"
	"strconv"

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

	constraintEventResp := &constraintM.ConstraintEventResp{}

	matches := []string{}
	misses := []string{}

	customizationMap := make(map[string]bool)

	for _, c := range e.Customizations {
		customizationMap[c] = true
	}
	fmt.Println(customizationMap)

	for _, c := range im.constraintResp.Customizations {
		if _, ok := customizationMap[c.ID]; ok {
			matches = append(matches, c.ID)
		} else if c.DefaultPass {
			matches = append(matches, c.ID)
		} else {
			misses = append(misses, c.ID)
		}
	}

	fmt.Println("matches :")
	fmt.Println(matches)

	fmt.Println("misses: ")
	fmt.Println(misses)

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
	fmt.Println("-------")
	fmt.Println(len(matches))
	fmt.Println(im.constraintResp.Customizations)
	fmt.Println("customization pass : " + strconv.FormatBool(constraintEventResp.Pass))
	return constraintEventResp, nil
}
