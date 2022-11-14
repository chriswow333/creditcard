package transportation

import (
	"context"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	transportations []*channelM.Transportation
	channel         *channelM.Channel
}

func New(
	transportations []*channelM.Transportation,
	channel *channelM.Channel,
) channel.Component {
	return &impl{
		transportations: transportations,
		channel:         channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.TransportationType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}
	transportationMap := make(map[string]bool)

	for _, st := range e.Transportations {
		transportationMap[st] = true
	}

	for _, st := range im.channel.Transportations {
		if _, ok := transportationMap[st]; ok {
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
