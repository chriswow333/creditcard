package streaming

import (
	"context"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	channel *channelM.Channel
}

func New(
	channel *channelM.Channel,
) channel.Component {
	return &impl{
		channel: channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.StreamingType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}
	streamingMap := make(map[string]bool)

	for _, st := range e.Streamings {
		streamingMap[st] = true
	}

	for _, st := range im.channel.Streamings {
		if _, ok := streamingMap[st]; ok {
			matches = append(matches, st)
		} else {
			misses = append(misses, st)
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
