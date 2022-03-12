package customization

type Customization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CardID      string `json:"cardID"`
	DefaultPass bool   `json:"defaultPass"`
}
