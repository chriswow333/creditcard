package channel

import "example.com/creditcard/models/label"

type Supermarket struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	LabelTypes []label.LabelType `json:"labelTypes"`
	ImagePath  string            `json:"imagePath"`
}
