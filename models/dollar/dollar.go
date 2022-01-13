package dollar

// 現金回饋

type DollarType int32

const (
	NTD DollarType = iota
)

type Dollar struct {
	Current int64 `json:"current,omitempty"` // 該單筆花費多少
	Total   int64 `json:"total,omitempty"`   // 目前總共花費多少

	DollarBonusBack int64     `json:"dollarBonusGet,omitempty"` // 實際用多少錢得到回饋
	DollarBack      float64   `json:"dollarBack,omitempty"`     // 回饋多少
	PointBack       PointBack `json:"pointBack,omitempty"`      // 是否回饋全拿

	DollarType  DollarType  `json:"currency"`
	DollarLimit DollarLimit `json:"dollarLimit"`
}

type DollarLimit struct {
	Point float64 `json:"point"` // 現金回饋％數
	Min   int64   `json:"min"`
	Max   int64   `json:"max"`
}

type PointBack int32

const (
	Full = iota
	PartOf
	None
)
