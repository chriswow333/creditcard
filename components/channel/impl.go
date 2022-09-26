package channel

import (
	"context"

	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"
)

type impl struct {
	channelComps []*Component
	channel      *channelM.Channel
}

func New(
	channelComps []*Component,
	channel *channelM.Channel,
) Component {

	impl := &impl{
		channelComps: channelComps,
		channel:      channel,
	}
	return impl
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResps := []*channelM.ChannelEventResp{}

	for _, cc := range im.channelComps {
		channelEventResp, err := (*cc).Judge(ctx, e)
		if err != nil {
			logrus.New().Error(err)
			return nil, err
		}
		channelEventResps = append(channelEventResps, channelEventResp)
	}

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.InnerChannelType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	switch im.channel.ChannelOperatorType {
	case channelM.OR:
		for _, resp := range channelEventResps {
			if resp.Pass {
				channelEventResp.Pass = true
				break
			} else {
				channelEventResp.Pass = false
			}
		}
	case channelM.AND:
		for _, resp := range channelEventResps {
			if resp.Pass {
				channelEventResp.Pass = true
			} else {
				channelEventResp.Pass = false
				break
			}
		}
	}

	channelEventResp.Matches = []string{}
	channelEventResp.Misses = []string{}

	channelEventResp.ChannelEventResps = channelEventResps

	return channelEventResp, nil
}
