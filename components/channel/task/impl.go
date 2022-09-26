package customization

import (
	"context"
	"time"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/task"
)

type impl struct {
	channel *channelM.Channel
	tasks   []*task.Task
}

func New(
	channel *channelM.Channel,
	tasks []*task.Task,
) channel.Component {

	return &impl{
		channel: channel,
		tasks:   tasks,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.TaskType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}

	eventTaskMap := make(map[string]bool)

	for _, t := range e.Tasks {
		eventTaskMap[t] = true
	}

	for _, t := range im.tasks {

		taskType := t.TaskType

		pass := false

		switch taskType {

		case task.NONE:
			pass = im.processNoneType(t, eventTaskMap)
		case task.WEEKDAY:
			pass = im.processWeekDayType(e, t, eventTaskMap)

		}

		if pass {
			matches = append(matches, t.ID)
		} else {
			misses = append(misses, t.ID)
		}
	}

	channelEventResp.Matches = matches
	channelEventResp.Misses = misses

	switch im.channel.ChannelOperatorType {
	case channelM.OR:
		if len(matches) > 0 {
			channelEventResp.Pass = true
		} else {
			channelEventResp.Pass = false
		}
	case channelM.AND:
		if len(misses) > 0 || len(matches) == 0 {
			channelEventResp.Pass = false
		} else {
			channelEventResp.Pass = true
		}
	}

	if im.channel.ChannelMappingType == channelM.MISMATCH {
		channelEventResp.Pass = !channelEventResp.Pass
	}

	return channelEventResp, nil
}

func (im *impl) processNoneType(t *task.Task, eventTaskMap map[string]bool) bool {
	if _, ok := eventTaskMap[t.ID]; ok {
		return true
	} else if t.DefaultPass {
		return true
	} else {
		return false
	}

}

func (im *impl) processWeekDayType(e *eventM.Event, t *task.Task, taskMap map[string]bool) bool {

	if _, ok := taskMap[t.ID]; ok {
		return true
	} else {
		weekDay := time.Unix(e.EffectiveTime, 0).Weekday()

		weekdayLimit := t.TaskTypeModel.WeekDayLimit

		for _, d := range weekdayLimit.WeekDays {
			if d == int(weekDay) {
				return true
			}
		}
		return false
	}

}
