package channel

type Food struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	ChannelLabels []ChannelLabel `json:"channelLabels"`
	ImagePath     string         `json:"imagePath"`
}
