package channel

type Delivery struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	ChannelLabels []ChannelLabel `json:"channelLabels"`
}
