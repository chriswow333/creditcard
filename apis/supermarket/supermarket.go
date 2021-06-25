package supermarket

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/supermarket"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
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

}

func (h *supermarketHandler) updateByID(ctx *gin.Context) {

}

func (h *supermarketHandler) getAll(ctx *gin.Context) {

}
