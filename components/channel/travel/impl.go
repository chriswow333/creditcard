package travel

import (
	"context"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"
)

type impl struct {
	travels []*channelM.Travel
	channel *channelM.Channel
}

func New(
	travels []*channelM.Travel,
	channel *channelM.Channel,
) channel.Component {
	return &impl{
		travels: travels,
		channel: channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.TravelType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}
	travelMap := make(map[string]bool)

	for _, st := range e.Travels {
		travelMap[st] = true
	}

	for _, st := range im.channel.Travels {
		if _, ok := travelMap[st]; ok {
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

	logrus.Info("travelComp.Judge ", channelEventResp)
	return channelEventResp, nil

}
