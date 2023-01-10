package reward_channel

type RewardChannel struct {
	ID           string `json:"id"`
	Order        int32  `json:"order"`
	CardID       string `json:"cardID"`
	CardRewardID string `json:"cardRewardID"`
	ChannelID    string `json:"channelID"`
	ChannelType  int32  `json:"channelType"`
}
