package customization

type Customization struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Descs []string `json:"desc"`
}
