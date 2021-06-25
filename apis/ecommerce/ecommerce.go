package ecommerce

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/ecommerce"
	"github.com/gin-gonic/gin"
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

}

func (h *ecommerceHandler) updateByID(ctx *gin.Context) {

}

func (h *ecommerceHandler) getAll(ctx *gin.Context) {

}
