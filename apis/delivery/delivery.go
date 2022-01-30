package delivery

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	deliveryM "example.com/creditcard/models/delivery"
	"example.com/creditcard/service/delivery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type deliveryHandler struct {
	dig.In

	deliveryService delivery.Service
}

func NewDeliveryHandler(
	rg *gin.RouterGroup,

	deliveryService delivery.Service,

) {
	eh := &deliveryHandler{
		deliveryService: deliveryService,
	}

	apis.Handle(rg, http.MethodPost, "", eh.create)
	apis.Handle(rg, http.MethodPost, "/:ID", eh.updateByID)
	apis.Handle(rg, http.MethodGet, "", eh.getAll)
}

func (h *deliveryHandler) create(ctx *gin.Context) {
	var deliveryModel deliveryM.Delivery
	ctx.BindJSON(&deliveryModel)

	if err := h.deliveryService.Create(ctx, &deliveryModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (h *deliveryHandler) updateByID(ctx *gin.Context) {

	var deliveryModel deliveryM.Delivery

	ctx.BindJSON(&deliveryModel)
	ID := ctx.Param("ID")
	deliveryModel.ID = ID

	if err := h.deliveryService.UpdateByID(ctx, &deliveryModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *deliveryHandler) getAll(ctx *gin.Context) {
	deliverys, err := h.deliveryService.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, deliverys)
}
