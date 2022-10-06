package cardreward

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	cardComp "example.com/creditcard/components/card"
	channelComp "example.com/creditcard/components/channel"

	amusementComp "example.com/creditcard/components/channel/amusement"
	appStoreComp "example.com/creditcard/components/channel/appstore"
	cinemaComp "example.com/creditcard/components/channel/cinema"
	convenienceStoreComp "example.com/creditcard/components/channel/conveniencestore"
	deliveryComp "example.com/creditcard/components/channel/delivery"
	ecommerceComp "example.com/creditcard/components/channel/ecommerce"
	foodComp "example.com/creditcard/components/channel/food"
	hotelComp "example.com/creditcard/components/channel/hotel"
	insuranceComp "example.com/creditcard/components/channel/insurance"
	mallComp "example.com/creditcard/components/channel/mall"
	mobilepayComp "example.com/creditcard/components/channel/mobilepay"
	onlinegameComp "example.com/creditcard/components/channel/onlinegame"
	sportComp "example.com/creditcard/components/channel/sport"
	streamingComp "example.com/creditcard/components/channel/streaming"
	supermarketComp "example.com/creditcard/components/channel/supermarket"
	taskComp "example.com/creditcard/components/channel/task"
	transportationComp "example.com/creditcard/components/channel/transportation"
	travelComp "example.com/creditcard/components/channel/travel"

	feedbackComp "example.com/creditcard/components/feedback"
	cashbackComp "example.com/creditcard/components/feedback/cashback"
	pointbackComp "example.com/creditcard/components/feedback/pointback"
	redbackComp "example.com/creditcard/components/feedback/redback"

	payloadComp "example.com/creditcard/components/payload"

	rewardComp "example.com/creditcard/components/reward"
	"example.com/creditcard/service/bank"
	"example.com/creditcard/service/channel"

	cardM "example.com/creditcard/models/card"
	channelM "example.com/creditcard/models/channel"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
	"example.com/creditcard/models/task"
)

type impl struct {
	*dig.In

	channelService channel.Service
	bankService    bank.Service
}

func New(
	channelService channel.Service,
	bankService bank.Service,
) Builder {
	return &impl{
		channelService: channelService,
		bankService:    bankService,
	}

}

func (im *impl) BuildCardComponent(ctx context.Context, card *cardM.Card) (cardComp.Component, error) {

	rewardMapper := make(map[string][]*rewardComp.Component)

	cardRewardOperatorMapper := make(map[rewardM.RewardType]cardM.CardRewardOperator)

	for _, cr := range card.CardRewards {

		rewardType := cr.RewardType

		cardRewardOperatorMapper[rewardType] = cr.CardRewardOperator

		for _, r := range cr.Rewards {

			payloadComponents := []*payloadComp.Component{}

			for _, p := range r.Payloads {
				channelComponent, err := im.getChannelComponent(ctx, p.Channel)
				if err != nil {
					logrus.New().Error(err)
					return nil, err
				}

				feedbackComponent, err := im.getFeedbackComponent(ctx, cr.RewardType, p.Feedback)
				if err != nil {
					logrus.New().Error(err)
					return nil, err
				}

				payloadComponent := payloadComp.New(p, channelComponent, feedbackComponent)
				payloadComponents = append(payloadComponents, &payloadComponent)
			}

			rewardComponent := rewardComp.New(cr.RewardType, r, payloadComponents)

			if rewardCmp, ok := rewardMapper[cr.ID]; ok {
				rewardMapper[cr.ID] = append(rewardCmp, &rewardComponent)
			} else {
				rewardComponents := []*rewardComp.Component{}
				rewardComponents = append(rewardComponents, &rewardComponent)
				rewardMapper[cr.ID] = rewardComponents
			}

		}

	}

	cardComponent := cardComp.New(card, rewardMapper, cardRewardOperatorMapper, im.bankService)

	return cardComponent, nil
}

