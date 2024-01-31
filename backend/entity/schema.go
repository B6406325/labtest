package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    string `gorm:"uniqueIndex" valid:"required~ID is required, matches(^[B]\\d{7}$)"`
	Name  string `valid:"required~Name is required"`
	Phone string `valid:"required~Phone is required, stringlength(10|10)"`
	Age   int    `valid:"required~Age is required, int~Age is not number"`
}
