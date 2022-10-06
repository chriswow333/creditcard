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

	GetAllEcommerces(ctx context.Context) ([]*channel.Ecommerce, error)
	GetEcommerce(ctx context.Context, ID string) (*channel.Ecommerce, error)

	GetAllDeliverys(ctx context.Context) ([]*channel.Delivery, error)
	GetDelivery(ctx context.Context, ID string) (*channel.Delivery, error)

	GetAllMobilepays(ctx context.Context) ([]*channel.Mobilepay, error)
	GetMobilepay(ctx context.Context, ID string) (*channel.Mobilepay, error)

	GetAllOnlinegames(ctx context.Context) ([]*channel.Onlinegame, error)
	GetOnlinegame(ctx context.Context, ID string) (*channel.Onlinegame, error)

	GetAllStreamings(ctx context.Context) ([]*channel.Streaming, error)
	GetStreaming(ctx context.Context, ID string) (*channel.Streaming, error)

	GetAllSupermarkets(ctx context.Context) ([]*channel.Supermarket, error)
	GetSupermarket(ctx context.Context, ID string) (*channel.Supermarket, error)

	GetAllFoods(ctx context.Context) ([]*channel.Food, error)
	GetFood(ctx context.Context, ID string) (*channel.Food, error)

	GetAllTransportations(ctx context.Context) ([]*channel.Transportation, error)
	GetTransportation(ctx context.Context, ID string) (*channel.Transportation, error)

	GetAllTravels(ctx context.Context) ([]*channel.Travel, error)
	GetTravel(ctx context.Context, ID string) (*channel.Travel, error)

	GetAllInsurances(ctx context.Context) ([]*channel.Insurance, error)
	GetInsurance(ctx context.Context, ID string) (*channel.Insurance, error)

	GetAllSports(ctx context.Context) ([]*channel.Sport, error)
	GetSport(ctx context.Context, ID string) (*channel.Sport, error)

	GetAllConvenienceStores(ctx context.Context) ([]*channel.ConvenienceStore, error)
	GetConvenienceStore(ctx context.Context, ID string) (*channel.ConvenienceStore, error)

	GetAllMalls(ctx context.Context) ([]*channel.Mall, error)
	GetMall(ctx context.Context, ID string) (*channel.Mall, error)

	GetAllAppstores(ctx context.Context) ([]*channel.AppStore, error)
	GetAppstore(ctx context.Context, ID string) (*channel.AppStore, error)

	GetAllHotels(ctx context.Context) ([]*channel.Hotel, error)
	GetHotel(ctx context.Context, ID string) (*channel.Hotel, error)

	GetAllAmusemnets(ctx context.Context) ([]*channel.Amusement, error)
	GetAmusement(ctx context.Context, ID string) (*channel.Amusement, error)

	GetAllCinemas(ctx context.Context) ([]*channel.Cinema, error)
	GetCinema(ctx context.Context, ID string) (*channel.Cinema, error)
}
