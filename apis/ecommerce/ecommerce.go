package ecommerce

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	ecommerceM "example.com/creditcard/models/ecommerce"
	"example.com/creditcard/service/ecommerce"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type ecommerceHandler struct {
	dig.In

	ecommerceService ecommerce.Service
}

func NewEcommerceHandler(
	rg *gin.RouterGroup,

	ecommerceService ecommerce.Service,

) {
	eh := &ecommerceHandler{
		ecommerceService: ecommerceService,
	}

	apis.Handle(rg, http.MethodPost, "", eh.create)
	apis.Handle(rg, http.MethodPost, "/:ID", eh.updateByID)
	apis.Handle(rg, http.MethodGet, "", eh.getAll)
}

func (h *ecommerceHandler) create(ctx *gin.Context) {
	var ecommerceModel ecommerceM.Ecommerce
	ctx.BindJSON(&ecommerceModel)

	if err := h.ecommerceService.Create(ctx, &ecommerceModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (h *ecommerceHandler) updateByID(ctx *gin.Context) {

	var ecommerceModel ecommerceM.Ecommerce

	ctx.BindJSON(&ecommerceModel)
	ID := ctx.Param("ID")
	ecommerceModel.ID = ID

	if err := h.ecommerceService.UpdateByID(ctx, &ecommerceModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *ecommerceHandler) getAll(ctx *gin.Context) {
	ecommerces, err := h.ecommerceService.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, ecommerces)
}
