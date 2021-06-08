package models

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model

	Year      int
	Make      string
	ModelName string
	DriverID  int
}
