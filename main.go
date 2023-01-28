package main

import (
	"runtime/debug"

	"example.com/creditcard/base/psql"
	_ "example.com/creditcard/base/psql"
	"go.uber.org/dig"

	"github.com/braintree/manners"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"example.com/creditcard/apis/bank"
	"example.com/creditcard/apis/card"
	"example.com/creditcard/apis/channel"
	"example.com/creditcard/apis/evaluator"
	"example.com/creditcard/apis/image"
	"example.com/creditcard/apis/reward"
	"example.com/creditcard/apis/reward_channel"

	bankService "example.com/creditcard/service/bank"
	cardService "example.com/creditcard/service/card"
	channelService "example.com/creditcard/service/channel"
	rewardService "example.com/creditcard/service/reward"
	rewardChannelService "example.com/creditcard/service/reward_channel"

	cardrewardBuilder "example.com/creditcard/builder/cardreward"
	amusementStore "example.com/creditcard/stores/amusement"
	appStoreStore "example.com/creditcard/stores/appstore"
	bankStore "example.com/creditcard/stores/bank"
	cardStore "example.com/creditcard/stores/card"
	cardRewardStore "example.com/creditcard/stores/card_reward"
	cinemaStore "example.com/creditcard/stores/cinema"
	conveniencestoreStore "example.com/creditcard/stores/conveniencestore"
	deliveryStore "example.com/creditcard/stores/delivery"
	ecommerceStore "example.com/creditcard/stores/ecommerce"

	// feedbackDescStore "example.com/creditcard/stores/feedback_desc"
	evaluatorModule "example.com/creditcard/modules/evaluator"
	foodStore "example.com/creditcard/stores/food"
	hotelStore "example.com/creditcard/stores/hotel"
	insuranceStore "example.com/creditcard/stores/insurance"
	mallStore "example.com/creditcard/stores/mall"
	mobilepayStore "example.com/creditcard/stores/mobilepay"
	onlinegameStore "example.com/creditcard/stores/onlinegame"
	payloadStore "example.com/creditcard/stores/payload"
	publicutilityStore "example.com/creditcard/stores/publicutility"
	rewardStore "example.com/creditcard/stores/reward"
	rewardChannelStore "example.com/creditcard/stores/reward_channel"
	sportStore "example.com/creditcard/stores/sport"
	streamingStore "example.com/creditcard/stores/streaming"
	supermarketStore "example.com/creditcard/stores/supermarket"
	taskStore "example.com/creditcard/stores/task"
	transportationStore "example.com/creditcard/stores/transportation"
	travelStore "example.com/creditcard/stores/travel"
)

func BuildContainer() *dig.Container {

	container := dig.New()
	container.Provide(psql.NewPsql) // new postgres

	// service
	container.Provide(bankService.New)
	container.Provide(cardService.New)
	container.Provide(rewardService.New)
	container.Provide(channelService.New)
	container.Provide(rewardChannelService.New)
	// container.Provide(feedbackDescService.New)

	// store
	container.Provide(bankStore.New)
	container.Provide(cardStore.New)
	container.Provide(rewardStore.New)
	container.Provide(ecommerceStore.New)
	container.Provide(mobilepayStore.New)
	container.Provide(deliveryStore.New)
	container.Provide(onlinegameStore.New)
	container.Provide(streamingStore.New)
	container.Provide(supermarketStore.New)
	container.Provide(payloadStore.New)
	container.Provide(taskStore.New)
	container.Provide(cardRewardStore.New)
	container.Provide(rewardChannelStore.New)
	container.Provide(transportationStore.New)
	container.Provide(foodStore.New)
	container.Provide(travelStore.New)
	container.Provide(insuranceStore.New)
	container.Provide(conveniencestoreStore.New)
	container.Provide(mallStore.New)
	container.Provide(sportStore.New)
	container.Provide(appStoreStore.New)
	container.Provide(hotelStore.New)
	container.Provide(amusementStore.New)
	container.Provide(cinemaStore.New)
	container.Provide(publicutilityStore.New)

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
	channelSrc channelService.Service,
	rewardChannelSrc rewardChannelService.Service,
	// feedbackSrc feedbackDescService.Service,

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
	reward_channel.NewRewardChannelHandler(v1.Group("/reward_channel"), rewardChannelSrc)

	channel.NewChannelHandler(v1.Group("/channel"), channelSrc)

	evaluator.NewEvaluatorHandler(v1.Group("/evaluator"), evaluatorMod)

	image.NewImageHandler(v1.Group("/image"))

	// feedback_desc.NewFeedbackDescHandler(v1.Group("/feedback_desc"), feedbackSrc)

	return router
}

func main() {

	container := BuildContainer()

	if err := container.Invoke(func(router *gin.Engine) {
		logrus.Info("start serving http request")

		manners.ListenAndServe(":8080", router)
	}); err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		panic(err)
	}

}
