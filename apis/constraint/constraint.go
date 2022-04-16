package constraint

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/constraint"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type constraintHandler struct {
	dig.In

	constraintService constraint.Service
}

func NewConstraintHandler(
	rg *gin.RouterGroup,
	constraintService constraint.Service,
) {

	ch := &constraintHandler{
		constraintService: constraintService,
	}

	apis.Handle(rg, http.MethodGet, "/customizations/cardID/:ID", ch.getCustomizationByCardID)
	apis.Handle(rg, http.MethodGet, "/customization/:ID", ch.getCustomizationByID)

	apis.Handle(rg, http.MethodGet, "/ecommerces", ch.getCustomizationByID)
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

}

func (h *constraintHandler) getCustomizationByCardID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.constraintService.GetCustomizationsByCardID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getCustomizationByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.constraintService.GetCustomizationByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getEcommerces(ctx *gin.Context) {

	resp, err := h.constraintService.GetAllEcommerces(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getEcommerceByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.constraintService.GetEcommerce(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getDeliveries(ctx *gin.Context) {

	resp, err := h.constraintService.GetAllDeliverys(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getDeliveryByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.constraintService.GetDelivery(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getMobilepays(ctx *gin.Context) {

	resp, err := h.constraintService.GetAllMobilepays(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getMobilepayByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.constraintService.GetMobilepay(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getOnlinegames(ctx *gin.Context) {

	resp, err := h.constraintService.GetAllOnlinegames(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getOnlinegameByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.constraintService.GetOnlinegame(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getStreamings(ctx *gin.Context) {

	resp, err := h.constraintService.GetAllStreamings(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getStreamingByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.constraintService.GetStreaming(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getSupermarkets(ctx *gin.Context) {

	resp, err := h.constraintService.GetAllSupermarkets(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *constraintHandler) getSupermarketByID(ctx *gin.Context) {

	ID := ctx.Param("ID")

	resp, err := h.constraintService.GetSupermarket(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
