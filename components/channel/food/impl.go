package food

import (
	"context"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"
)

type impl struct {
	foods   []*channelM.Food
	channel *channelM.Channel
}

func New(
	foods []*channelM.Food,
	channel *channelM.Channel,
) channel.Component {
	return &impl{
		foods:   foods,
		channel: channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.FoodType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}
	foodMap := make(map[string]bool)

	for _, st := range e.Foods {
		foodMap[st] = true
	}

	for _, fo := range im.foods {

		if _, ok := foodMap[fo.ID]; ok {
			matches = append(matches, fo.ID)
		} else {
			misses = append(misses, fo.ID)
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

	logrus.Info("foodComponent.Judge ", channelEventResp)

	return channelEventResp, nil

}
