package channel

import (
	"context"

	"example.com/creditcard/models/channel"
	"example.com/creditcard/models/task"
)

type Service interface {
	GetTaskByID(ctx context.Context, ID string) (*task.Task, error)
	GetTasksByCardID(ctx context.Context, cardID string) ([]*task.Task, error)
	CreateTask(ctx context.Context, task *task.Task) error

	GetAllEcommerces(ctx context.Context, offset, limit int) ([]*channel.Ecommerce, error)
	GetEcommerce(ctx context.Context, ID string) (*channel.Ecommerce, error)

	GetAllDeliverys(ctx context.Context, offset, limit int) ([]*channel.Delivery, error)
	GetDelivery(ctx context.Context, ID string) (*channel.Delivery, error)

	GetAllMobilepays(ctx context.Context, offset, limit int) ([]*channel.Mobilepay, error)
	GetMobilepay(ctx context.Context, ID string) (*channel.Mobilepay, error)

	GetAllOnlinegames(ctx context.Context, offset, limit int) ([]*channel.Onlinegame, error)
	GetOnlinegame(ctx context.Context, ID string) (*channel.Onlinegame, error)

	GetAllStreamings(ctx context.Context, offset, limit int) ([]*channel.Streaming, error)
	GetStreaming(ctx context.Context, ID string) (*channel.Streaming, error)

	GetAllSupermarkets(ctx context.Context, offset, limit int) ([]*channel.Supermarket, error)
	GetSupermarket(ctx context.Context, ID string) (*channel.Supermarket, error)

	GetAllFoods(ctx context.Context, offset, limit int) ([]*channel.Food, error)
	GetFood(ctx context.Context, ID string) (*channel.Food, error)

	GetAllTransportations(ctx context.Context, offset, limit int) ([]*channel.Transportation, error)
	GetTransportation(ctx context.Context, ID string) (*channel.Transportation, error)

	GetAllTravels(ctx context.Context, offset, limit int) ([]*channel.Travel, error)
	GetTravel(ctx context.Context, ID string) (*channel.Travel, error)

	GetAllInsurances(ctx context.Context, offset, limit int) ([]*channel.Insurance, error)
	GetInsurance(ctx context.Context, ID string) (*channel.Insurance, error)

	GetAllSports(ctx context.Context, offset, limit int) ([]*channel.Sport, error)
	GetSport(ctx context.Context, ID string) (*channel.Sport, error)

	GetAllConvenienceStores(ctx context.Context, offset, limit int) ([]*channel.ConvenienceStore, error)
	GetConvenienceStore(ctx context.Context, ID string) (*channel.ConvenienceStore, error)

	GetAllMalls(ctx context.Context, offset, limit int) ([]*channel.Mall, error)
	GetMall(ctx context.Context, ID string) (*channel.Mall, error)

	GetAllAppstores(ctx context.Context, offset, limit int) ([]*channel.AppStore, error)
	GetAppstore(ctx context.Context, ID string) (*channel.AppStore, error)

	GetAllHotels(ctx context.Context, offset, limit int) ([]*channel.Hotel, error)
	GetHotel(ctx context.Context, ID string) (*channel.Hotel, error)

	GetAllAmusemnets(ctx context.Context, offset, limit int) ([]*channel.Amusement, error)
	GetAmusement(ctx context.Context, ID string) (*channel.Amusement, error)

	GetAllCinemas(ctx context.Context, offset, limit int) ([]*channel.Cinema, error)
	GetCinema(ctx context.Context, ID string) (*channel.Cinema, error)

	GetAllPublicUtilities(ctx context.Context, offset, limit int) ([]*channel.PublicUtility, error)
	GetPublicUtility(ctx context.Context, ID string) (*channel.PublicUtility, error)

	FindLike(ctx context.Context, names []string) ([]*channel.ChannelResp, error)
}
