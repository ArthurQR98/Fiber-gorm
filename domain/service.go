package domain

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service , interface defining all CRUD operations
type Service interface {
	Find(id int) (*User, error)
	Store(user *User) error
	Update(id int, user *User) error
	FindAll() ([]*User, error)
	Delete(id int) error
}

// Repository , interface acting like a port for the database implementation
type Repository interface {
	Find(id int) (*User, error)
	Store(user *User) error
	Update(id int, user *User) error
	FindAll() ([]*User, error)
	Delete(id int) error
}

func InitDB(dsn string) (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{}, &Order{})
	return db, nil
}
