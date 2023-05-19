package channel

import (
	"context"
	"runtime/debug"

	"example.com/creditcard/models/channel"
	"example.com/creditcard/models/task"
	amusementStore "example.com/creditcard/stores/amusement"
	appStoreStore "example.com/creditcard/stores/appstore"
	cinemaStore "example.com/creditcard/stores/cinema"
	convenienceStoreStore "example.com/creditcard/stores/conveniencestore"
	deliveryStore "example.com/creditcard/stores/delivery"
	ecommerceStore "example.com/creditcard/stores/ecommerce"
	foodStore "example.com/creditcard/stores/food"
	hotelStore "example.com/creditcard/stores/hotel"
	insuranceStore "example.com/creditcard/stores/insurance"
	mallStore "example.com/creditcard/stores/mall"
	mobilepayStore "example.com/creditcard/stores/mobilepay"
	onlinegameStore "example.com/creditcard/stores/onlinegame"
	publicUtilityStore "example.com/creditcard/stores/publicutility"
	sportStore "example.com/creditcard/stores/sport"
	streamingStore "example.com/creditcard/stores/streaming"
	supermarketStore "example.com/creditcard/stores/supermarket"
	taskStore "example.com/creditcard/stores/task"
	transportationStore "example.com/creditcard/stores/transportation"
	travelStore "example.com/creditcard/stores/travel"

	"github.com/sirupsen/logrus"

	uuid "github.com/nu7hatch/gouuid"
)

type impl struct {
	taskStore        taskStore.Store
	ecommerceStore   ecommerceStore.Store
	deliveryStore    deliveryStore.Store
	mobilepayStore   mobilepayStore.Store
	onlinegameStore  onlinegameStore.Store
	streamingStore   streamingStore.Store
	supermarketStore supermarketStore.Store

	foodStore           foodStore.Store
	transportationStore transportationStore.Store
	travelStore         travelStore.Store
	insuranceStore      insuranceStore.Store

	sportStore            sportStore.Store
	convenienceStoreStore convenienceStoreStore.Store
	mallStore             mallStore.Store

	appStoreStore appStoreStore.Store

	hotelStore     hotelStore.Store
	amusementStore amusementStore.Store

	cinemaStore cinemaStore.Store

	publicUtilityStore publicUtilityStore.Store
}

func New(
	taskStore taskStore.Store,
	ecommerceStore ecommerceStore.Store,
	deliveryStore deliveryStore.Store,
	mobilepayStore mobilepayStore.Store,
	onlinegameStore onlinegameStore.Store,
	streamingStore streamingStore.Store,
	supermarketStore supermarketStore.Store,
	foodStore foodStore.Store,
	transportationStore transportationStore.Store,
	travelStore travelStore.Store,
	insuranceStore insuranceStore.Store,
	mallStore mallStore.Store,
	sportStore sportStore.Store,
	convenienceStoreStore convenienceStoreStore.Store,
	appStoreStore appStoreStore.Store,
	hotelStore hotelStore.Store,
	amusementStore amusementStore.Store,
	cinemaStore cinemaStore.Store,
	publicUtilityStore publicUtilityStore.Store,
) Service {
	return &impl{
		taskStore:             taskStore,
		ecommerceStore:        ecommerceStore,
		deliveryStore:         deliveryStore,
		mobilepayStore:        mobilepayStore,
		onlinegameStore:       onlinegameStore,
		streamingStore:        streamingStore,
		supermarketStore:      supermarketStore,
		foodStore:             foodStore,
		transportationStore:   transportationStore,
		travelStore:           travelStore,
		insuranceStore:        insuranceStore,
		mallStore:             mallStore,
		sportStore:            sportStore,
		convenienceStoreStore: convenienceStoreStore,
		appStoreStore:         appStoreStore,
		hotelStore:            hotelStore,
		amusementStore:        amusementStore,
		cinemaStore:           cinemaStore,
		publicUtilityStore:    publicUtilityStore,
	}
}

