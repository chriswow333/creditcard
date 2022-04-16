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
)

type Service interface {
	GetCustomizationsByCardID(ctx context.Context, cardID string) ([]*customization.Customization, error)
	GetCustomizationByID(ctx context.Context, ID string) (*customization.Customization, error)

	GetAllEcommerces(ctx context.Context) ([]*ecommerce.Ecommerce, error)
	GetEcommerce(ctx context.Context, ID string) (*ecommerce.Ecommerce, error)

	GetAllDeliverys(ctx context.Context) ([]*delivery.Delivery, error)
	GetDelivery(ctx context.Context, ID string) (*delivery.Delivery, error)

	GetAllMobilepays(ctx context.Context) ([]*mobilepay.Mobilepay, error)
	GetMobilepay(ctx context.Context, ID string) (*mobilepay.Mobilepay, error)

	GetAllOnlinegames(ctx context.Context) ([]*onlinegame.Onlinegame, error)
	GetOnlinegame(ctx context.Context, ID string) (*onlinegame.Onlinegame, error)

	GetAllStreamings(ctx context.Context) ([]*streaming.Streaming, error)
	GetStreaming(ctx context.Context, ID string) (*streaming.Streaming, error)

	GetAllSupermarkets(ctx context.Context) ([]*supermarket.Supermarket, error)
	GetSupermarket(ctx context.Context, ID string) (*supermarket.Supermarket, error)

	GetAllTimeIntervals(ctx context.Context) ([]*timeinterval.TimeInterval, error)
	GetTimeInterval(ctx context.Context, ID string) (*timeinterval.TimeInterval, error)
}
