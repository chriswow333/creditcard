package card

import (
	"context"
	"errors"
	"strings"
	"time"

	cardM "example.com/creditcard/models/card"
	channelM "example.com/creditcard/models/channel"
	rewardM "example.com/creditcard/models/reward"
	rewardChannelM "example.com/creditcard/models/reward_channel"
	taskM "example.com/creditcard/models/task"

	"example.com/creditcard/service/bank"
	"example.com/creditcard/service/channel"
	"example.com/creditcard/stores/card"
	"example.com/creditcard/stores/card_reward"

	// feedbackDescStore "example.com/creditcard/stores/feedback_desc"
	"example.com/creditcard/stores/reward"
	"example.com/creditcard/stores/reward_channel"

	"github.com/sirupsen/logrus"

	uuid "github.com/nu7hatch/gouuid"

	"go.uber.org/dig"
)

var (
	timeNow = time.Now
)

const DATE_FORMAT = "2006/01/02"

type impl struct {
	dig.In

	cardStore       card.Store
	rewardStore     reward.Store
	cardRewardStore card_reward.Store
	// feedbackDescStore feedbackDescStore.Store

	bankService          bank.Service
	rewardChannelService reward_channel.Store
	channelService       channel.Service
}

func New(
	cardStore card.Store,
	rewardStore reward.Store,
	cardRewardStore card_reward.Store,
	bankService bank.Service,
	rewardChannelService reward_channel.Store,
	channelService channel.Service,
	// feedbackDescStore feedbackDescStore.Store,
) Service {
	return &impl{
		cardStore:            cardStore,
		rewardStore:          rewardStore,
		cardRewardStore:      cardRewardStore,
		bankService:          bankService,
		rewardChannelService: rewardChannelService,
		channelService:       channelService,
		// feedbackDescStore:    feedbackDescStore,
	}
}

