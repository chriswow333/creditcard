package main

import (
	"example.com/creditcard/apis/bank"
	"example.com/creditcard/apis/bonus"
	"example.com/creditcard/apis/card"
	"example.com/creditcard/apis/constraint"
	"example.com/creditcard/apis/evaluator"
	"example.com/creditcard/apis/reward"
	"example.com/creditcard/base/psql"
	_ "example.com/creditcard/base/psql"
	"go.uber.org/dig"

	"github.com/braintree/manners"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	bankService "example.com/creditcard/service/bank"
	bonusService "example.com/creditcard/service/bonus"
	cardService "example.com/creditcard/service/card"
	constraintService "example.com/creditcard/service/constraint"
	rewardService "example.com/creditcard/service/reward"

	bankStore "example.com/creditcard/stores/bank"
	cardStore "example.com/creditcard/stores/card"
	rewardStore "example.com/creditcard/stores/reward"

	cardrewardBuilder "example.com/creditcard/builder/cardreward"

	evaluatorModule "example.com/creditcard/modules/evaluator"
)

func BuildContainer() *dig.Container {

	container := dig.New()
	container.Provide(psql.NewPsql) // new postgres

	// service
	container.Provide(bankService.New)
	container.Provide(cardService.New)
	container.Provide(rewardService.New)
	container.Provide(constraintService.New)
	container.Provide(bonusService.New)

	// store
	container.Provide(bankStore.New)
	container.Provide(cardStore.New)
	container.Provide(rewardStore.New)

	// builder
	container.Provide(cardrewardBuilder.New)

	// module
	container.Provide(evaluatorModule.New)

	// gin server
	container.Provide(NewServer)
	return container
}

func NewServer(
	bankSrc bankService.Service,
	cardSrc cardService.Service,
	rewardSrc rewardService.Service,
	constraintSrc constraintService.Service,
	bonusSrc bonusService.Service,

	evaluatorMod evaluatorModule.Module,

) *gin.Engine {

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	v1 := router.Group("api/v1")

	bank.NewBankHandle(v1.Group("/bank"), bankSrc)
	card.NewCardHandler(v1.Group("/card"), cardSrc)
	reward.NewrewardHandler(v1.Group("/reward"), rewardSrc)
	constraint.NewConstraintHandler(v1.Group("/constraint"), constraintSrc)
	bonus.NewBonusHandler(v1.Group("/bonus"), bonusSrc)

	evaluator.NewEvaluatorHandler(v1.Group("/evaluator"), evaluatorMod)

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
