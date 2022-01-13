package bonus

// 點數回饋

type BonusType int32

const (
	TWD BonusType = iota
	Percentage
)

type Bonus struct {
	Point float64 `json:"point,omitempty"`

	BonusType BonusType `json:"bonusType"`

	BonusLimit BonusLimit `json:"bonusLimit,omitempty"`
}

type BonusLimit struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`

	Desc string `json:"desc,omitempty"`

	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}
