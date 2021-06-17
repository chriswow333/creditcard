package streaming

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/streaming"
	streamingM "example.com/creditcard/models/streaming"
)

type impl struct {
	streamings     []*streamingM.Streaming
	operator       constraintM.OperatorType
	constraintType constraintM.ConstraintType
	name           string
	descs          []string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {
	return &impl{
		streamings:     constraintPayload.Streamings,
		operator:       constraintPayload.Operator,
		constraintType: constraintPayload.ConstraintType,
		name:           constraintPayload.Name,
		descs:          constraintPayload.Descs,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {
	constraint := &eventM.Constraint{
		Name:           im.name,
		Descs:          im.descs,
		ConstraintType: im.constraintType,
	}

	matches := []string{}
	misses := []string{}
	streamingMap := make(map[string]*streaming.Streaming)

	for _, st := range e.Streamings {
		streamingMap[st.ID] = st
	}

	for _, ec := range im.streamings {
		if _, ok := streamingMap[ec.ID]; ok {
			matches = append(matches, ec.ID)
		} else {
			misses = append(misses, ec.ID)
		}
	}

	constraint.Matches = matches
	constraint.Misses = misses

	switch im.operator {
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
