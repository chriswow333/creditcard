package onlinegame

import (
	"context"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	onlinegames []*channelM.Onlinegame
	channel     *channelM.Channel
}

func New(
	onlinegames []*channelM.Onlinegame,
	channel *channelM.Channel,

) channel.Component {

	return &impl{
		onlinegames: onlinegames,
		channel:     channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.OnlinegameType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}
	onlinegameMap := make(map[string]bool)

	for _, on := range e.Onlinegames {
		onlinegameMap[on] = true
	}

	for _, on := range im.onlinegames {
		if _, ok := onlinegameMap[on.ID]; ok {
			matches = append(matches, on.ID)
		} else {
			misses = append(misses, on.ID)
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
