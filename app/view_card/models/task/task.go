package task

/*
	回饋各項任務
*/

type Task struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	RewardID   string  `json:"rewardID"`
	Point      float64 `json:"point"`
	UpdateDate int64   `json:"updateDate"`
}

type Repr struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	RewardID string  `json:"rewardID"`
	Point    float64 `json:"point"`
}
