package customization

type Customization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	CardID      string `json:"cardID"`
	DefaultPass bool   `json:"defaultPass"`
}
