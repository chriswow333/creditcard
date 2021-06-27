package onlinegame

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	onlinegameM "example.com/creditcard/models/onlinegame"
	"example.com/creditcard/service/onlinegame"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type onlinegameHandler struct {
	dig.In

	onlinegameService onlinegame.Service
}

func NewOnlinegameandler(
	rg *gin.RouterGroup,

	onlinegameService onlinegame.Service,

) {
	eh := &onlinegameHandler{
		onlinegameService: onlinegameService,
	}

	apis.Handle(rg, http.MethodPost, "", eh.create)
	apis.Handle(rg, http.MethodPost, "/:ID", eh.updateByID)
	apis.Handle(rg, http.MethodGet, "", eh.getAll)
}

func (h *onlinegameHandler) create(ctx *gin.Context) {
	var onlinegameModel onlinegameM.Onlinegame

	ctx.BindJSON(&onlinegameModel)
	if err := h.onlinegameService.Create(ctx, &onlinegameModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "o"})

}

func (h *onlinegameHandler) updateByID(ctx *gin.Context) {
	var onlinegameModel onlinegameM.Onlinegame

	ctx.BindJSON(&onlinegameModel)

	ID := ctx.Param("ID")
	onlinegameModel.ID = ID
	if err := h.onlinegameService.UpdateByID(ctx, &onlinegameModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *onlinegameHandler) getAll(ctx *gin.Context) {

	onlinegames, err := h.onlinegameService.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, onlinegames)
}
