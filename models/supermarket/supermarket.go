package supermarket

type Supermarket struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Desc    string `json:"desc,omitempty"`
	LinkURL string `json:"linkURL,omitempty"`
}