func (im *impl) Create(ctx context.Context, card *cardM.Card) error {

	card.UpdateDate = timeNow().Unix()

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)

		return err
	}
	card.ID = id.String()

	if err := im.cardStore.Create(ctx, card); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.Card, error) {

	card, err := im.cardStore.GetByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	cardRewards, err := im.cardRewardStore.GetByCardID(ctx, card.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for _, cr := range cardRewards {
		rewards, err := im.rewardStore.GetByCardRewardID(ctx, cr.ID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		cr.Rewards = rewards
	}

	card.CardRewards = cardRewards

	return card, nil
}

// func (im *impl) getFeedbackDesc(ctx context.Context, feedbackID string) (*feedback.FeedbackDesc, error) {

// 	feedbackDesc, err := im.feedbackDescStore.GetByID(ctx, feedbackID)

// 	if err != nil {
// 		logrus.WithFields(logrus.Fields{
// 			"": "",
// 		}).Error(err)
// 		return nil, err
// 	}

// 	return feedbackDesc, nil
// }

func (im *impl) transCardRewardResp(ctx context.Context, cardRewards []*cardM.CardReward) ([]*cardM.CardRewardResp, error) {

	cardRewardResps := []*cardM.CardRewardResp{}

	for _, cr := range cardRewards {

		// feedbackDesc, err := im.getFeedbackDesc(ctx, cr.FeedbackDescID)
		// if err != nil {
		// 	logrus.WithFields(logrus.Fields{
		// 		"": "",
		// 	}).Error(err)
		// 	return nil, err
		// }

		startDate := time.Unix(cr.StartDate, 0).Format(DATE_FORMAT)
		endDate := time.Unix(cr.EndDate, 0).Format(DATE_FORMAT)

		cardRewardResp := &cardM.CardRewardResp{
			ID:         cr.ID,
			RewardType: cr.RewardType,
			// CardRewardBonus: cr.CardRewardBonus,

			ConstraintPassLogics: cr.ConstraintPassLogics,
			Title:                cr.Title,
			Descs:                cr.Descs,
			StartDate:            startDate,
			EndDate:              endDate,
			CardRewardLimitTypes: cr.CardRewardLimitTypes,
			FeedbackBonus:        cr.FeedbackBonus,
			// FeedbackDesc:         feedbackDesc,
		}

		cardRewardID := cr.ID
		rewardChannels, err := im.rewardChannelService.GetByRewardID(ctx, cardRewardID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		tasks := []*taskM.Task{}
		taskMap := make(map[string]bool)

		mobilepays := []*channelM.Mobilepay{}
		mobilepayMap := make(map[string]bool)

		ecommerces := []*channelM.Ecommerce{}
		ecommerceMap := make(map[string]bool)

		supermarkets := []*channelM.Supermarket{}
		supermarketMap := make(map[string]bool)

		onlinegames := []*channelM.Onlinegame{}
		onlinegameMap := make(map[string]bool)

		streamings := []*channelM.Streaming{}
		streamingMap := make(map[string]bool)

		foods := []*channelM.Food{}
		foodMap := make(map[string]bool)

		transportations := []*channelM.Transportation{}
		transportationMap := make(map[string]bool)

		travels := []*channelM.Travel{}
		travelMap := make(map[string]bool)

		deliveries := []*channelM.Delivery{}
		deliveryMap := make(map[string]bool)

		insurances := []*channelM.Insurance{}
		insuranceMap := make(map[string]bool)

		malls := []*channelM.Mall{}
		mallMap := make(map[string]bool)

		sports := []*channelM.Sport{}
		sportMap := make(map[string]bool)

		convenienceStores := []*channelM.ConvenienceStore{}
		convenienceStoreMap := make(map[string]bool)

		appstores := []*channelM.AppStore{}
		appstoreMap := make(map[string]bool)

		hotels := []*channelM.Hotel{}
		hotelMap := make(map[string]bool)

		amusements := []*channelM.Amusement{}
		amusementMap := make(map[string]bool)

		cinemas := []*channelM.Cinema{}
		cinemaMap := make(map[string]bool)

		channelResps := []*channelM.ChannelResp{}

		for _, rc := range rewardChannels {

			switch rc.ChannelType {

			case int32(channelM.TaskType):

				task, err := im.channelService.GetTaskByID(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if task.TaskType == taskM.CHANNEL {

					for _, ch := range task.TaskTypeModel.ChannelLimit.Channels {
						switch ch {
						case int32(channelM.MobilepayType):

							allMobilepays, err := im.channelService.GetAllMobilepays(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, m := range allMobilepays {
								if _, ok := mobilepayMap[m.ID]; !ok {
									mobilepayMap[m.ID] = true
									mobilepays = append(mobilepays, m)
								}
							}

							break

						case int32(channelM.EcommerceType):
							allEcommerces, err := im.channelService.GetAllEcommerces(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, e := range allEcommerces {
								if _, ok := ecommerceMap[e.ID]; !ok {
									ecommerceMap[e.ID] = true
									ecommerces = append(ecommerces, e)
								}
							}

							break

						case int32(channelM.SupermarketType):
							allSupermarkets, err := im.channelService.GetAllSupermarkets(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, s := range allSupermarkets {
								if _, ok := supermarketMap[s.ID]; !ok {
									supermarketMap[s.ID] = true
									supermarkets = append(supermarkets, s)
								}
							}

							break

						case int32(channelM.OnlinegameType):
							allOnlinegames, err := im.channelService.GetAllOnlinegames(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, o := range allOnlinegames {
								if _, ok := onlinegameMap[o.ID]; !ok {
									onlinegameMap[o.ID] = true
									onlinegames = append(onlinegames, o)
								}
							}

							break

						case int32(channelM.StreamingType):

							allStreamings, err := im.channelService.GetAllStreamings(ctx)

							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, s := range allStreamings {
								if _, ok := streamingMap[s.ID]; !ok {
									streamingMap[s.ID] = true
									streamings = append(streamings, s)
								}
							}

							break

						case int32(channelM.FoodType):
							allFoods, err := im.channelService.GetAllFoods(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, f := range allFoods {
								if _, ok := foodMap[f.ID]; !ok {
									foodMap[f.ID] = true
									foods = append(foods, f)
								}
							}
							break

						case int32(channelM.TransportationType):
							allTransportations, err := im.channelService.GetAllTransportations(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, t := range allTransportations {
								if _, ok := transportationMap[t.ID]; !ok {
									transportationMap[t.ID] = true
									transportations = append(transportations, t)
								}
							}

							break

						case int32(channelM.TravelType):
							allTravels, err := im.channelService.GetAllTravels(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, t := range allTravels {
								if _, ok := travelMap[t.ID]; !ok {
									travelMap[t.ID] = true
									travels = append(travels, t)
								}
							}

							break

						case int32(channelM.DeliveryType):
							allDeliveries, err := im.channelService.GetAllDeliverys(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, d := range allDeliveries {
								if _, ok := deliveryMap[d.ID]; !ok {
									deliveryMap[d.ID] = true
									deliveries = append(deliveries, d)
								}
							}

							break

						case int32(channelM.InsuranceType):
							allInurances, err := im.channelService.GetAllInsurances(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, i := range allInurances {
								if _, ok := insuranceMap[i.ID]; !ok {
									insuranceMap[i.ID] = true
									insurances = append(insurances, i)
								}
							}

							break

						case int32(channelM.MallType):
							allMalls, err := im.channelService.GetAllMalls(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, m := range allMalls {
								if _, ok := mallMap[m.ID]; !ok {
									mallMap[m.ID] = true
									malls = append(malls, m)
								}
							}

							break

						case int32(channelM.SportType):
							allSports, err := im.channelService.GetAllSports(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, s := range allSports {
								if _, ok := sportMap[s.ID]; !ok {
									sportMap[s.ID] = true
									sports = append(sports, s)
								}
							}

							break

						case int32(channelM.ConvenienceStoreType):
							allConvenienceStores, err := im.channelService.GetAllConvenienceStores(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, c := range allConvenienceStores {
								if _, ok := convenienceStoreMap[c.ID]; !ok {
									convenienceStoreMap[c.ID] = true
									convenienceStores = append(convenienceStores, c)
								}
							}

							break

						case int32(channelM.AppStoreType):
							allAppStores, err := im.channelService.GetAllAppstores(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, a := range allAppStores {
								if _, ok := appstoreMap[a.ID]; !ok {
									appstoreMap[a.ID] = true
									appstores = append(appstores, a)
								}
							}

							break

						case int32(channelM.HotelType):
							allHotels, err := im.channelService.GetAllHotels(ctx)
							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, h := range allHotels {
								if _, ok := hotelMap[h.ID]; !ok {
									hotelMap[h.ID] = true
									hotels = append(hotels, h)
								}
							}

							break

						case int32(channelM.AmusementType):

							allAmusements, err := im.channelService.GetAllAmusemnets(ctx)

							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, a := range allAmusements {
								if _, ok := amusementMap[a.ID]; !ok {
									amusementMap[a.ID] = true
									amusements = append(amusements, a)
								}
							}

							break

						case int32(channelM.CinemaType):

							allCinemas, err := im.channelService.GetAllCinemas(ctx)

							if err != nil {
								logrus.WithFields(logrus.Fields{
									"": "",
								}).Error(err)
								break
							}

							for _, c := range allCinemas {
								if _, ok := cinemaMap[c.ID]; !ok {
									cinemaMap[c.ID] = true
									cinemas = append(cinemas, c)
								}
							}

							break
						}

					}
				}

				if _, ok := taskMap[task.ID]; !ok {
					taskMap[task.ID] = true
					tasks = append(tasks, task)
				}

				break

			case int32(channelM.MobilepayType):

				mobilepay, err := im.channelService.GetMobilepay(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := mobilepayMap[mobilepay.ID]; !ok {
					mobilepayMap[mobilepay.ID] = true
					mobilepays = append(mobilepays, mobilepay)
				}

				break

			case int32(channelM.EcommerceType):

				ecommerce := &channelM.Ecommerce{}

				ecommerce, err = im.channelService.GetEcommerce(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := ecommerceMap[ecommerce.ID]; !ok {
					ecommerceMap[ecommerce.ID] = true
					ecommerces = append(ecommerces, ecommerce)
				}

				break

			case int32(channelM.SupermarketType):

				supermarket, err := im.channelService.GetSupermarket(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := supermarketMap[supermarket.ID]; !ok {
					supermarketMap[supermarket.ID] = true
					supermarkets = append(supermarkets, supermarket)
				}

				break

			case int32(channelM.OnlinegameType):

				onlinegame, err := im.channelService.GetOnlinegame(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := onlinegameMap[onlinegame.ID]; !ok {
					onlinegameMap[onlinegame.ID] = true
					onlinegames = append(onlinegames, onlinegame)
				}

				break

			case int32(channelM.StreamingType):

				streaming, err := im.channelService.GetStreaming(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := streamingMap[streaming.ID]; !ok {
					streamingMap[streaming.ID] = true
					streamings = append(streamings, streaming)
				}

				break

			case int32(channelM.FoodType):

				food, err := im.channelService.GetFood(ctx, rc.ChannelID)

				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := foodMap[food.ID]; !ok {
					foodMap[food.ID] = true
					foods = append(foods, food)
				}

				break

			case int32(channelM.TransportationType):

				transportation, err := im.channelService.GetTransportation(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := transportationMap[transportation.ID]; !ok {
					transportationMap[transportation.ID] = true
					transportations = append(transportations, transportation)
				}

				break

			case int32(channelM.TravelType):

				travel, err := im.channelService.GetTravel(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := travelMap[travel.ID]; !ok {
					travelMap[travel.ID] = true
					travels = append(travels, travel)
				}

				break

			case int32(channelM.DeliveryType):

				delivery, err := im.channelService.GetDelivery(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := deliveryMap[delivery.ID]; !ok {
					deliveryMap[delivery.ID] = true
					deliveries = append(deliveries, delivery)
				}

				break

			case int32(channelM.InsuranceType):

				insurance, err := im.channelService.GetInsurance(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := insuranceMap[insurance.ID]; !ok {
					insuranceMap[insurance.ID] = true
					insurances = append(insurances, insurance)
				}

				break

			case int32(channelM.MallType):

				mall := &channelM.Mall{}

				mall, err = im.channelService.GetMall(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := mallMap[mall.ID]; !ok {
					mallMap[mall.ID] = true
					malls = append(malls, mall)
				}

				break

			case int32(channelM.SportType):

				sport, err := im.channelService.GetSport(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := sportMap[sport.ID]; !ok {
					sportMap[sport.ID] = true
					sports = append(sports, sport)
				}

				break

			case int32(channelM.ConvenienceStoreType):

				convenienceStore, err := im.channelService.GetConvenienceStore(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := convenienceStoreMap[convenienceStore.ID]; !ok {
					convenienceStoreMap[convenienceStore.ID] = true
					convenienceStores = append(convenienceStores, convenienceStore)
				}

				break

			case int32(channelM.AppStoreType):

				appstore, err := im.channelService.GetAppstore(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := appstoreMap[appstore.ID]; !ok {
					appstoreMap[appstore.ID] = true
					appstores = append(appstores, appstore)
				}

				break

			case int32(channelM.HotelType):

				hotel, err := im.channelService.GetHotel(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := hotelMap[hotel.ID]; !ok {
					hotelMap[hotel.ID] = true
					hotels = append(hotels, hotel)
				}

				break

			case int32(channelM.AmusementType):

				amusement, err := im.channelService.GetAmusement(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					break
				}

				if _, ok := amusementMap[amusement.ID]; !ok {
					amusementMap[amusement.ID] = true
					amusements = append(amusements, amusement)
				}

				break

			case int32(channelM.CinemaType):

				cinema, err := im.channelService.GetCinema(ctx, rc.ChannelID)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"card.transCardRewardResp": "",
					}).Error(err)
					break
				}

				if _, ok := cinemaMap[cinema.ID]; !ok {
					cinemaMap[cinema.ID] = true
					cinemas = append(cinemas, cinema)
				}

				break

			}
		}

		if len(tasks) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.TaskType,
				Tasks:       tasks,
			})
		}

		if len(mobilepays) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.MobilepayType,
				Mobilepays:  mobilepays,
			})
		}

		if len(ecommerces) > 0 {

			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.EcommerceType,
				Ecommerces:  ecommerces,
			})
		}

		if len(supermarkets) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType:  channelM.SupermarketType,
				Supermarkets: supermarkets,
			})
		}

		if len(onlinegames) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.OnlinegameType,
				Onlinegames: onlinegames,
			})
		}

		if len(streamings) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.StreamingType,
				Streamings:  streamings,
			})
		}

		if len(foods) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.FoodType,
				Foods:       foods,
			})
		}

		if len(transportations) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType:     channelM.TransportationType,
				Transportations: transportations,
			})
		}

		if len(travels) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.TravelType,
				Travels:     travels,
			})
		}

		if len(deliveries) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.DeliveryType,
				Deliveries:  deliveries,
			})
		}

		if len(insurances) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.InsuranceType,
				Insurances:  insurances,
			})
		}

		if len(malls) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.MallType,
				Malls:       malls,
			})
		}

		if len(sports) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.SportType,
				Sports:      sports,
			})
		}

		if len(convenienceStores) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType:       channelM.ConvenienceStoreType,
				ConvenienceStores: convenienceStores,
			})
		}

		if len(appstores) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.AppStoreType,
				Appstores:   appstores,
			})
		}

		if len(hotels) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.HotelType,
				Hotels:      hotels,
			})
		}

		if len(amusements) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.AmusementType,
				Amusements:  amusements,
			})
		}

		if len(cinemas) > 0 {
			channelResps = append(channelResps, &channelM.ChannelResp{
				ChannelType: channelM.CinemaType,
				Cinemas:     cinemas,
			})
		}

		cardRewardResp.ChannelResps = channelResps

		cardRewardResps = append(cardRewardResps, cardRewardResp)

	}

	return cardRewardResps, nil

}

