package channel

import (
	"example.com/creditcard/models/task"
)

type ChannelResp struct {
	// ChannelOperatorType ChannelOperatorType `json:"ChannelOperatorType,omitempty"`
	// ChannelMappingType ChannelMappingType `json:"ChannelMappingType,omitempty"`

	ChannelType ChannelType `json:"channelType,omitempty"`

	Tasks []*task.Task `json:"tasks,omitempty"`

	Mobilepays        []*Mobilepay        `json:"mobilepays,omitempty"`
	Ecommerces        []*Ecommerce        `json:"ecommerces,omitempty"`
	Supermarkets      []*Supermarket      `json:"supermarkets,omitempty"`
	Onlinegames       []*Onlinegame       `json:"onlinegames,omitempty"`
	Streamings        []*Streaming        `json:"streamings,omitempty"`
	Foods             []*Food             `json:"foods,omitempty"`
	Transportations   []*Transportation   `json:"transportations,omitempty"`
	Travels           []*Travel           `json:"travels,omitempty"`
	Deliveries        []*Delivery         `json:"deliveries,omitempty"`
	Insurances        []*Insurance        `json:"insurances,omitempty"`
	Malls             []*Mall             `json:"malls,omitempty"`
	ConvenienceStores []*ConvenienceStore `json:"conveniencestores,omitempty"`
	Sports            []*Sport            `json:"sports,omitempty"`
	Appstores         []*AppStore         `json:"appstores,omitempty"`
}
