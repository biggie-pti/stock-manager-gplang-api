package stock

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model
	Name        string `gorm:"Not Null" json:"name"`
	Description string `json:"description"`
	      string `gorm:"Not Null" json:"status"`
}
