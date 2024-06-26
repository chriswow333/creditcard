package event

import (
	"example.com/creditcard/models/card"
	"example.com/creditcard/models/reward"
)

type SortType int32

const (
	NONE = iota
	MATCH
	MAX_REWARD_BONUS
	MAX_REWARD_RETURN
	MAX_REWARD_EXPECTED_BONUS
)

type Event struct {
	ID string `json:"id"`

	Cash int64 `json:"cash"`

	SortType SortType `json:"sortType"`

	CardIDs       []string `json:"cardIDs"`
	CardRewardIDs []string `json:"cardRewardIDs"`

	RewardType reward.RewardType `json:"rewardType"`

	EffectiveTime int64 `json:"effectiveTime"`

	DefaultCustomization bool `json:"defaultCustomization"`

	ChannelIDs []string `json:"channelIDs"`

	Tasks []string `json:"tasks"`

	Mobilepays        []string `json:"mobilepays"`
	Ecommerces        []string `json:"ecommerces"`
	Supermarkets      []string `json:"supermarkets"`
	Onlinegames       []string `json:"onlinegames"`
	Streamings        []string `json:"streamings"`
	Foods             []string `json:"foods"`
	Transportations   []string `json:"transportations"`
	Deliveries        []string `json:"deliveries"`
	Travels           []string `json:"travels"`
	Insurances        []string `json:"insurances"`
	Malls             []string `json:"malls"`
	Conveniencestores []string `json:"conveniencestores"`
	Sports            []string `json:"sports"`
	AppStores         []string `json:"appstores"`
	Hotels            []string `json:"hotels"`
	Amusements        []string `json:"amusements"`
	Cinemas           []string `json:"cinemas"`
	Publicutilities   []string `json:"publicutilities"`
}

type Response struct {
	EventID        string                `json:"eventID"`
	CardEventResps []*card.CardEventResp `json:"cardEventResps"`
}
