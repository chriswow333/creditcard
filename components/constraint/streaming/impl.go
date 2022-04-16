package streaming

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
		ConstraintType:         constraintM.StreamingType,
		ConstraintOperatorType: im.constraintResp.ConstraintOperatorType,
		ConstraintMappingType:  im.constraintResp.ConstraintMappingType,
	}

	matches := []string{}
	misses := []string{}
	streamingMap := make(map[string]bool)

	for _, st := range e.Streamings {
		streamingMap[st] = true
	}

	for _, st := range im.constraintResp.Streamings {
		if _, ok := streamingMap[st.ID]; ok {
			matches = append(matches, st.ID)
		} else {
			misses = append(misses, st.ID)
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
