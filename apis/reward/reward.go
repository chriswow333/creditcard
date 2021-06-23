package reward

import (
	"net/http"

	"go.uber.org/dig"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"example.com/creditcard/middlewares/apis"
	rewardM "example.com/creditcard/models/reward"
	"example.com/creditcard/service/reward"
)

type rewardHandler struct {
	dig.In

	rewardSrc reward.Service
}

func NewrewardHandler(
	rg *gin.RouterGroup,

	rewardSrc reward.Service,
) {

	ph := &rewardHandler{
		rewardSrc: rewardSrc,
	}

	apis.Handle(rg, http.MethodPost, "", ph.create)
	apis.Handle(rg, http.MethodGet, "/:ID", ph.get)
	apis.Handle(rg, http.MethodPost, "/:ID", ph.updateByID)
	apis.Handle(rg, http.MethodGet, "/cardID/:cardID", ph.getByCardID)
}

func (h *rewardHandler) create(ctx *gin.Context) {
	var rewardModel rewardM.Reward

	ctx.BindJSON(&rewardModel)
	if err := h.rewardSrc.Create(ctx, &rewardModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (h *rewardHandler) get(ctx *gin.Context) {
	ID := ctx.Param("ID")
	reward, err := h.rewardSrc.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, reward)
}

func (h *rewardHandler) getByCardID(ctx *gin.Context) {

	cardID := ctx.Param("cardID")
	rewards, err := h.rewardSrc.GetByCardID(ctx, cardID)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, rewards)
}

func (h *rewardHandler) updateByID(ctx *gin.Context) {
	var rewardModel rewardM.Reward

	ctx.BindJSON(&rewardModel)
	if err := h.rewardSrc.UpdateByID(ctx, &rewardModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
