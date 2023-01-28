package feedback

type PointCalculateType int32

const (
	FIXED_POINT_RETURN = iota + 1
	BONUS_MULTIPLY_POINT
)

type Pointback struct {
	PointCalculateType PointCalculateType `json:"pointCalculateType,omitempty"`

	Fixed float64 `json:"fixed,omitempty"`
	Bonus float64 `json:"bonus,omitempty"`
	Min   int64   `json:"min"`
	Max   int64   `json:"max"`
}

type PointReturnStatus int32

const (
	ALL_RETURN_POINT = iota
	SOME_RETURN_POINT
	NONE_RETURN_POINT
)

type PointReturn struct {
	Cash int64 `json:"cash"` // 花費金額

	PointReturnStatus PointReturnStatus `json:"pointReturnStatus"`

	ActualUseCash     int64   `json:"actualUseCash"`     // 實際用多少錢得到回饋
	ActualPointReturn float64 `json:"actualPointReturn"` // 回饋多少point

	// TotalBonus       float64 `json:"totalBonus"`     // 總%數
	PointReturnBonus float64 `json:"pointReturnBonus"` // 實際拿多少%回饋
}

type PointFeedbackBonus struct {
	TotalBonus         float64            `json:"totalBonus"`
	PointCalculateType PointCalculateType `json:"pointCalculateType"`

	Title                  string `json:"title"`                  // ex. LINE POINTS
	ReturnBonusTitle       string `json:"returnBonusTitle"`       // ex. 3%"現金回饋"
	PointReturnTitlePrefix string `json:"pointReturnTitlePrefix"` // ex. "現省"xx元
	PointReturnTitleSuffix string `json:"pointReturnTitleSuffix"` // ex. 現省xx"元"
}
