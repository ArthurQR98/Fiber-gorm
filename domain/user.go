package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Orders []Order
	Data   string `gorm:"type:JSONB"`
}

func (User) TableName() string {
	return "user"
}
