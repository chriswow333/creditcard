package streaming

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	streamingM "example.com/creditcard/models/streaming"
)

type impl struct {
	streamings         []*streamingM.Streaming
	constraintOperator constraintM.OperatorType
	constraintType     constraintM.ConstraintType
	name               string
}

func New(
	constraint *constraintM.Constraint,
) constraint.Component {
	return &impl{
		streamings:         constraint.Streamings,
		constraintOperator: constraint.ConstraintOperator,
		constraintType:     constraint.ConstraintType,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error) {
	constraint := &eventM.ConstraintResp{
		ConstraintType: im.constraintType,
	}

	matches := []string{}
	misses := []string{}
	streamingMap := make(map[string]*streamingM.Streaming)

	for _, st := range e.Streamings {
		streamingMap[st.ID] = st
	}

	for _, st := range im.streamings {
		if _, ok := streamingMap[st.ID]; ok {
			matches = append(matches, st.ID)
		} else {
			misses = append(misses, st.ID)
		}
	}

	constraint.Matches = matches
	constraint.Misses = misses

	switch im.constraintOperator {
	case constraintM.OrOperator:
		if len(matches) > 0 {
			constraint.Pass = true
		} else {
			constraint.Pass = false
		}
	case constraintM.AndOperator:
		if len(misses) > 0 {
			constraint.Pass = false
		} else {
			constraint.Pass = true
		}
	}

	return constraint, nil

}
