package main

import (
	"example.com/creditcard/base/psql"
	_ "example.com/creditcard/base/psql"
	"example.com/creditcard/routes/bank"
	"example.com/creditcard/routes/card"
	"example.com/creditcard/routes/constraint"
	"example.com/creditcard/routes/privilage"
	"go.uber.org/dig"

	"github.com/braintree/manners"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	bankService "example.com/creditcard/service/bank"
	bankStore "example.com/creditcard/stores/bank"

	cardService "example.com/creditcard/service/card"
	cardStore "example.com/creditcard/stores/card"

	privilageService "example.com/creditcard/service/privilage"
	privilageStore "example.com/creditcard/stores/privilage"

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
	container.Provide(privilageService.New)
	container.Provide(privilageStore.New)

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
	privilageSrc privilageService.Service,
	constraintSrc constraintService.Service,

) *gin.Engine {

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	v1 := router.Group("api/v1")

	bank.NewBankHandle(v1.Group("/bank"), bankSrc)
	card.NewCardHandler(v1.Group("/card"), cardSrc)
	privilage.NewPrivilageHandler(v1.Group("/privilage"), privilageSrc)
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
