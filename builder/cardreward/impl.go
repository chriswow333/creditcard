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

	amusementStore "example.com/creditcard/stores/amusement"
	appstoreStore "example.com/creditcard/stores/appstore"
	cinemaStore "example.com/creditcard/stores/cinema"
	conveniencestoreStore "example.com/creditcard/stores/conveniencestore"
	deliveryStore "example.com/creditcard/stores/delivery"
	ecommerceStore "example.com/creditcard/stores/ecommerce"
	foodStore "example.com/creditcard/stores/food"
	hotelStore "example.com/creditcard/stores/hotel"
	insuraceStore "example.com/creditcard/stores/insurance"
	mallStore "example.com/creditcard/stores/mall"
	mobilepayStore "example.com/creditcard/stores/mobilepay"
	onlinegameStore "example.com/creditcard/stores/onlinegame"
	sportStore "example.com/creditcard/stores/sport"
	streamingStore "example.com/creditcard/stores/streaming"
	supermarketStore "example.com/creditcard/stores/supermarket"
	transportationStore "example.com/creditcard/stores/transportation"
	travelStore "example.com/creditcard/stores/travel"
)

type impl struct {
	*dig.In

	channelService channel.Service
	bankService    bank.Service

	// feedbackDescStore     feedbackDescStore.Store
	mobilepayStore        mobilepayStore.Store
	ecommerceStore        ecommerceStore.Store
	amusementStore        amusementStore.Store
	supermarketStore      supermarketStore.Store
	onlinegameStore       onlinegameStore.Store
	streamingStore        streamingStore.Store
	foodStore             foodStore.Store
	transportationStore   transportationStore.Store
	deliveryStore         deliveryStore.Store
	travelStore           travelStore.Store
	insuraceStore         insuraceStore.Store
	mallStore             mallStore.Store
	conveniencestoreStore conveniencestoreStore.Store
	sportStore            sportStore.Store
	appstoreStore         appstoreStore.Store
	hotelStore            hotelStore.Store
	cinemaStore           cinemaStore.Store
}

