package bonus

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/bonus"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	bonusM "example.com/creditcard/models/bonus"
)

type bonusHandler struct {
	dig.In

	bonusSrc bonus.Service
}

func NewBonusHandler(
	rg *gin.RouterGroup,

	bonusSrc bonus.Service,

) {

	bh := &bonusHandler{
		bonusSrc: bonusSrc,
	}

	apis.Handle(rg, http.MethodGet, "/rewardID/:ID", bh.get)
	apis.Handle(rg, http.MethodPost, "/rewardID/:ID", bh.update)

}

func (h *bonusHandler) update(ctx *gin.Context) {
	rewardID := ctx.Param("ID")
	var bonusModel *bonusM.Bonus
	ctx.BindJSON(&bonusModel)

	if err := h.bonusSrc.UpdateByRewardID(ctx, rewardID, bonusModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *bonusHandler) get(ctx *gin.Context) {

	rewardID := ctx.Param("ID")
	bonus, err := h.bonusSrc.GetByRewardID(ctx, rewardID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, bonus)
}
