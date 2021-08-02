package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"google.golang.org/grpc"

	postConn "example.com/creditcard/app/view_card/utils/conn/postgresql"
	"example.com/creditcard/base/psql"

	bankRoute "example.com/creditcard/app/view_card/routes/bank"
	bankService "example.com/creditcard/app/view_card/services/bank"
	bankStore "example.com/creditcard/app/view_card/stores/bank"
)

func BuildContainer() *dig.Container {

	container := dig.New()

	container.Provide(psql.NewPsql) // new postgres
	container.Provide(postConn.New)

	container.Provide(bankService.New)
	container.Provide(bankStore.New)

	// service

	container.Provide(NewServer)

	return container
}

func NewServer(
	bankSrc bankService.Service,
) *grpc.Server {

	s := grpc.NewServer()
	bankRoute.NewRoute(s, bankSrc)
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

	container.Invoke(func(ss bankService.Service) {
		fmt.Println(ss)
	})

	if err := container.Invoke(func(s *grpc.Server) {

		logrus.Info("start serving grpc request", port)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}); err != nil {
		panic(err)
	}

}
