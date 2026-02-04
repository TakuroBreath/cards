package service

import (
	"github.com/TakuroBreath/cards/internal/models"
	"github.com/TakuroBreath/cards/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Username: req.Username,
		Password: HashPassword(req.Password),
	}
	
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	
	return user, nil
}

func HashPassword(p string) string {
	return p
}
