package card

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/card"

	cardM "example.com/creditcard/models/card"

	"github.com/gin-gonic/gin"

	"go.uber.org/dig"
)

type cardHandler struct {
	dig.In

	cardSrc card.Service
}

func NewCardHandler(
	rg *gin.RouterGroup,

	cardSrc card.Service,
) {

	ch := &cardHandler{
		cardSrc: cardSrc,
	}

	apis.Handle(rg, http.MethodGet, "", ch.getAll)
	apis.Handle(rg, http.MethodPost, "", ch.create)
	apis.Handle(rg, http.MethodGet, "/:ID", ch.get)
	apis.Handle(rg, http.MethodGet, "/bankID/:bankID", ch.getByBankID)
}

func (h *cardHandler) create(ctx *gin.Context) {
	var cardModel cardM.Card
	ctx.BindJSON(&cardModel)

	err := h.cardSrc.Create(ctx, &cardModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *cardHandler) get(ctx *gin.Context) {

	ID := ctx.Param("ID")
	card, err := h.cardSrc.GetByID(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, card)
}

func (h *cardHandler) getAll(ctx *gin.Context) {
	cards, err := h.cardSrc.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, cards)
}

func (h *cardHandler) getByBankID(ctx *gin.Context) {

	bankID := ctx.Param("bankID")
	cards, err := h.cardSrc.GetByBankID(ctx, bankID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, cards)
}