func (im *impl) GetRespByID(ctx context.Context, ID string) (*cardM.CardResp, error) {

	card, err := im.cardStore.GetByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	updateDate := time.Unix(card.UpdateDate, 0).Format(DATE_FORMAT)

	bank, err := im.bankService.GetByID(ctx, card.BankID)
	cardResp := &cardM.CardResp{
		ID:    ID,
		Name:  card.Name,
		Descs: card.Descs,

		BankID:   card.BankID,
		BankName: bank.Name,

		UpdateDate: updateDate,

		ImagePath: card.ImagePath,
		LinkURL:   card.LinkURL,

		CardStatus:       card.CardStatus,
		OtherRewardResps: card.OtherRewards,
	}

	cardRewards, err := im.cardRewardStore.GetByCardID(ctx, card.ID)

	cardRewardResps, err := im.transCardRewardResp(ctx, cardRewards)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	cardResp.CardRewardResps = cardRewardResps
	return cardResp, nil

}

func (im *impl) UpdateByID(ctx context.Context, card *cardM.Card) error {
	if err := im.cardStore.UpdateByID(ctx, card); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*cardM.Card, error) {
	cards, err := im.cardStore.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return cards, nil
}

func (im *impl) GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error) {
	cards, err := im.cardStore.GetByBankID(ctx, bankID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return cards, nil
}

func (im *impl) CreateCardReward(ctx context.Context, cardReward *cardM.CardReward) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)

		return err
	}
	cardReward.ID = id.String()

	if err := im.cardRewardStore.Create(ctx, cardReward); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	for _, r := range cardReward.Rewards {
		r.CardRewardID = cardReward.ID

		id, err := uuid.NewV4()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": "",
			}).Error(err)
			return err
		}

		r.ID = id.String()

		for _, p := range r.Payloads {

			pid, err := uuid.NewV4()

			if err != nil {
				logrus.WithFields(logrus.Fields{
					"msg": "",
				}).Fatal(err)
				return err
			}
			p.ID = pid.String()
		}

		if err := im.rewardStore.Create(ctx, r); err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": "",
			}).Fatal(err)
			return err
		}

		if err := im.createRewardChannels(ctx, cardReward.CardID, cardReward.ID, r); err != nil {
			logrus.Error(err)
			return err
		}

	}

	return nil
}

