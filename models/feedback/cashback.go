package feedback

type CashCalculateType int32

const (
	FIXED_CASH_RETURN = iota + 1
	BONUS_MULTIPLY_CASH
)

type Cashback struct {
	CashCalculateType CashCalculateType `json:"cashCalculateType,omitempty"`

	Fixed float64 `json:"fixed,omitempty"`
	Bonus float64 `json:"bonus,omitempty"`

	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

type CashReturnStatus int32

const (
	ALL_RETURN_CASH = iota
	SOME_RETURN_CASH
	NONE_RETURN_CASH
)

type CashReturn struct {
	Cash int64 `json:"cash"` // 花費金額

	CashReturnStatus CashReturnStatus `json:"cashReturnStatus"`

	ActualUseCash    int64   `json:"actualUseCash"`    // 實際用多少錢得到回饋
	ActualCashReturn float64 `json:"actualCashReturn"` // 回饋多少cash

	// TotalBonus      float64 `json:"totalBonus"`      // 總%數
	CashReturnBonus float64 `json:"cashReturnBonus"` // 實際拿多少%回饋

}

type CashFeedbackBonusType int32

const (
	PERCENTAGE = iota + 1 // %數回饋
	DISCOUNT              // 打幾折

)

type CashFeedbackBonus struct {
	CashFeedbackBonusType CashFeedbackBonusType `json:"cashFeedbackBonusType"`
	CashCalculateType     CashCalculateType     `json:"cashCalculateType"`

	TotalBonus float64 `json:"totalBonus"` // ex. "10%"回饋 or 可以是"9"折優惠

	Title                 string `json:"title"`                 // like 現金回饋
	ReturnBonusTitle      string `json:"returnBonusTitle"`      // ex. 3%"現金回饋"
	CashReturnTitlePrefix string `json:"cashReturnTitlePrefix"` // ex. "現省"xx元
	CashReturnTitleSuffix string `json:"cashReturnTitleSuffix"` // ex. 現省xx"元"

}

// var MULTIPLY_CASH_RETURN_TITLE = &CashFeedbackBonusTitle{
// 	Title:                 "現金回饋",
// 	ReturnBonusTitle:      "現金回饋",
// 	CashReturnTitlePrefix: "現省",
// 	CashReturnTitleSuffix: "元",
// }

// var FIXED_CASH_RETURN_TITLE = &CashTitle{
// 	Title:                 "現金回饋",
// 	ReturnBonusTitle:      "現金回饋",
// 	CashReturnTitlePrefix: "現省",
// 	CashReturnTitleSuffix: "元",
// }
