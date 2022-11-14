package channel

type ChannelEventResp struct {
	Pass bool `json:"pass"`

	AllPass bool `json:"allPass"`

	ChannelType ChannelType `json:"channelType,omitempty"`

	ChannelOperatorType ChannelOperatorType `json:"channelOperatorType,omitempty"`

	ChannelMappingType ChannelMappingType `json:"channelMappingType,omitempty"`

	ChannelEventResps []*ChannelEventResp `json:"channelEventResps,omitempty"`

	Matches []string `json:"matches"`
	Misses  []string `json:"misses"`
}