func (im *impl) GetTaskByID(ctx context.Context, ID string) (*task.Task, error) {

	task, err := im.taskStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	return task, nil
}

func (im *impl) GetTasksByCardID(ctx context.Context, cardID string) ([]*task.Task, error) {

	tasks, err := im.taskStore.GetByCardID(ctx, cardID)

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return tasks, nil
}

func (im *impl) CreateTask(ctx context.Context, task *task.Task) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	task.ID = id.String()

	if err := im.taskStore.Create(ctx, task); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	return nil
}

func (im *impl) GetAllEcommerces(ctx context.Context, offset, limit int) ([]*channel.Ecommerce, error) {

	ecommerces, err := im.ecommerceStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return ecommerces, nil
}

func (im *impl) GetEcommerce(ctx context.Context, ID string) (*channel.Ecommerce, error) {

	ecommerce, err := im.ecommerceStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return ecommerce, nil
}

func (im *impl) GetAllDeliverys(ctx context.Context, offset, limit int) ([]*channel.Delivery, error) {

	deliveries, err := im.deliveryStore.GetAll(ctx, offset, limit)

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return deliveries, nil
}

func (im *impl) GetDelivery(ctx context.Context, ID string) (*channel.Delivery, error) {
	delivery, err := im.deliveryStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return delivery, nil
}

func (im *impl) GetAllMobilepays(ctx context.Context, offset, limit int) ([]*channel.Mobilepay, error) {
	mobilepays, err := im.mobilepayStore.GetAll(ctx, offset, limit)

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return mobilepays, nil
}

func (im *impl) GetMobilepay(ctx context.Context, ID string) (*channel.Mobilepay, error) {

	mobilepay, err := im.mobilepayStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return mobilepay, nil
}

func (im *impl) GetAllOnlinegames(ctx context.Context, offset, limit int) ([]*channel.Onlinegame, error) {

	onlinegames, err := im.onlinegameStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return onlinegames, nil
}

func (im *impl) GetOnlinegame(ctx context.Context, ID string) (*channel.Onlinegame, error) {
	onlinegame, err := im.onlinegameStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return onlinegame, nil
}

func (im *impl) GetAllStreamings(ctx context.Context, offset, limit int) ([]*channel.Streaming, error) {

	streamings, err := im.streamingStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return streamings, nil
}

func (im *impl) GetStreaming(ctx context.Context, ID string) (*channel.Streaming, error) {

	streaming, err := im.streamingStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return streaming, nil
}

func (im *impl) GetAllSupermarkets(ctx context.Context, offset, limit int) ([]*channel.Supermarket, error) {

	supermarkets, err := im.supermarketStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return supermarkets, nil
}

func (im *impl) GetSupermarket(ctx context.Context, ID string) (*channel.Supermarket, error) {
	supermarket, err := im.supermarketStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return supermarket, nil
}

func (im *impl) GetAllFoods(ctx context.Context, offset, limit int) ([]*channel.Food, error) {
	foods, err := im.foodStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return foods, nil
}

func (im *impl) GetFood(ctx context.Context, ID string) (*channel.Food, error) {
	food, err := im.foodStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return food, nil
}

func (im *impl) GetAllTransportations(ctx context.Context, offset, limit int) ([]*channel.Transportation, error) {
	transportations, err := im.transportationStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return transportations, nil
}

func (im *impl) GetTransportation(ctx context.Context, ID string) (*channel.Transportation, error) {
	transportation, err := im.transportationStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return transportation, nil
}

func (im *impl) GetAllTravels(ctx context.Context, offset, limit int) ([]*channel.Travel, error) {
	travels, err := im.travelStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return travels, nil
}

func (im *impl) GetTravel(ctx context.Context, ID string) (*channel.Travel, error) {
	travel, err := im.travelStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return travel, nil
}

func (im *impl) GetAllInsurances(ctx context.Context, offset, limit int) ([]*channel.Insurance, error) {

	insurance, err := im.insuranceStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return insurance, nil
}
func (im *impl) GetInsurance(ctx context.Context, ID string) (*channel.Insurance, error) {
	travel, err := im.insuranceStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return travel, nil

}

