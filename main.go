package main

import (
	"example.com/creditcard/apis/bank"
	"example.com/creditcard/apis/card"
	"example.com/creditcard/apis/constraint"
	"example.com/creditcard/apis/customization"
	"example.com/creditcard/apis/delivery"
	"example.com/creditcard/apis/ecommerce"
	"example.com/creditcard/apis/evaluator"
	"example.com/creditcard/apis/mobilepay"
	"example.com/creditcard/apis/onlinegame"
	"example.com/creditcard/apis/reward"
	"example.com/creditcard/apis/streaming"
	"example.com/creditcard/apis/supermarket"
	"example.com/creditcard/base/psql"
	_ "example.com/creditcard/base/psql"
	"go.uber.org/dig"

	"github.com/braintree/manners"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	bankService "example.com/creditcard/service/bank"
	cardService "example.com/creditcard/service/card"
	constraintService "example.com/creditcard/service/constraint"
	customizationService "example.com/creditcard/service/customization"
	deliveryService "example.com/creditcard/service/delivery"
	ecommerceService "example.com/creditcard/service/ecommerce"
	mobilepayService "example.com/creditcard/service/mobilepay"
	onlinegameService "example.com/creditcard/service/onlinegame"
	rewardService "example.com/creditcard/service/reward"
	streamingService "example.com/creditcard/service/streaming"
	supermarketService "example.com/creditcard/service/supermarket"

	bankStore "example.com/creditcard/stores/bank"
	cardStore "example.com/creditcard/stores/card"
	customizationStore "example.com/creditcard/stores/customization"
	deliveryStore "example.com/creditcard/stores/delivery"
	ecommerceStore "example.com/creditcard/stores/ecommerce"
	mobilepayStore "example.com/creditcard/stores/mobilepay"
	onlinegameStore "example.com/creditcard/stores/onlinegame"
	rewardStore "example.com/creditcard/stores/reward"
	streamingStore "example.com/creditcard/stores/streaming"
	supermarketStore "example.com/creditcard/stores/supermarket"

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
	container.Provide(customizationService.New)
	container.Provide(ecommerceService.New)
	container.Provide(mobilepayService.New)
	container.Provide(deliveryService.New)
	container.Provide(onlinegameService.New)
	container.Provide(streamingService.New)
	container.Provide(supermarketService.New)

	// store
	container.Provide(bankStore.New)
	container.Provide(cardStore.New)
	container.Provide(rewardStore.New)
	container.Provide(customizationStore.New)
	container.Provide(ecommerceStore.New)
	container.Provide(mobilepayStore.New)
	container.Provide(deliveryStore.New)
	container.Provide(onlinegameStore.New)
	container.Provide(streamingStore.New)
	container.Provide(supermarketStore.New)

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
	customizationSrc customizationService.Service,
	ecommerceSrc ecommerceService.Service,
	mobilepaySrc mobilepayService.Service,
	deliverySrc deliveryService.Service,
	onlinegameSrc onlinegameService.Service,
	streamingSrc streamingService.Service,
	supermarketSrc supermarketService.Service,

	evaluatorMod evaluatorModule.Module,

) *gin.Engine {

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
	}))

	v1 := router.Group("api/v1")

	bank.NewBankHandle(v1.Group("/bank"), bankSrc)
	card.NewCardHandler(v1.Group("/card"), cardSrc)
	reward.NewrewardHandler(v1.Group("/reward"), rewardSrc)
	constraint.NewConstraintHandler(v1.Group("/constraint"), constraintSrc)
	customization.NewCustomizationHandler(v1.Group("/customization"), customizationSrc)
	ecommerce.NewEcommerceHandler(v1.Group("/ecommerce"), ecommerceSrc)
	mobilepay.NewMobilepayHandler(v1.Group("/mobilepay"), mobilepaySrc)
	delivery.NewDeliveryHandler(v1.Group("/delivery"), deliverySrc)
	onlinegame.NewOnlinegameandler(v1.Group("/onlinegame"), onlinegameSrc)
	streaming.NewStreamingHandler(v1.Group("/streaming"), streamingSrc)
	supermarket.NewSupermarketHandler(v1.Group("/supermarket"), supermarketSrc)

	evaluator.NewEvaluatorHandler(v1.Group("/evaluator"), evaluatorMod)

	return router
}

func main() {

	container := BuildContainer()

	if err := container.Invoke(func(router *gin.Engine) {
		logrus.Info("start serving http request")

		manners.ListenAndServe(":8080", router)
	}); err != nil {
		panic(err)
	}

}
