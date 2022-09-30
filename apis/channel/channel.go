package channel

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/models/task"
	"example.com/creditcard/service/channel"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type channelHandler struct {
	dig.In

	channelService channel.Service
}

func NewChannelHandler(
	rg *gin.RouterGroup,
	channelService channel.Service,
) {

	ch := &channelHandler{
		channelService: channelService,
	}

	apis.Handle(rg, http.MethodGet, "/tasks/cardID/:ID", ch.getTasksByCardID)
	apis.Handle(rg, http.MethodGet, "/task/:ID", ch.getTaskByID)
	apis.Handle(rg, http.MethodPost, "/task", ch.createTask)

	apis.Handle(rg, http.MethodGet, "/ecommerces", ch.getEcommerces)
	apis.Handle(rg, http.MethodGet, "/ecommerce/:ID", ch.getEcommerceByID)

	apis.Handle(rg, http.MethodGet, "/deliveries", ch.getDeliveries)
	apis.Handle(rg, http.MethodGet, "/delivery/:ID", ch.getDeliveryByID)

	apis.Handle(rg, http.MethodGet, "/mobilepays", ch.getMobilepays)
	apis.Handle(rg, http.MethodGet, "/mobilepay/:ID", ch.getMobilepayByID)

	apis.Handle(rg, http.MethodGet, "/onlinegames", ch.getOnlinegames)
	apis.Handle(rg, http.MethodGet, "/onlinegame/:ID", ch.getOnlinegameByID)

	apis.Handle(rg, http.MethodGet, "/streamings", ch.getStreamings)
	apis.Handle(rg, http.MethodGet, "/streaming/:ID", ch.getStreamingByID)

	apis.Handle(rg, http.MethodGet, "/supermarkets", ch.getSupermarkets)
	apis.Handle(rg, http.MethodGet, "/supermarkets/:ID", ch.getSupermarketByID)

	apis.Handle(rg, http.MethodGet, "/transportations", ch.getTransportations)
	apis.Handle(rg, http.MethodGet, "/transportation/:ID", ch.getTransportationByID)

	apis.Handle(rg, http.MethodGet, "/foods", ch.getFoods)
	apis.Handle(rg, http.MethodGet, "/food/:ID", ch.getFoodByID)

	apis.Handle(rg, http.MethodGet, "/travels", ch.getTravels)
	apis.Handle(rg, http.MethodGet, "/travel/:ID", ch.getTravelByID)

	apis.Handle(rg, http.MethodGet, "/insurances", ch.getInsurances)
	apis.Handle(rg, http.MethodGet, "/insurance/:ID", ch.getInsuranceByID)

	apis.Handle(rg, http.MethodGet, "/malls", ch.getMalls)
	apis.Handle(rg, http.MethodGet, "/mall/:ID", ch.getMallByID)

	apis.Handle(rg, http.MethodGet, "/sports", ch.getSports)
	apis.Handle(rg, http.MethodGet, "/sport/:ID", ch.getSportByID)

	apis.Handle(rg, http.MethodGet, "/conveniencestores", ch.getConvenienceStores)
	apis.Handle(rg, http.MethodGet, "/conveniencestore/:ID", ch.getConvenienceStoreByID)

	apis.Handle(rg, http.MethodGet, "/appstores", ch.getAppstores)
	apis.Handle(rg, http.MethodGet, "/appstore/:ID", ch.getAppstoreByID)
}

func (h *channelHandler) getTasksByCardID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetTasksByCardID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTaskByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetTaskByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) createTask(ctx *gin.Context) {

	var task task.Task

	ctx.BindJSON(&task)

	err := h.channelService.CreateTask(ctx, &task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *channelHandler) getEcommerces(ctx *gin.Context) {

	resp, err := h.channelService.GetAllEcommerces(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getEcommerceByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetEcommerce(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getDeliveries(ctx *gin.Context) {

	resp, err := h.channelService.GetAllDeliverys(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getDeliveryByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetDelivery(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getMobilepays(ctx *gin.Context) {

	resp, err := h.channelService.GetAllMobilepays(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getMobilepayByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetMobilepay(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getOnlinegames(ctx *gin.Context) {

	resp, err := h.channelService.GetAllOnlinegames(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getOnlinegameByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetOnlinegame(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getStreamings(ctx *gin.Context) {

	resp, err := h.channelService.GetAllStreamings(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getStreamingByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetStreaming(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getSupermarkets(ctx *gin.Context) {

	resp, err := h.channelService.GetAllSupermarkets(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getSupermarketByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetSupermarket(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTransportations(ctx *gin.Context) {

	resp, err := h.channelService.GetAllTransportations(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTransportationByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetTransportation(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getFoods(ctx *gin.Context) {

	resp, err := h.channelService.GetAllFoods(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getFoodByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetFood(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTravels(ctx *gin.Context) {

	resp, err := h.channelService.GetAllTravels(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTravelByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetTravel(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getInsurances(ctx *gin.Context) {

	resp, err := h.channelService.GetAllInsurances(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getInsuranceByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetInsurance(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getSports(ctx *gin.Context) {

	resp, err := h.channelService.GetAllSports(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getSportByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetSport(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getConvenienceStores(ctx *gin.Context) {

	resp, err := h.channelService.GetAllConvenienceStores(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getConvenienceStoreByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetConvenienceStore(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getMalls(ctx *gin.Context) {

	resp, err := h.channelService.GetAllMalls(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getMallByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetMall(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getAppstores(ctx *gin.Context) {

	resp, err := h.channelService.GetAllAppstores(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getAppstoreByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetAppstore(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
