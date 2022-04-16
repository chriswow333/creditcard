package customization

type CustomizationType int32

const (
	NONE CustomizationType = iota + 1
	CASH
)

type Customization struct {
	ID string `json:"id"`

	Name   string `json:"name"`
	Desc   string `json:"desc"`
	CardID string `json:"cardID"`

	DefaultPass bool `json:"defaultPass"`

	CustomizationType CustomizationType `json:"customizationType"`

	CustomizationTypeModel CustomizationTypeModel `json:"customizationTypeModel,omitempty"`
}

type CustomizationTypeModel struct {
	CashLimit CashLimit `json:"cashLimit,omitempty"`
}

type CashLimit struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}
