package feedback

// 點數回饋

type PointbackType int32

const (
	LINE_POINT PointbackType = iota + 1
	KUO_BROTHERS_POINT
	WOWPRIME_POINT
	OPEN_POINT
	YIDA_POINT
)

type PointCalculateType int32

const (
	FIXED_POINT_RETURN = iota + 1
	BONUS_MULTIPLY_POINT
)

type Pointback struct {
	PointbackType      PointbackType      `json:"pointbackType,omitempty"`
	PointCalculateType PointCalculateType `json:"pointCalculateType,omitempty"`

	Fixed float64 `json:"fixed,omitempty"`
	Bonus float64 `json:"bonus,omitempty"`
	Min   int64   `json:"min"`
	Max   int64   `json:"max"`
}

type PointReturn struct {
	IsPointbackGet bool    `json:"isPointbackGet"`
	PointbackBonus float64 `json:"pointbackBonus"`

	ActualUseCash   int64   `json:"actualUseCash"` // 實際用多少錢得到回饋
	ActualPointBack float64 `json:"actualPointBack"`

	CurrentCash int64   `json:"current"`
	TotalCash   float64 `json:"total"`
}
