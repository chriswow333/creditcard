package main

import (
	"example.com/creditcard/apis/bank"
	"example.com/creditcard/apis/card"
	"example.com/creditcard/apis/constraint"
	"example.com/creditcard/apis/reward"
	"example.com/creditcard/base/psql"
	_ "example.com/creditcard/base/psql"
	"go.uber.org/dig"

	"github.com/braintree/manners"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	bankService "example.com/creditcard/service/bank"
	bankStore "example.com/creditcard/stores/bank"

	cardService "example.com/creditcard/service/card"
	cardStore "example.com/creditcard/stores/card"

	rewardService "example.com/creditcard/service/reward"
	rewardStore "example.com/creditcard/stores/reward"

	constraintService "example.com/creditcard/service/constraint"
	constraintStore "example.com/creditcard/stores/constraint"
)

func BuildContainer() *dig.Container {

	container := dig.New()
	container.Provide(psql.NewPsql) // new postgres

	// bank module
	container.Provide(bankService.New)
	container.Provide(bankStore.New)

	// card module
	container.Provide(cardService.New)
	container.Provide(cardStore.New)

	// privlage module
	container.Provide(rewardService.New)
	container.Provide(rewardStore.New)

	// constraint module
	container.Provide(constraintService.New)
	container.Provide(constraintStore.New)

	// gin server
	container.Provide(NewServer)
	return container
}

func NewServer(
	bankSrc bankService.Service,
	cardSrc cardService.Service,
	rewardSrc rewardService.Service,
	constraintSrc constraintService.Service,

) *gin.Engine {

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	v1 := router.Group("api/v1")

	bank.NewBankHandle(v1.Group("/bank"), bankSrc)
	card.NewCardHandler(v1.Group("/card"), cardSrc)
	reward.NewrewardHandler(v1.Group("/reward"), rewardSrc)
	constraint.NewConstraintHandler(v1.Group("/constraint"), constraintSrc)

	return router
}

func main() {

	container := BuildContainer()

	if err := container.Invoke(func(router *gin.Engine) {
		manners.ListenAndServe(":8080", router)
	}); err != nil {
		panic(err)
	}

	logrus.Info("start serving http request")

}
