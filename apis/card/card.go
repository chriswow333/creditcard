package card

import (
	"net/http"
	"runtime/debug"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/card"

	cardM "example.com/creditcard/models/card"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

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
	apis.Handle(rg, http.MethodPost, "", ch.createCard)
	apis.Handle(rg, http.MethodPost, "/cardReward", ch.createCardReward)

	apis.Handle(rg, http.MethodPost, "/:ID", ch.update)

	apis.Handle(rg, http.MethodGet, "/:ID", ch.get)

	apis.Handle(rg, http.MethodGet, "/resp/:ID", ch.getResp)

	apis.Handle(rg, http.MethodGet, "/bankID/:bankID", ch.getByBankID)
	apis.Handle(rg, http.MethodPost, "/evaluateConstraintLogic/:ID", ch.evaluateConstraintLogic)

	apis.Handle(rg, http.MethodPost, "/likecard", ch.findLikeName)

}

func (h *cardHandler) createCardReward(ctx *gin.Context) {

	var carRewardModel cardM.CardReward

	ctx.BindJSON(&carRewardModel)

	if err := h.cardSrc.CreateCardReward(ctx, &carRewardModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *cardHandler) createCard(ctx *gin.Context) {
	var cardModel cardM.Card
	ctx.BindJSON(&cardModel)

	err := h.cardSrc.Create(ctx, &cardModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *cardHandler) update(ctx *gin.Context) {
	var cardModel cardM.Card
	ctx.BindJSON(&cardModel)

	ID := ctx.Param("ID")
	cardModel.ID = ID
	err := h.cardSrc.UpdateByID(ctx, &cardModel)

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

func (h *cardHandler) getResp(ctx *gin.Context) {

	ID := ctx.Param("ID")
	cardResp, err := h.cardSrc.GetRespByID(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, cardResp)
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

type ConstraintIDsPayload struct {
	ConstraintIDs []string `json:"constraintIDs"`
}

func (h *cardHandler) evaluateConstraintLogic(ctx *gin.Context) {

	id := ctx.Param("ID")
	var constraintIDsPayload ConstraintIDsPayload

	err := ctx.BindJSON(&constraintIDsPayload)

	pass, msg, err := h.cardSrc.EvaluateConstraintLogic(ctx, id, constraintIDsPayload.ConstraintIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	constraintLogicResp := &ConstraintLogicResp{
		Pass:    pass,
		Meesage: msg,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": constraintLogicResp})

}

type ConstraintLogicResp struct {
	Pass    bool   `json:"pass"`
	Meesage string `json:"message"`
}

type LikeParam struct {
	Likes []string `json:"likes"`
}

func (h *cardHandler) findLikeName(ctx *gin.Context) {

	// name := ctx.Param("name")

	var likeParam LikeParam
	ctx.BindJSON(&likeParam)
	logrus.Info(likeParam)

	resp, err := h.cardSrc.FindByLike(ctx, likeParam.Likes)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}
