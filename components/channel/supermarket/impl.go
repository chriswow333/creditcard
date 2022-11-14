package supermarket

import (
	"context"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	supermarkets []*channelM.Supermarket
	channel      *channelM.Channel
}

func New(
	supermarkets []*channelM.Supermarket,
	channel *channelM.Channel,
) channel.Component {
	return &impl{
		supermarkets: supermarkets,
		channel:      channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.SupermarketType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}
	supermarketMap := make(map[string]bool)

	for _, su := range e.Supermarkets {
		supermarketMap[su] = true
	}

	for _, su := range im.supermarkets {
		if _, ok := supermarketMap[su.ID]; ok {
			matches = append(matches, su.ID)
		} else {
			misses = append(misses, su.ID)
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
