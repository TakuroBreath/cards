package handlers

import (
	"net/http"
	"strconv"

	"github.com/TakuroBreath/cards/internal/models"
	"github.com/TakuroBreath/cards/internal/service"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	service *service.CardService
}

func NewCardHandler(service *service.CardService) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) RegisterRoutes(r *gin.Engine) {
	cards := r.Group("/api/cards/:user_id")
	{
		cards.POST("", h.Create)
		cards.GET("", h.GetAll)
	}
}

func (h *CardHandler) Create(c *gin.Context) {
	uid, ok := parseUserID(c)
	if !ok {
		return
	}

	var req models.CreateCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card, err := h.service.CreateCard(uid, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, card)
}

func (h *CardHandler) GetAll(c *gin.Context) {
	uid, ok := parseUserID(c)
	if !ok {
		return
	}
	
	cards, err := h.service.GetAllCards(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	c.JSON(http.StatusOK, cards)
}

func parseUserID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return 0, false
	}

	return uint(id), true
}
