package bonus

type BonusType int32

const (
	TWD BonusType = iota
	Percentage
)

type Bonus struct {
	Point     float64   `json:"point"`
	BonusType BonusType `json:"bonusType"`
}

type BonusLimit struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	BonusType BonusType `json:"bonusType"`
	AtLeast   float64   `json:"atLeast"`
	AtMost    float64   `json:"atMost"`
}
