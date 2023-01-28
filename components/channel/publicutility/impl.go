package publicutility

import (
	"context"

	"example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	publicutilities []*channelM.PublicUtility
	channel         *channelM.Channel
}

func New(
	publicutilities []*channelM.PublicUtility,
	channel *channelM.Channel,

) channel.Component {

	return &impl{
		publicutilities: publicutilities,
		channel:         channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.PublicUtilityType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}
	publicutilityMap := make(map[string]bool)

	for _, on := range e.Publicutilities {
		publicutilityMap[on] = true
	}

	for _, p := range im.publicutilities {
		if _, ok := publicutilityMap[p.ID]; ok {
			matches = append(matches, p.ID)
		} else {
			misses = append(misses, p.ID)
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
