package customization

type Customization struct {
	ID       string `json:"id"`
	RewardID string `json:"rewardID"`
	Name     string `json:"name,omitempty"`
	Desc     string `json:"desc,omitempty"`
}
