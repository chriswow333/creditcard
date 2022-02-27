package customization

type Customization struct {
	ID          string `json:"id"`
	RewardID    string `json:"rewardID"`
	DefaultPass bool   `json:"defaultPass"`
}
