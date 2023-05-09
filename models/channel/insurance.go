package channel

type Insurance struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	ChannelLabels []ChannelLabel `json:"channelLabels"`
	ImagePath     string         `json:"imagePath"`
}
