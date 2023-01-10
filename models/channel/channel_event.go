package channel

type ChannelEventResp struct {
	Pass bool `json:"pass"`

	ChannelType ChannelType `json:"channelType,omitempty"`

	ChannelOperatorType ChannelOperatorType `json:"channelOperatorType,omitempty"`

	ChannelMappingType ChannelMappingType `json:"channelMappingType,omitempty"`

	ChannelEventResps []*ChannelEventResp `json:"channelEventResps,omitempty"`

	Matches []string `json:"matches"`
	Misses  []string `json:"misses"`
}