func (im *impl) GetAllSports(ctx context.Context, offset, limit int) ([]*channel.Sport, error) {

	sports, err := im.sportStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return sports, nil
}

func (im *impl) GetSport(ctx context.Context, ID string) (*channel.Sport, error) {

	sport, err := im.sportStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return sport, nil
}

func (im *impl) GetAllConvenienceStores(ctx context.Context, offset, limit int) ([]*channel.ConvenienceStore, error) {

	convenienceStores, err := im.convenienceStoreStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return convenienceStores, nil
}

func (im *impl) GetConvenienceStore(ctx context.Context, ID string) (*channel.ConvenienceStore, error) {
	convenienceStore, err := im.convenienceStoreStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return convenienceStore, nil
}

func (im *impl) GetAllMalls(ctx context.Context, offset, limit int) ([]*channel.Mall, error) {
	malls, err := im.mallStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return malls, nil
}

func (im *impl) GetMall(ctx context.Context, ID string) (*channel.Mall, error) {
	mall, err := im.mallStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return mall, nil
}

func (im *impl) GetAllAppstores(ctx context.Context, offset, limit int) ([]*channel.AppStore, error) {
	appstoreStores, err := im.appStoreStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return appstoreStores, nil
}

func (im *impl) GetAppstore(ctx context.Context, ID string) (*channel.AppStore, error) {
	appStore, err := im.appStoreStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return appStore, nil
}

func (im *impl) GetAllHotels(ctx context.Context, offset, limit int) ([]*channel.Hotel, error) {
	hotelStores, err := im.hotelStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return hotelStores, nil
}
func (im *impl) GetHotel(ctx context.Context, ID string) (*channel.Hotel, error) {

	hotelStore, err := im.hotelStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return hotelStore, nil
}

func (im *impl) GetAllAmusemnets(ctx context.Context, offset, limit int) ([]*channel.Amusement, error) {
	amusements, err := im.amusementStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return amusements, nil
}
func (im *impl) GetAmusement(ctx context.Context, ID string) (*channel.Amusement, error) {
	amusement, err := im.amusementStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n [msg]%s \n %s", err, string(debug.Stack()))
		return nil, err
	}

	return amusement, nil
}

func (im *impl) GetAllCinemas(ctx context.Context, offset, limit int) ([]*channel.Cinema, error) {
	cinemas, err := im.cinemaStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return cinemas, nil
}

func (im *impl) GetCinema(ctx context.Context, ID string) (*channel.Cinema, error) {
	cinema, err := im.cinemaStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return cinema, nil
}

func (im *impl) GetAllPublicUtilities(ctx context.Context, offset, limit int) ([]*channel.PublicUtility, error) {
	publicUtilities, err := im.publicUtilityStore.GetAll(ctx, offset, limit)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return publicUtilities, nil
}

