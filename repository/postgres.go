package repository

import (
	"github.com/ArthurQR98/fiber-gorm/domain"
	"gorm.io/gorm"
)

type postgresRepository struct {
	conection *gorm.DB
}

func (r *postgresRepository) Find(id int) (*domain.User, error) {
	var user *domain.User
	r.conection.First(&user, id)
	return user, nil
}
func (r *postgresRepository) Store(user *domain.User) error {
	result := r.conection.Create(&user)
	return result.Error
}
func (r *postgresRepository) Update(id int, user *domain.User) error {
	// var userup *domain.User
	result := r.conection.Model(&user).Where("id = ?", id).Update("data", user.Data)
	return result.Error
}

func (r *postgresRepository) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	result := r.conection.Find(&users)
	return users, result.Error
}

func (r *postgresRepository) Delete(id int) error {
	var user *domain.User
	r.conection.Delete(&user, id)
	return nil
}

func NewPostgresRepository(db *gorm.DB) (domain.Repository, error) {
	repo := &postgresRepository{
		conection: db,
	}
	return repo, nil
}
