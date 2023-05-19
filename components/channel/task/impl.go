package customization

import (
	"context"
	"errors"
	"time"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	labelM "example.com/creditcard/models/label"
	"github.com/sirupsen/logrus"

	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/task"
	channelSvc "example.com/creditcard/service/channel"
)

type impl struct {
	channel        *channelM.Channel
	tasks          []*task.Task
	channelService channelSvc.Service
}

func New(
	channel *channelM.Channel,
	tasks []*task.Task,
	channelService channelSvc.Service,

) channel.Component {

	return &impl{
		channel:        channel,
		tasks:          tasks,
		channelService: channelService,
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
			break
		case task.WEEKDAY:
			pass = im.processWeekDayType(e, t, eventTaskMap)
			break
		case task.LABEL:
			pass = im.processLabel(ctx, e, t, eventTaskMap)
			break
		default:
			return nil, errors.New("not found task type")
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
		break
	case channelM.AND:
		if len(misses) > 0 || len(matches) == 0 {
			channelEventResp.Pass = false
		} else {
			channelEventResp.Pass = true
		}
		break
	default:
		return nil, errors.New("not found channel operatortype")
	}

	if im.channel.ChannelMappingType == channelM.MISMATCH {
		channelEventResp.Pass = !channelEventResp.Pass
	}

	return channelEventResp, nil
}

func (im *impl) processNoneType(t *task.Task, eventTaskMap map[string]bool) bool {
	var pass bool
	if _, ok := eventTaskMap[t.ID]; ok {
		pass = true
	} else if t.DefaultPass {
		pass = true
	} else {
		pass = false
	}
	logrus.Info("processNoneType status: ", pass, " ", t)
	return pass
}

func (im *impl) processWeekDayType(e *eventM.Event, t *task.Task, taskMap map[string]bool) bool {

	var pass bool = false
	if _, ok := taskMap[t.ID]; ok {
		pass = true
	} else {
		weekDay := time.Unix(e.EffectiveTime, 0).Weekday()

		weekdayLimit := t.TaskTypeModel.WeekDayLimit

		for _, d := range weekdayLimit.WeekDays {
			if d == int(weekDay) {
				pass = true
				break
			}
		}
	}

	logrus.Info("processWeekDayType status: ", pass, " ", t)
	return pass
}

// TODO 根據Label掃描所有的通路，並且確認有符合/無符合的通路掃出來做確認
// 只要有一個錯，就回錯

func (im *impl) processLabel(ctx context.Context, e *eventM.Event, t *task.Task, taskMap map[string]bool) bool {

	label := t.TaskTypeModel.Label
	var pass bool = false
	switch label.LabelType {
	case labelM.ALL:
		if label.Match {
			for _, id := range e.Tasks {
				for _, tid := range im.tasks {
					if id == tid.ID {
						pass = true
						break
					}
				}
				if pass {
					break
				}
			}
			if !pass {
				if len(e.Amusements) > 0 ||
					len(e.AppStores) > 0 ||
					len(e.Cinemas) > 0 ||
					len(e.Conveniencestores) > 0 ||
					len(e.Deliveries) > 0 ||
					len(e.Ecommerces) > 0 ||
					len(e.Foods) > 0 ||
					len(e.Hotels) > 0 ||
					len(e.Insurances) > 0 ||
					len(e.Malls) > 0 ||
					len(e.Onlinegames) > 0 ||
					len(e.Publicutilities) > 0 ||
					len(e.Sports) > 0 ||
					len(e.Streamings) > 0 ||
					len(e.Supermarkets) > 0 ||
					len(e.Transportations) > 0 ||
					len(e.Travels) > 0 {
					pass = true
				}
			}
		}

	case labelM.Channel:
		// TODO for loop 所有 channel
		pass = false

		break

	case labelM.OVERSEA:
		pass = im.processOversea(ctx, e, t)

	default:
		pass = false
	}

	logrus.Info("processLabel status: ", pass)
	return pass
}

func (im *impl) processOversea(ctx context.Context, e *eventM.Event, t *task.Task) bool {

	var pass bool = false
	for _, eTaskID := range e.Tasks {
		if eTaskID == t.ID {
			pass = true
			break
		}
	}
	return pass
}
