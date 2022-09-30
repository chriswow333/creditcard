package appstore

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
		ChannelType:         channelM.AppStoreType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}

	appStoreMap := make(map[string]bool)

	for _, ec := range e.AppStores {
		appStoreMap[ec] = true
	}

	for _, ec := range im.channel.AppStores {
		if _, ok := appStoreMap[ec]; ok {
			matches = append(matches, ec)
		} else {
			misses = append(misses, ec)
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
