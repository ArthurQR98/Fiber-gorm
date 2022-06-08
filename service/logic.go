package service

import "github.com/ArthurQR98/fiber-gorm/domain"

type service struct {
	userRepo domain.Repository
}

// NewUserService a service struct that attaches to a repository via the Repository interface
func NewUserService(userRepo domain.Repository) *service {
	return &service{userRepo: userRepo}
}

func (s *service) Find(id int) (*domain.User, error) {
	return s.userRepo.Find(id)
}
func (s *service) Store(user *domain.User) error {
	return s.userRepo.Store(user)
}
func (s *service) Update(id int, user *domain.User) error {
	return s.userRepo.Update(id, user)
}

func (s *service) FindAll() ([]*domain.User, error) {
	return s.userRepo.FindAll()
}

func (s *service) Delete(id int) error {
	return s.userRepo.Delete(id)
}
