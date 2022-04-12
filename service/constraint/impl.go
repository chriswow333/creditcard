package constraint

import (
	"context"

	"example.com/creditcard/models/customization"
	"example.com/creditcard/models/delivery"
	"example.com/creditcard/models/ecommerce"
	"example.com/creditcard/models/mobilepay"
	"example.com/creditcard/models/onlinegame"
	"example.com/creditcard/models/streaming"
	"example.com/creditcard/models/supermarket"
	"example.com/creditcard/models/timeinterval"
	customizationStore "example.com/creditcard/stores/customization"
	deliveryStore "example.com/creditcard/stores/delivery"
	ecommerceStore "example.com/creditcard/stores/ecommerce"
	mobilepayStore "example.com/creditcard/stores/mobilepay"
	onlinegameStore "example.com/creditcard/stores/onlinegame"
	streamingStore "example.com/creditcard/stores/streaming"
	supermarketStore "example.com/creditcard/stores/supermarket"
	"github.com/sirupsen/logrus"
)

type impl struct {
	customizationStore customizationStore.Store
	ecommerceStore     ecommerceStore.Store
	deliveryStore      deliveryStore.Store
	mobilepayStore     mobilepayStore.Store
	onlinegameStore    onlinegameStore.Store
	streamingStore     streamingStore.Store
	supermarketStore   supermarketStore.Store
}

func New(
	customizationStore customizationStore.Store,
	ecommerceStore ecommerceStore.Store,
	deliveryStore deliveryStore.Store,
	mobilepayStore mobilepayStore.Store,
	onlinegameStore onlinegameStore.Store,
	streamingStore streamingStore.Store,
	supermarketStore supermarketStore.Store,
) Service {
	return &impl{
		customizationStore: customizationStore,
		ecommerceStore:     ecommerceStore,
		deliveryStore:      deliveryStore,
		mobilepayStore:     mobilepayStore,
		onlinegameStore:    onlinegameStore,
		streamingStore:     streamingStore,
		supermarketStore:   supermarketStore,
	}
}

func (im *impl) GetCustmozationsByCardID(ctx context.Context, cardID string) ([]*customization.Customization, error) {

	customizations, err := im.customizationStore.GetByCardID(ctx, cardID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return customizations, nil
}

func (im *impl) GetCustomizationByID(ctx context.Context, ID string) (*customization.Customization, error) {
	customization, err := im.customizationStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return customization, nil
}

func (im *impl) GetAllEcommerces(ctx context.Context) ([]*ecommerce.Ecommerce, error) {

	ecommerces, err := im.ecommerceStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return ecommerces, nil
}

func (im *impl) GetEcommerce(ctx context.Context, ID string) (*ecommerce.Ecommerce, error) {

	ecommerce, err := im.ecommerceStore.GetByID(ctx, ID)
	if err != nil {
		logrus.New().Error(err)
		return nil, err
	}
	return ecommerce, nil
}

func (im *impl) GetAllDeliverys(ctx context.Context) ([]*delivery.Delivery, error) {

	deliveries, err := im.deliveryStore.GetAll(ctx)

	if err != nil {
		logrus.New().Error(err)
		return nil, err
	}

	return deliveries, nil
}

func (im *impl) GetDelivery(ctx context.Context, ID string) (*delivery.Delivery, error) {
	delivery, err := im.deliveryStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return delivery, nil
}

func (im *impl) GetAllMobilepays(ctx context.Context) ([]*mobilepay.Mobilepay, error) {
	mobilepays, err := im.mobilepayStore.GetAll(ctx)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return mobilepays, nil
}

func (im *impl) GetMobilepay(ctx context.Context, ID string) (*mobilepay.Mobilepay, error) {

	mobilepay, err := im.mobilepayStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return mobilepay, nil
}

func (im *impl) GetAllOnlinegames(ctx context.Context) ([]*onlinegame.Onlinegame, error) {

	onlinegames, err := im.onlinegameStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return onlinegames, nil
}

func (im *impl) GetOnlinegame(ctx context.Context, ID string) (*onlinegame.Onlinegame, error) {
	onlinegame, err := im.onlinegameStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return onlinegame, nil
}

func (im *impl) GetAllStreamings(ctx context.Context) ([]*streaming.Streaming, error) {

	streamings, err := im.streamingStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return streamings, nil
}

func (im *impl) GetStreaming(ctx context.Context, ID string) (*streaming.Streaming, error) {

	streaming, err := im.streamingStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return streaming, nil
}

func (im *impl) GetAllSupermarkets(ctx context.Context) ([]*supermarket.Supermarket, error) {

	supermarkets, err := im.supermarketStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return supermarkets, nil
}

func (im *impl) GetSupermarket(ctx context.Context, ID string) (*supermarket.Supermarket, error) {
	supermarket, err := im.supermarketStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return supermarket, nil
}

func (im *impl) GetAllTimeIntervals(ctx context.Context) ([]*timeinterval.TimeInterval, error) {
	return nil, nil
}

func (im *impl) GetTimeInterval(ctx context.Context, ID string) (*timeinterval.TimeInterval, error) {

	return nil, nil
}