func (im *impl) createRewardChannels(ctx context.Context, cardID, cardRewardID string, reward *rewardM.Reward) error {

	channelTypeMap := make(map[channelM.ChannelType]map[string]bool)

	for _, p := range reward.Payloads {
		if err := findAllChannelID(p.Channel, channelTypeMap); err != nil {
			return err
		}
	}

	for channelType, channelIDMap := range channelTypeMap {

		for channelID := range channelIDMap {
			id, err := uuid.NewV4()

			if err != nil {
				logrus.WithFields(logrus.Fields{
					"msg": "",
				}).Fatal(err)

				return err
			}

			rewardChannelM := &rewardChannelM.RewardChannel{
				ID:           id.String(),
				Order:        0,
				CardID:       cardID,
				CardRewardID: cardRewardID,
				ChannelID:    channelID,
				ChannelType:  int32(channelType),
			}

			if err := im.rewardChannelService.Create(ctx, rewardChannelM); err != nil {

				logrus.WithFields(logrus.Fields{
					"": "",
				}).Error(err)

				return err
			}

		}

	}

	return nil
}

func findAllChannelID(channel *channelM.Channel, channelTypeMap map[channelM.ChannelType]map[string]bool) error {

	if channel.ChannelMappingType != channelM.MATCH {
		return nil
	}

	switch channel.ChannelType {
	case channelM.InnerChannelType:
		for _, c := range channel.InnerChannels {
			findAllChannelID(c, channelTypeMap)
		}
		break
	case channelM.TaskType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, t := range channel.Tasks {
			channelTypeMap[channel.ChannelType][t] = true
		}

		break
	case channelM.MobilepayType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, m := range channel.Mobilepays {
			channelTypeMap[channel.ChannelType][m] = true
		}

		break
	case channelM.EcommerceType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, e := range channel.Ecommerces {
			channelTypeMap[channel.ChannelType][e] = true
		}

		break
	case channelM.SupermarketType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, s := range channel.Supermarkets {
			channelTypeMap[channel.ChannelType][s] = true
		}

		break
	case channelM.OnlinegameType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, o := range channel.Onlinegames {
			channelTypeMap[channel.ChannelType][o] = true
		}

		break
	case channelM.StreamingType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, s := range channel.Streamings {
			channelTypeMap[channel.ChannelType][s] = true
		}

		break
	case channelM.FoodType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, f := range channel.Foods {
			channelTypeMap[channel.ChannelType][f] = true
		}

		break
	case channelM.TransportationType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, t := range channel.Transportations {
			channelTypeMap[channel.ChannelType][t] = true
		}

		break
	case channelM.TravelType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, t := range channel.Travels {
			channelTypeMap[channel.ChannelType][t] = true
		}

		break
	case channelM.DeliveryType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, d := range channel.Deliveries {
			channelTypeMap[channel.ChannelType][d] = true
		}

		break
	case channelM.InsuranceType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, i := range channel.Insurances {
			channelTypeMap[channel.ChannelType][i] = true
		}

		break
	case channelM.MallType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, m := range channel.Malls {
			channelTypeMap[channel.ChannelType][m] = true
		}

		break
	case channelM.SportType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, s := range channel.Sports {
			channelTypeMap[channel.ChannelType][s] = true
		}

		break
	case channelM.ConvenienceStoreType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, c := range channel.Conveniencestores {
			channelTypeMap[channel.ChannelType][c] = true
		}

		break

	case channelM.AppStoreType:
		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, c := range channel.AppStores {
			channelTypeMap[channel.ChannelType][c] = true
		}

		break
	case channelM.HotelType:
		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, c := range channel.Hotels {
			channelTypeMap[channel.ChannelType][c] = true
		}

		break

	case channelM.AmusementType:
		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, c := range channel.Amusements {
			channelTypeMap[channel.ChannelType][c] = true
		}
		break

	case channelM.CinemaType:

		if _, ok := channelTypeMap[channel.ChannelType]; !ok {
			channelTypeMap[channel.ChannelType] = make(map[string]bool)
		}

		for _, c := range channel.Cinemas {
			channelTypeMap[channel.ChannelType][c] = true
		}

		break

	default:
		logrus.WithFields(logrus.Fields{
			"": "findAllChannelID",
		}).Error("")
		// return logrus.Error(err)
		return errors.New("no suitable channelType")
	}

	return nil
}

