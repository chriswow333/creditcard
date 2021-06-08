package bonus

type Bonus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`

	score float64 `json:"score"`
}
