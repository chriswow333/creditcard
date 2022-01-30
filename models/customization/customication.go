package customization

type Customization struct {
	ID       string `json:"id"`
	CardID   string `json:"cardID"`
	RewardID string `json:"rewardID"`

	Name        string `json:"name"`
	Desc        string `json:"desc"`
	DefaultPass bool   `json:"defaultPass"`
	LinkURL     string `json:"linkURL"`
}
