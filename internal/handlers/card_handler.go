package handlers

import "github.com/TakuroBreath/cards/internal/service"

type CardHandler struct {
	service *service.CardService
}

func NewCardHandler(service *service.CardService) *CardHandler {
	return &CardHandler{service: service}
}
