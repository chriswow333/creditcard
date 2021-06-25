package bank

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/bank"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	bankM "example.com/creditcard/models/bank"
)

type bankHandler struct {
	dig.In

	bankSrc bank.Service
}

func NewBankHandle(
	rg *gin.RouterGroup,

	bankSrc bank.Service,
) {

	bh := &bankHandler{
		bankSrc: bankSrc,
	}

	apis.Handle(rg, http.MethodGet, "", bh.getAll)
	apis.Handle(rg, http.MethodPost, "/:ID", bh.update)
	apis.Handle(rg, http.MethodPost, "", bh.create)
	apis.Handle(rg, http.MethodGet, "/:ID", bh.get)
}

func (h *bankHandler) create(ctx *gin.Context) {
	var bankModel bankM.Bank
	ctx.BindJSON(&bankModel)

	err := h.bankSrc.Create(ctx, &bankModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *bankHandler) update(ctx *gin.Context) {
	var bankModel bankM.Bank
	ctx.BindJSON(&bankModel)

	ID := ctx.Param("ID")
	bankModel.ID = ID

	err := h.bankSrc.Create(ctx, &bankModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *bankHandler) get(ctx *gin.Context) {

	ID := ctx.Param("ID")
	bank, err := h.bankSrc.GetByID(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, bank)
}

func (h *bankHandler) getAll(ctx *gin.Context) {
	banks, err := h.bankSrc.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, banks)
}
