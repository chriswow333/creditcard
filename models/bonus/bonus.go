package bonus

// 點數回饋

type BonusType int32

const (
	TWD BonusType = iota
	Percentage
)

type Bonus struct {
	Point      float64    `json:"point"`
	BonusType  BonusType  `json:"bonusType"`
	BonusLimit BonusLimit `json:"bonusLimit"`
}

type BonusLimit struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Desc string  `json:"desc"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
}
