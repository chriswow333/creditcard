package customization

type Customization struct {
	ID       string   `json:"id"`
	RewardID string   `json:"rewardID"`
	Name     string   `json:"name"`
	Descs    []string `json:"descs"`
}
