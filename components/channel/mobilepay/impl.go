package mobilepay

import (
	"context"

	channelComp "example.com/creditcard/components/channel"
	channelM "example.com/creditcard/models/channel"
	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"
)

type impl struct {
	mobilepays []*channelM.Mobilepay
	channel    *channelM.Channel
}

func New(

	mobilepays []*channelM.Mobilepay,
	channel *channelM.Channel,

) channelComp.Component {
	return &impl{
		mobilepays: mobilepays,
		channel:    channel,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*channelM.ChannelEventResp, error) {

	channelEventResp := &channelM.ChannelEventResp{
		ChannelType:         channelM.MobilepayType,
		ChannelOperatorType: im.channel.ChannelOperatorType,
		ChannelMappingType:  im.channel.ChannelMappingType,
	}

	matches := []string{}
	misses := []string{}

	mobilepayMap := make(map[string]bool)

	for _, mo := range e.Mobilepays {
		mobilepayMap[mo] = true

	}

	for _, mo := range im.mobilepays {

		if _, ok := mobilepayMap[mo.ID]; ok {
			matches = append(matches, mo.ID)
		} else {
			misses = append(misses, mo.ID)
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

	logrus.Info("mobilepayComponent.Judge ", channelEventResp)

	return channelEventResp, nil
}
