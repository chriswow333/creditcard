package customization

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	customizationM "example.com/creditcard/models/customization"
	"example.com/creditcard/service/customization"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type customizationHandler struct {
	dig.In

	customizationService customization.Service
}

func NewCustomizationHandler(
	rg *gin.RouterGroup,

	customizationService customization.Service,
) {
	ch := &customizationHandler{
		customizationService: customizationService,
	}

	apis.Handle(rg, http.MethodPost, "", ch.create)
	apis.Handle(rg, http.MethodGet, "/:ID", ch.getByID)
	apis.Handle(rg, http.MethodPost, "/:ID", ch.update)
	apis.Handle(rg, http.MethodGet, "/rewardID/:ID", ch.getByRewardID)
}

func (h *customizationHandler) create(ctx *gin.Context) {
	var customizationModel customizationM.Customization
	ctx.BindJSON(&customizationModel)

	err := h.customizationService.Create(ctx, &customizationModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *customizationHandler) getByID(ctx *gin.Context) {
	ID := ctx.Param("ID")
	customization, err := h.customizationService.GetByID(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, customization)
}

func (h *customizationHandler) update(ctx *gin.Context) {

	var customizationModel customizationM.Customization
	ctx.BindJSON(&customizationModel)

	ID := ctx.Param("ID")
	customizationModel.ID = ID

	err := h.customizationService.UpdateByID(ctx, &customizationModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (h *customizationHandler) getByRewardID(ctx *gin.Context) {

	ID := ctx.Param("ID")
	customizations, err := h.customizationService.GetByRewardID(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, customizations)
}
