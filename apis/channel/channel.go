package channel

import (
	"net/http"
	"runtime/debug"
	"strconv"

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

	apis.Handle(rg, http.MethodGet, "/hotels", ch.getHotels)
	apis.Handle(rg, http.MethodGet, "/hotel/:ID", ch.getHotelByID)

	apis.Handle(rg, http.MethodGet, "/amusements", ch.getAmusements)
	apis.Handle(rg, http.MethodGet, "/amusement/:ID", ch.getAmusementlByID)

	apis.Handle(rg, http.MethodGet, "/cinemas", ch.getCinemas)
	apis.Handle(rg, http.MethodGet, "/cinema/:ID", ch.getCinemaByID)

	apis.Handle(rg, http.MethodGet, "/publicutilities", ch.getPublicUtilities)
	apis.Handle(rg, http.MethodGet, "/publicutility/:ID", ch.getPublicUtilityByID)

	apis.Handle(rg, http.MethodPost, "/likename", ch.findLikeName)

}

func (h *channelHandler) getTasksByCardID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetTasksByCardID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTaskByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetTaskByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
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

	offset := ctx.Query("offset")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		logrus.Error("[Error]{}", err)
		ctx.JSON(http.StatusInternalServerError, "parameter error")
		return
	}

	resp, err := h.channelService.GetAllEcommerces(ctx, offsetInt, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getEcommerceByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetEcommerce(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getDeliveries(ctx *gin.Context) {

	resp, err := h.channelService.GetAllDeliverys(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getDeliveryByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetDelivery(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getMobilepays(ctx *gin.Context) {

	resp, err := h.channelService.GetAllMobilepays(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getMobilepayByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetMobilepay(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getOnlinegames(ctx *gin.Context) {

	resp, err := h.channelService.GetAllOnlinegames(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getOnlinegameByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetOnlinegame(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getStreamings(ctx *gin.Context) {

	resp, err := h.channelService.GetAllStreamings(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getStreamingByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetStreaming(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getSupermarkets(ctx *gin.Context) {

	resp, err := h.channelService.GetAllSupermarkets(ctx, 0, 10)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getSupermarketByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetSupermarket(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTransportations(ctx *gin.Context) {

	resp, err := h.channelService.GetAllTransportations(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTransportationByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetTransportation(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getFoods(ctx *gin.Context) {

	resp, err := h.channelService.GetAllFoods(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getFoodByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetFood(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTravels(ctx *gin.Context) {

	resp, err := h.channelService.GetAllTravels(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getTravelByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetTravel(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getInsurances(ctx *gin.Context) {

	resp, err := h.channelService.GetAllInsurances(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getInsuranceByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetInsurance(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getSports(ctx *gin.Context) {

	resp, err := h.channelService.GetAllSports(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getSportByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetSport(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getConvenienceStores(ctx *gin.Context) {

	resp, err := h.channelService.GetAllConvenienceStores(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getConvenienceStoreByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetConvenienceStore(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getMalls(ctx *gin.Context) {

	resp, err := h.channelService.GetAllMalls(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getMallByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetMall(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getAppstores(ctx *gin.Context) {

	resp, err := h.channelService.GetAllAppstores(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getAppstoreByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetAppstore(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getHotels(ctx *gin.Context) {

	resp, err := h.channelService.GetAllHotels(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getHotelByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetHotel(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getAmusements(ctx *gin.Context) {

	resp, err := h.channelService.GetAllAmusemnets(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getAmusementlByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetAmusement(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getCinemas(ctx *gin.Context) {

	resp, err := h.channelService.GetAllCinemas(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getCinemaByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetCinema(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getPublicUtilities(ctx *gin.Context) {

	resp, err := h.channelService.GetAllPublicUtilities(ctx, 0, 1000)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *channelHandler) getPublicUtilityByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.channelService.GetPublicUtility(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

type NameParam struct {
	Names []string `json:"names"`
}

func (h *channelHandler) findLikeName(ctx *gin.Context) {

	// name := ctx.Param("name")

	var nameParam NameParam
	ctx.BindJSON(&nameParam)
	logrus.Info(nameParam)

	resp, err := h.channelService.FindLike(ctx, nameParam.Names)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}