func New(
	channelService channel.Service,
	bankService bank.Service,

	// feedbackDescStore feedbackDescStore.Store,
	mobilepayStore mobilepayStore.Store,
	ecommerceStore ecommerceStore.Store,
	amusementStore amusementStore.Store,
	supermarketStore supermarketStore.Store,
	onlinegameStore onlinegameStore.Store,
	streamingStore streamingStore.Store,
	foodStore foodStore.Store,
	transportationStore transportationStore.Store,
	deliveryStore deliveryStore.Store,
	travelStore travelStore.Store,
	insuraceStore insuraceStore.Store,
	mallStore mallStore.Store,
	conveniencestoreStore conveniencestoreStore.Store,
	sportStore sportStore.Store,
	appstoreStore appstoreStore.Store,
	hotelStore hotelStore.Store,
	cinemaStore cinemaStore.Store,

) Builder {
	return &impl{
		channelService: channelService,
		bankService:    bankService,
		// feedbackDescStore:     feedbackDescStore,
		mobilepayStore:        mobilepayStore,
		ecommerceStore:        ecommerceStore,
		amusementStore:        amusementStore,
		supermarketStore:      supermarketStore,
		onlinegameStore:       onlinegameStore,
		streamingStore:        streamingStore,
		foodStore:             foodStore,
		transportationStore:   transportationStore,
		deliveryStore:         deliveryStore,
		travelStore:           travelStore,
		insuraceStore:         insuraceStore,
		mallStore:             mallStore,
		conveniencestoreStore: conveniencestoreStore,
		sportStore:            sportStore,
		appstoreStore:         appstoreStore,
		hotelStore:            hotelStore,
		cinemaStore:           cinemaStore,
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

		channelComponent = taskComp.New(channel, tasks, im.channelService)

		break
	case channelM.MobilepayType:
		mobilepayIDs := channel.Mobilepays
		mobilepays := []*channelM.Mobilepay{}

		for _, id := range mobilepayIDs {
			mobilepay, err := im.mobilepayStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"mobilepay get error": err,
					},
				).Error(err)
				return nil, err
			}
			mobilepays = append(mobilepays, mobilepay)
		}

		channelComponent = mobilepayComp.New(mobilepays, channel)
		break
	case channelM.EcommerceType:

		ecommerceIDs := channel.Ecommerces
		ecommerces := []*channelM.Ecommerce{}

		for _, id := range ecommerceIDs {
			ecommerce, err := im.ecommerceStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"mobilepay get error": err,
					},
				).Error(err)
				return nil, err
			}
			ecommerces = append(ecommerces, ecommerce)
		}

		channelComponent = ecommerceComp.New(ecommerces, channel)
		break
	case channelM.SupermarketType:

		supermarketIDs := channel.Supermarkets
		supermarkets := []*channelM.Supermarket{}

		for _, id := range supermarketIDs {
			supermarket, err := im.supermarketStore.GetByID(ctx, id)

			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"supermarket get error": err,
					},
				).Error(err)
				return nil, err
			}
			supermarkets = append(supermarkets, supermarket)
		}

		channelComponent = supermarketComp.New(supermarkets, channel)
		break
	case channelM.OnlinegameType:

		onlinegameIDs := channel.Onlinegames
		onlinegames := []*channelM.Onlinegame{}

		for _, id := range onlinegameIDs {
			onlinegame, err := im.onlinegameStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"onlinegamge get error": err,
					},
				).Error(err)
				return nil, err
			}
			onlinegames = append(onlinegames, onlinegame)
		}

		channelComponent = onlinegameComp.New(onlinegames, channel)
		break

	case channelM.StreamingType:

		streamingIDs := channel.Streamings
		streamings := []*channelM.Streaming{}

		for _, id := range streamingIDs {
			steraming, err := im.streamingStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"streaming get error": err,
					},
				).Error(err)
				return nil, err
			}
			streamings = append(streamings, steraming)
		}

		channelComponent = streamingComp.New(streamings, channel)
		break
	case channelM.FoodType:

		foodIDs := channel.Foods
		foods := []*channelM.Food{}

		for _, id := range foodIDs {

			food, err := im.foodStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"food get error": err,
					},
				).Error(err)
				return nil, err
			}

			foods = append(foods, food)
		}

		channelComponent = foodComp.New(foods, channel)

		break

	case channelM.TransportationType:

		transportationIDs := channel.Transportations
		transportations := []*channelM.Transportation{}

		for _, id := range transportationIDs {
			transportation, err := im.transportationStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"transportation get error": err,
					},
				).Error(err)
				return nil, err
			}
			transportations = append(transportations, transportation)
		}

		channelComponent = transportationComp.New(transportations, channel)

		break
	case channelM.DeliveryType:

		deliveryIDs := channel.Deliveries
		deliveries := []*channelM.Delivery{}

		for _, id := range deliveryIDs {
			delivery, err := im.deliveryStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"delivery get error": err,
					},
				).Error(err)
				return nil, err
			}
			deliveries = append(deliveries, delivery)
		}

		channelComponent = deliveryComp.New(deliveries, channel)
		break

	case channelM.TravelType:

		travelIDs := channel.Travels
		travels := []*channelM.Travel{}

		for _, id := range travelIDs {
			travel, err := im.travelStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"travel get error": err,
					},
				).Error(err)
				return nil, err
			}
			travels = append(travels, travel)
		}

		channelComponent = travelComp.New(travels, channel)
		break

	case channelM.InsuranceType:

		insuranceIDs := channel.Insurances
		insurances := []*channelM.Insurance{}

		for _, id := range insuranceIDs {
			insurance, err := im.insuraceStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"insurance get error": err,
					},
				).Error(err)
				return nil, err
			}
			insurances = append(insurances, insurance)
		}

		channelComponent = insuranceComp.New(insurances, channel)
		break
	case channelM.MallType:

		mallIDs := channel.Malls
		malls := []*channelM.Mall{}

		for _, id := range mallIDs {
			mall, err := im.mallStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"mall get error": err,
					},
				).Error(err)
				return nil, err
			}
			malls = append(malls, mall)
		}

		channelComponent = mallComp.New(malls, channel)
		break
	case channelM.ConvenienceStoreType:
		conveniencestoreIDs := channel.Conveniencestores
		conveniencestores := []*channelM.ConvenienceStore{}

		for _, id := range conveniencestoreIDs {
			conveniencestore, err := im.conveniencestoreStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"conveniencestore get error": err,
					},
				).Error(err)
				return nil, err
			}
			conveniencestores = append(conveniencestores, conveniencestore)
		}

		channelComponent = convenienceStoreComp.New(conveniencestores, channel)
		break
	case channelM.SportType:

		sportIDs := channel.Sports
		sports := []*channelM.Sport{}

		for _, id := range sportIDs {
			sport, err := im.sportStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"sport get error": err,
					},
				).Error(err)
				return nil, err
			}
			sports = append(sports, sport)
		}

		channelComponent = sportComp.New(sports, channel)
		break

	case channelM.AppStoreType:

		appstoreIDs := channel.AppStores
		appstores := []*channelM.AppStore{}

		for _, id := range appstoreIDs {
			appstore, err := im.appstoreStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"appstore get error": err,
					},
				).Error(err)
				return nil, err
			}
			appstores = append(appstores, appstore)
		}
		channelComponent = appStoreComp.New(appstores, channel)
		break

	case channelM.HotelType:

		hotelIDs := channel.Hotels
		hotels := []*channelM.Hotel{}

		for _, id := range hotelIDs {
			hotel, err := im.hotelStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"hotel get error": err,
					},
				).Error(err)
				return nil, err
			}
			hotels = append(hotels, hotel)
		}

		channelComponent = hotelComp.New(hotels, channel)
		break
	case channelM.AmusementType:

		amusementIDs := channel.Amusements
		amusements := []*channelM.Amusement{}

		for _, id := range amusementIDs {
			amusement, err := im.amusementStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"amusement get error": err,
					},
				).Error(err)
				return nil, err
			}
			amusements = append(amusements, amusement)
		}

		channelComponent = amusementComp.New(amusements, channel)
		break
	case channelM.CinemaType:

		cinemaIDs := channel.Cinemas
		cinemas := []*channelM.Cinema{}

		for _, id := range cinemaIDs {
			cinema, err := im.cinemaStore.GetByID(ctx, id)
			if err != nil {
				logrus.WithFields(
					logrus.Fields{
						"cinema get error": err,
					},
				).Error(err)
				return nil, err
			}
			cinemas = append(cinemas, cinema)
		}

		channelComponent = cinemaComp.New(cinemas, channel)

		break
	default:
		return nil, errors.New("failed in mapping contraint type")

	}

	return &channelComponent, nil
}

func (im *impl) getFeedbackComponent(ctx context.Context, rewardType rewardM.RewardType, feedback *feedbackM.Feedback) (*feedbackComp.Component, error) {

	switch rewardType {
	case rewardM.CASH:
		cashbackComponent := cashbackComp.New(feedback.Cashback)
		return &cashbackComponent, nil
	case rewardM.POINT:
		pointbackComponent := pointbackComp.New(feedback.Pointback)
		return &pointbackComponent, nil
	case rewardM.RED:
		redbackComponent := redbackComp.New(feedback.Redback)
		return &redbackComponent, nil
	default:
		logrus.WithFields(logrus.Fields{
			"cardReward.getFeedbackComponent": "reward not found",
		}).Error()
		return nil, errors.New("not found reward type")
	}
}