func (im *impl) getChannelComponent(ctx context.Context, channel *channelM.Channel) (*channelComp.Component, error) {

	channelType := channel.ChannelType

	var channelComponent channelComp.Component

	switch channelType {
	case channelM.InnerChannelType:

		channelComponents := []*channelComp.Component{}

		for _, c := range channel.InnerChannels {
			channelComponent, err := im.getChannelComponent(ctx, c)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"": err,
					},
				).Error(err)
				return nil, err
			}
			channelComponents = append(channelComponents, channelComponent)
		}

		channelComponent = channelComp.New(channelComponents, channel)

		break
	case channelM.TaskType:

		//
		tasks := []*task.Task{}
		for _, t := range channel.Tasks {
			task, err := im.channelService.GetTaskByID(ctx, t)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"": err,
					},
				).Error(err)
				return nil, err
			}

			tasks = append(tasks, task)
		}

		channelComponent = taskComp.New(channel, tasks)

		break
	case channelM.MobilepayType:
		channelComponent = mobilepayComp.New(channel)
		break
	case channelM.EcommerceType:
		channelComponent = ecommerceComp.New(channel)
		break
	case channelM.SupermarketType:
		channelComponent = supermarketComp.New(channel)
		break
	case channelM.OnlinegameType:
		channelComponent = onlinegameComp.New(channel)
		break
	case channelM.StreamingType:
		channelComponent = streamingComp.New(channel)
		break
	case channelM.FoodType:
		channelComponent = foodComp.New(channel)
		break
	case channelM.TransportationType:
		channelComponent = transportationComp.New(channel)
		break
	case channelM.DeliveryType:
		channelComponent = deliveryComp.New(channel)
		break
	case channelM.TravelType:
		channelComponent = travelComp.New(channel)
		break
	case channelM.InsuranceType:
		channelComponent = insuranceComp.New(channel)
		break
	case channelM.MallType:
		channelComponent = mallComp.New(channel)
		break
	case channelM.ConvenienceStoreType:
		channelComponent = convenienceStoreComp.New(channel)
		break
	case channelM.SportType:
		channelComponent = sportComp.New(channel)
		break
	case channelM.AppStoreType:
		channelComponent = appStoreComp.New(channel)
		break
	case channelM.HotelType:
		channelComponent = hotelComp.New(channel)
		break
	case channelM.AmusementType:
		channelComponent = amusementComp.New(channel)
		break
	case channelM.CinemaType:
		channelComponent = cinemaComp.New(channel)
		break
	default:
		return nil, errors.New("failed in mapping contraint type")

	}

	return &channelComponent, nil
}

func (im *impl) getFeedbackComponent(ctx context.Context, rewardType rewardM.RewardType, feedback *feedbackM.Feedback) (*feedbackComp.Component, error) {

	switch rewardType {
	case rewardM.CASH_TWD:
		cashbackComponent := cashbackComp.New(feedback.Cashback)
		return &cashbackComponent, nil
	case rewardM.LINE_POINT:
		pointbackComponent := pointbackComp.New(feedback.Pointback)
		return &pointbackComponent, nil
	case rewardM.KUO_BROTHERS_POINT:
		pointbackComponent := pointbackComp.New(feedback.Pointback)
		return &pointbackComponent, nil
	case rewardM.WOWPRIME_POINT:
		pointbackComponent := pointbackComp.New(feedback.Pointback)
		return &pointbackComponent, nil
	case rewardM.OPEN_POINT:
		pointbackComponent := pointbackComp.New(feedback.Pointback)
		return &pointbackComponent, nil
	case rewardM.YIDA_POINT:
		pointbackComponent := pointbackComp.New(feedback.Pointback)
		return &pointbackComponent, nil

	case rewardM.RED_POINT:
		redbackComponent := redbackComp.New(feedback.Redback)
		return &redbackComponent, nil
	default:
		logrus.WithFields(logrus.Fields{
			"cardReward.getFeedbackComponent": "reward not found",
		}).Error()
		return nil, errors.New("not found reward type")
	}
}
