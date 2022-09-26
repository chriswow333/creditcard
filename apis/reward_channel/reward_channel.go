package reward_channel

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	rewardChannelM "example.com/creditcard/models/reward_channel"
	"example.com/creditcard/service/reward_channel"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type rewardChannelHandler struct {
	dig.In

	rewardChannelSrc reward_channel.Service
}

func NewRewardChannelHandler(
	rg *gin.RouterGroup,

	rewardChannelSrc reward_channel.Service,
) {

	h := &rewardChannelHandler{
		rewardChannelSrc: rewardChannelSrc,
	}

	apis.Handle(rg, http.MethodPost, "", h.createRewardChannel)
}

type RewardChannels struct {
	RewardChannels []*rewardChannelM.RewardChannel `json:"rewardChannels"`
}

func (h *rewardChannelHandler) createRewardChannel(ctx *gin.Context) {

	var rewardChannelsModel RewardChannels

	ctx.BindJSON(&rewardChannelsModel)

	err := h.rewardChannelSrc.Create(ctx, rewardChannelsModel.RewardChannels)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
