package amusement

import (
	"context"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	amusements []*channelM.Amusement
	channel    *channelM.Channel
}

func New(
	amusements []*channelM.Amusement,
	channel *channelM.Channel,
) channel.Component {
	return &impl{
		amusements: amusements,
		channel:    channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.AmusementType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}
	amusementMap := make(map[string]bool)

	for _, st := range e.Amusements {
		amusementMap[st] = true
	}

	for _, au := range im.amusements {
		if _, ok := amusementMap[au.ID]; ok {
			matches = append(matches, au.ID)
		} else {
			misses = append(misses, au.ID)
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
