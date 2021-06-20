package bonus

type Bonus struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	Point    float64  `json:"point"`
	UnitType UnitType `json:"unit"`
}

type UnitType int32

const (
	TWD UnitType = iota
	Percentage
)
