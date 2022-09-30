package channel

import (
	"context"

	"example.com/creditcard/models/channel"
	"example.com/creditcard/models/task"
	appStoreStore "example.com/creditcard/stores/appstore"
	convenienceStoreStore "example.com/creditcard/stores/conveniencestore"
	deliveryStore "example.com/creditcard/stores/delivery"
	ecommerceStore "example.com/creditcard/stores/ecommerce"
	foodStore "example.com/creditcard/stores/food"
	insuranceStore "example.com/creditcard/stores/insurance"
	mallStore "example.com/creditcard/stores/mall"
	mobilepayStore "example.com/creditcard/stores/mobilepay"
	onlinegameStore "example.com/creditcard/stores/onlinegame"
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
	}
}

func (im *impl) GetTaskByID(ctx context.Context, ID string) (*task.Task, error) {

	task, err := im.taskStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return task, nil
}

func (im *impl) GetTasksByCardID(ctx context.Context, cardID string) ([]*task.Task, error) {

	tasks, err := im.taskStore.GetByCardID(ctx, cardID)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return tasks, nil
}

func (im *impl) CreateTask(ctx context.Context, task *task.Task) error {

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)

		return err
	}

	task.ID = id.String()

	if err := im.taskStore.Create(ctx, task); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAllEcommerces(ctx context.Context) ([]*channel.Ecommerce, error) {

	ecommerces, err := im.ecommerceStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return ecommerces, nil
}

func (im *impl) GetEcommerce(ctx context.Context, ID string) (*channel.Ecommerce, error) {

	ecommerce, err := im.ecommerceStore.GetByID(ctx, ID)
	if err != nil {
		logrus.New().Error(err)
		return nil, err
	}
	return ecommerce, nil
}

func (im *impl) GetAllDeliverys(ctx context.Context) ([]*channel.Delivery, error) {

	deliveries, err := im.deliveryStore.GetAll(ctx)

	if err != nil {
		logrus.New().Error(err)
		return nil, err
	}

	return deliveries, nil
}

func (im *impl) GetDelivery(ctx context.Context, ID string) (*channel.Delivery, error) {
	delivery, err := im.deliveryStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return delivery, nil
}

func (im *impl) GetAllMobilepays(ctx context.Context) ([]*channel.Mobilepay, error) {
	mobilepays, err := im.mobilepayStore.GetAll(ctx)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return mobilepays, nil
}

func (im *impl) GetMobilepay(ctx context.Context, ID string) (*channel.Mobilepay, error) {

	mobilepay, err := im.mobilepayStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return mobilepay, nil
}

func (im *impl) GetAllOnlinegames(ctx context.Context) ([]*channel.Onlinegame, error) {

	onlinegames, err := im.onlinegameStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return onlinegames, nil
}

func (im *impl) GetOnlinegame(ctx context.Context, ID string) (*channel.Onlinegame, error) {
	onlinegame, err := im.onlinegameStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return onlinegame, nil
}

func (im *impl) GetAllStreamings(ctx context.Context) ([]*channel.Streaming, error) {

	streamings, err := im.streamingStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return streamings, nil
}

func (im *impl) GetStreaming(ctx context.Context, ID string) (*channel.Streaming, error) {

	streaming, err := im.streamingStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return streaming, nil
}

func (im *impl) GetAllSupermarkets(ctx context.Context) ([]*channel.Supermarket, error) {

	supermarkets, err := im.supermarketStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return supermarkets, nil
}

func (im *impl) GetSupermarket(ctx context.Context, ID string) (*channel.Supermarket, error) {
	supermarket, err := im.supermarketStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return supermarket, nil
}

func (im *impl) GetAllFoods(ctx context.Context) ([]*channel.Food, error) {
	foods, err := im.foodStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return foods, nil
}

func (im *impl) GetFood(ctx context.Context, ID string) (*channel.Food, error) {
	food, err := im.foodStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return food, nil
}

func (im *impl) GetAllTransportations(ctx context.Context) ([]*channel.Transportation, error) {
	transportations, err := im.transportationStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return transportations, nil
}

func (im *impl) GetTransportation(ctx context.Context, ID string) (*channel.Transportation, error) {
	transportation, err := im.transportationStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return transportation, nil
}

func (im *impl) GetAllTravels(ctx context.Context) ([]*channel.Travel, error) {
	travels, err := im.travelStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return travels, nil
}

func (im *impl) GetTravel(ctx context.Context, ID string) (*channel.Travel, error) {
	travel, err := im.travelStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return travel, nil
}

func (im *impl) GetAllInsurances(ctx context.Context) ([]*channel.Insurance, error) {

	insurance, err := im.insuranceStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return insurance, nil
}
func (im *impl) GetInsurance(ctx context.Context, ID string) (*channel.Insurance, error) {
	travel, err := im.insuranceStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return travel, nil

}

func (im *impl) GetAllSports(ctx context.Context) ([]*channel.Sport, error) {

	sports, err := im.sportStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return sports, nil
}

func (im *impl) GetSport(ctx context.Context, ID string) (*channel.Sport, error) {

	sport, err := im.sportStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return sport, nil
}

func (im *impl) GetAllConvenienceStores(ctx context.Context) ([]*channel.ConvenienceStore, error) {

	convenienceStores, err := im.convenienceStoreStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return convenienceStores, nil
}

func (im *impl) GetConvenienceStore(ctx context.Context, ID string) (*channel.ConvenienceStore, error) {
	convenienceStore, err := im.convenienceStoreStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return convenienceStore, nil
}

func (im *impl) GetAllMalls(ctx context.Context) ([]*channel.Mall, error) {
	malls, err := im.mallStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return malls, nil
}

func (im *impl) GetMall(ctx context.Context, ID string) (*channel.Mall, error) {
	mall, err := im.mallStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return mall, nil
}

func (im *impl) GetAllAppstores(ctx context.Context) ([]*channel.AppStore, error) {
	appstoreStores, err := im.appStoreStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return appstoreStores, nil
}

func (im *impl) GetAppstore(ctx context.Context, ID string) (*channel.AppStore, error) {
	appStore, err := im.appStoreStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return appStore, nil
}
