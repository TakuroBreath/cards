package service

import (
	"errors"

	"github.com/TakuroBreath/cards/internal/models"
	"github.com/TakuroBreath/cards/internal/repository"
)

type CardService struct {
	repo     *repository.CardRepository
	userRepo *repository.UserRepository
}

func NewCardService(repo *repository.CardRepository, userRepo *repository.UserRepository) *CardService {
	return &CardService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *CardService) CreateCard(userID uint, req models.CreateCardRequest) (*models.Card, error) {
	if _, err := s.userRepo.GetById(userID); err != nil {
		return nil, errors.New("user not found")
	}

	card := &models.Card{
		UserID:   userID,
		Topic:    req.Topic,
		Question: req.Question,
		Answer:   req.Answer,
	}
	if err := s.repo.Create(card); err != nil {
		return nil, err
	}

	return card, nil
}

func (s *CardService) GetAllCards(userID uint) ([]models.Card, error) {
	return s.repo.GetAllByUser(userID)
}
