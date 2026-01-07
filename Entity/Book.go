package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string  `valid:"stringlength(3|100)~Title must atleast 3 and limit at 100"`
	Price float64 `valid:"range(50|5000)~Price must be between 50 and 5000"`
	Code  string  `valid:"matches(^BK[0-9]{6}$)~Code must start with BK followed by 6 digits (0-9)"`
}
