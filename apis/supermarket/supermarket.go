package supermarket

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/supermarket"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	supermarketM "example.com/creditcard/models/supermarket"
)

type supermarketHandler struct {
	dig.In

	supermarketService supermarket.Service
}

func NewEcommerceHandler(
	rg *gin.RouterGroup,

	supermarketService supermarket.Service,

) {
	eh := &supermarketHandler{
		supermarketService: supermarketService,
	}

	apis.Handle(rg, http.MethodPost, "", eh.create)
	apis.Handle(rg, http.MethodPost, "/:ID", eh.updateByID)
	apis.Handle(rg, http.MethodGet, "", eh.getAll)
}

func (h *supermarketHandler) create(ctx *gin.Context) {
	var supermarketModel supermarketM.Supermarket

	ctx.BindJSON(&supermarketModel)

	if err := h.supermarketService.Create(ctx, &supermarketModel); err != nil {

		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *supermarketHandler) updateByID(ctx *gin.Context) {
	var supermarketModel supermarketM.Supermarket

	ctx.BindJSON(&supermarketModel)
	ID := ctx.Param("ID")
	supermarketModel.ID = ID
	if err := h.supermarketService.UpdateByID(ctx, &supermarketModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *supermarketHandler) getAll(ctx *gin.Context) {
	supermarkets, err := h.supermarketService.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, supermarkets)
}
