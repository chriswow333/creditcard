package channel

type ChannelOperatorType int32

const (
	AND ChannelOperatorType = iota + 1
	OR
)

type ChannelMappingType int32

const (
	MATCH ChannelMappingType = iota + 1
	MISMATCH
)

type ChannelType int32

const (
	InnerChannelType ChannelType = iota + 1 //  abstract layer, there are several nested layers.

	TaskType        // setting layer, ex. 綁定數位帳號,
	MobilepayType   // ex. line pay, google pay.
	EcommerceType   // ex. shopee, momo
	SupermarketType // ex. px mart
	OnlinegameType  // ex.
	StreamingType   // ex. netflix
	FoodType
	TransportationType
	TravelType
	DeliveryType
	InsuranceType
	MallType
	SportType
	ConvenienceStoreType
	AppStoreType
)

type Channel struct {
	ChannelOperatorType ChannelOperatorType `json:"channelOperatorType,omitempty"`

	ChannelType ChannelType `json:"channelType,omitempty"`

	ChannelMappingType ChannelMappingType `json:"channelMappingType,omitempty"`

	InnerChannels []*Channel `json:"innerChannels,omitempty"`

	Tasks []string `json:"tasks,omitempty"`

	Mobilepays        []string `json:"mobilepays,omitempty"`
	Ecommerces        []string `json:"ecommerces,omitempty"`
	Supermarkets      []string `json:"supermarkets,omitempty"`
	Onlinegames       []string `json:"onlinegames,omitempty"`
	Streamings        []string `json:"streamings,omitempty"`
	Foods             []string `json:"foods,omitempty"`
	Transportations   []string `json:"transportations,omitempty"`
	Deliveries        []string `json:"deliveries,omitempty"`
	Travels           []string `json:"travels,omitempty"`
	Insurances        []string `json:"insurances,omitempty"`
	Malls             []string `json:"malls,omitempty"`
	Conveniencestores []string `json:"conveniencestores,omitempty"`
	Sports            []string `json:"sports,omitempty"`
	AppStores         []string `json:"appstores,omitempty"`
}
