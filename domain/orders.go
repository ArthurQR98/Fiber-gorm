package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	User   User   `gorm:"foreignKey:UserID;references:ID"`
	Data   string `gorm:"type:JSONB"`
}

func (Order) TableName() string {
	return "order"
}
