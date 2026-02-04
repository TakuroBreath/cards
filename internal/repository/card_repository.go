package repository

import (
	"github.com/TakuroBreath/cards/internal/models"
	"gorm.io/gorm"
)

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) Create(card *models.Card) error {
	return r.db.Create(card).Error
}

func (r *CardRepository) GetAllByUser(userID uint) ([]models.Card, error) {
	var cards []models.Card
	if err := r.db.Where("user_id = ?", userID).Find(&cards).Error; err != nil {
		return nil, err
	}

	return cards, nil
}
