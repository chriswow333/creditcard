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
	bankService "example.com/creditcard/app/view_card/services/bank"
	bankStore "example.com/creditcard/app/view_card/stores/bank"
)

func BuildContainer() *dig.Container {

	container := dig.New()
	container.Provide(grpc.NewServer)

	container.Provide(psql.NewPsql, dig.Name("psql")) // new postgres
	container.Provide(postConn.New, dig.Name("connService"))

	container.Provide(bankService.New, dig.Name("bankService"))
	container.Provide(bankStore.New, dig.Name("bankStore"))

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

	if err := container.Invoke(func(s *grpc.Server) {
		logrus.Info("start serving grpc request")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}); err != nil {
		panic(err)
	}

}
