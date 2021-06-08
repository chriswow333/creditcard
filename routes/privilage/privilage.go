package privilage

import (
	"net/http"

	"go.uber.org/dig"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"example.com/creditcard/middlewares/apis"
	privilageM "example.com/creditcard/models/privilage"
	"example.com/creditcard/service/privilage"
)

type privilageHandler struct {
	dig.In

	privilageSrc privilage.Service
}

func NewPrivilageHandler(
	rg *gin.RouterGroup,

	privilageSrc privilage.Service,
) {

	ph := &privilageHandler{
		privilageSrc: privilageSrc,
	}

	apis.Handle(rg, http.MethodGet, "", ph.getAll)
	apis.Handle(rg, http.MethodPost, "", ph.create)
	apis.Handle(rg, http.MethodGet, "/:ID", ph.get)
}

func (h *privilageHandler) create(ctx *gin.Context) {
	var privilageModel privilageM.Privilage

	ctx.BindJSON(&privilageModel)

	if err := h.privilageSrc.Create(ctx, &privilageModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (h *privilageHandler) get(ctx *gin.Context) {
	ID := ctx.Param("ID")
	privilage, err := h.privilageSrc.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, privilage)
}

func (h *privilageHandler) getAll(ctx *gin.Context) {
	privilages, err := h.privilageSrc.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, privilages)
}