func (im *impl) GetPublicUtility(ctx context.Context, ID string) (*channel.PublicUtility, error) {
	publicUtility, err := im.publicUtilityStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return publicUtility, nil
}

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.ChannelResp, error) {

	channelResps := []*channel.ChannelResp{}

	amusements, err := im.amusementStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(amusements) != 0 {
		amusementResp := &channel.ChannelResp{
			ChannelType: channel.AmusementType,
			Amusements:  amusements,
		}
		channelResps = append(channelResps, amusementResp)
	}

	appstores, err := im.appStoreStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(appstores) != 0 {
		appstoreResp := &channel.ChannelResp{
			ChannelType: channel.AppStoreType,
			Appstores:   appstores,
		}
		channelResps = append(channelResps, appstoreResp)
	}

	cinemas, err := im.cinemaStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(cinemas) != 0 {
		cinemaResp := &channel.ChannelResp{
			ChannelType: channel.CinemaType,
			Cinemas:     cinemas,
		}
		channelResps = append(channelResps, cinemaResp)
	}

	conveniencestores, err := im.convenienceStoreStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(conveniencestores) != 0 {
		cinemaResp := &channel.ChannelResp{
			ChannelType:       channel.ConvenienceStoreType,
			ConvenienceStores: conveniencestores,
		}
		channelResps = append(channelResps, cinemaResp)
	}

	deliveries, err := im.deliveryStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(deliveries) != 0 {
		deliveryResp := &channel.ChannelResp{
			ChannelType: channel.DeliveryType,
			Deliveries:  deliveries,
		}
		channelResps = append(channelResps, deliveryResp)
	}

	ecommerces, err := im.ecommerceStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(ecommerces) != 0 {
		ecommerceResp := &channel.ChannelResp{
			ChannelType: channel.EcommerceType,
			Ecommerces:  ecommerces,
		}
		channelResps = append(channelResps, ecommerceResp)
	}

	foods, err := im.foodStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(foods) != 0 {
		foodResp := &channel.ChannelResp{
			ChannelType: channel.FoodType,
			Foods:       foods,
		}
		channelResps = append(channelResps, foodResp)
	}

	hotels, err := im.foodStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(hotels) != 0 {
		hotelResp := &channel.ChannelResp{
			ChannelType: channel.HotelType,
			Foods:       foods,
		}
		channelResps = append(channelResps, hotelResp)
	}

	insurances, err := im.insuranceStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(insurances) != 0 {
		insuranceResp := &channel.ChannelResp{
			ChannelType: channel.InsuranceType,
			Insurances:  insurances,
		}
		channelResps = append(channelResps, insuranceResp)
	}

	malls, err := im.mallStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(malls) != 0 {
		mallResp := &channel.ChannelResp{
			ChannelType: channel.MallType,
			Malls:       malls,
		}
		channelResps = append(channelResps, mallResp)
	}

	mobilepays, err := im.mobilepayStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(mobilepays) != 0 {
		mobilepayResp := &channel.ChannelResp{
			ChannelType: channel.MobilepayType,
			Mobilepays:  mobilepays,
		}
		channelResps = append(channelResps, mobilepayResp)
	}

	onlinegames, err := im.onlinegameStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(onlinegames) != 0 {
		onlinegameResp := &channel.ChannelResp{
			ChannelType: channel.OnlinegameType,
			Onlinegames: onlinegames,
		}
		channelResps = append(channelResps, onlinegameResp)
	}

	publicutilities, err := im.publicUtilityStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(publicutilities) != 0 {
		publicutilityResp := &channel.ChannelResp{
			ChannelType:     channel.PublicUtilityType,
			PublicUtilities: publicutilities,
		}
		channelResps = append(channelResps, publicutilityResp)
	}

	sports, err := im.sportStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(sports) != 0 {
		sportResp := &channel.ChannelResp{
			ChannelType: channel.SportType,
			Sports:      sports,
		}
		channelResps = append(channelResps, sportResp)
	}

	streamings, err := im.streamingStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(streamings) != 0 {
		streamingResp := &channel.ChannelResp{
			ChannelType: channel.StreamingType,
			Streamings:  streamings,
		}
		channelResps = append(channelResps, streamingResp)
	}

	supermarkets, err := im.supermarketStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(supermarkets) != 0 {
		supermarketResp := &channel.ChannelResp{
			ChannelType:  channel.SupermarketType,
			Supermarkets: supermarkets,
		}
		channelResps = append(channelResps, supermarketResp)
	}

	transportations, err := im.transportationStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(transportations) != 0 {
		transportationResp := &channel.ChannelResp{
			ChannelType:     channel.TransportationType,
			Transportations: transportations,
		}
		channelResps = append(channelResps, transportationResp)
	}

	travels, err := im.travelStore.FindLike(ctx, names)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	if len(travels) != 0 {
		travelResp := &channel.ChannelResp{
			ChannelType: channel.TravelType,
			Travels:     travels,
		}
		channelResps = append(channelResps, travelResp)
	}

	return channelResps, nil
}
