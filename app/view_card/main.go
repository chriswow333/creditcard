package main

import (
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"google.golang.org/grpc"

	postConn "example.com/creditcard/app/view_card/utils/conn/postgresql"
	"example.com/creditcard/base/psql"

	bankRoute "example.com/creditcard/app/view_card/routes/bank"
	cardRoute "example.com/creditcard/app/view_card/routes/card"
	rewardRoute "example.com/creditcard/app/view_card/routes/reward"
	bankService "example.com/creditcard/app/view_card/services/bank"
	cardService "example.com/creditcard/app/view_card/services/card"
	bankStore "example.com/creditcard/app/view_card/stores/bank"
	cardStore "example.com/creditcard/app/view_card/stores/card"
	featureStore "example.com/creditcard/app/view_card/stores/feature"

	rewardService "example.com/creditcard/app/view_card/services/reward"
	rewardStore "example.com/creditcard/app/view_card/stores/reward"

	taskStore "example.com/creditcard/app/view_card/stores/task"
)

func BuildContainer() *dig.Container {

	container := dig.New()

	container.Provide(psql.NewPsql) // new postgres
	container.Provide(postConn.New)

	container.Provide(bankService.New)
	container.Provide(bankStore.New)
	container.Provide(cardService.New)
	container.Provide(cardStore.New)
	container.Provide(featureStore.New)
	container.Provide(rewardService.New)
	container.Provide(rewardStore.New)
	container.Provide(taskStore.New)

	// service

	container.Provide(NewServer)

	return container
}

func NewServer(
	bankSrc bankService.Service,
	cardSrc cardService.Service,
	rewardSrc rewardService.Service,
) *grpc.Server {

	s := grpc.NewServer()

	bankRoute.NewRoute(s, bankSrc)
	cardRoute.NewRoute(s, cardSrc)
	rewardRoute.NewRoute(s, rewardSrc)

	return s
}

const (
	port = ":50051"
)

func main() {

	container := BuildContainer()

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := container.Invoke(func(s *grpc.Server) {

		logrus.Info("start serving grpc request", port)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}); err != nil {
		panic(err)
	}

}