func (im *impl) EvaluateConstraintLogic(ctx context.Context, cardRewardID string, constraintIDs []string) (bool, string, error) {

	cardReward, err := im.cardRewardStore.GetByID(ctx, cardRewardID)
	if err != nil {
		return false, "internal error", err
	}

	constraintSet := make(map[string]bool)
	for _, constraintID := range constraintIDs {
		constraintSet[constraintID] = true
	}
	for _, logic := range cardReward.ConstraintPassLogics {
		pass, _, err := checkConstraintLogic(logic.Logic, constraintSet)
		if err != nil {
			return false, "internal error", err
		}

		if !pass {
			return false, logic.Message, nil
		}
	}

	return true, "", nil
}

/**

A, B, C are constraint ID
((A^B)C)

if event has no constraint ID, that means true

*/
func checkConstraintLogic(constraintPassLogic string, constraintIDs map[string]bool) (bool, bool, error) {

	pos := 0

	for pos = 0; pos < len(constraintPassLogic); pos++ {

		ch := constraintPassLogic[pos : pos+1]

		if ch == "(" {
			lastPos := strings.LastIndex(constraintPassLogic, ")")
			if lastPos == -1 {
				return false, false, errors.New("constraintPassLogic is illegal")
			}

			pass, exist, err := checkConstraintLogic(constraintPassLogic[1:lastPos], constraintIDs)
			if err != nil {
				return false, exist, err
			} else {
				return pass, exist, nil
			}

		} else if ch == "&" || ch == "|" || ch == "^" {
			constraintPassLogicPrev := constraintPassLogic[0:pos]
			constraintPassLogicLast := constraintPassLogic[pos+1:]
			passPrev, existPrev, err := checkConstraintLogic(constraintPassLogicPrev, constraintIDs)
			if err != nil {
				return false, false, err
			}
			passLast, existLast, err := checkConstraintLogic(constraintPassLogicLast, constraintIDs)
			if err != nil {
				return false, false, err
			}

			switch ch {
			case "&":
				return (passPrev && passLast) || (!existPrev && !existLast), existPrev, nil // if no one exist, return true
			case "|":
				return (passPrev || passLast) || (!existPrev && !existLast), existLast, nil // if no one exist, return true
			case "^":
				return (passPrev || passLast) && !(passPrev && passLast) || (!existPrev && !existLast), existLast, nil // if no one exist, return true
			}
		}
	}

	if _, ok := constraintIDs[constraintPassLogic]; ok {
		return true, true, nil
	} else {
		return false, false, nil
	}
}
