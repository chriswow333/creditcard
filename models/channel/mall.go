package channel

import "example.com/creditcard/models/label"

type Mall struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	LabelTypes []label.LabelType `json:"labelTypes"`
	ImagePath  string            `json:"imagePath"`
}
