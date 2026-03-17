// Package handlers
package handlers

import (
	"net/http"
	"time"

	"github.com/Viggie1/gourl/internal/models"
	"github.com/Viggie1/gourl/internal/repository"
	"github.com/Viggie1/gourl/internal/utils"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

type CreateURLRequest struct {
	URL string `json:"url" binding:"required"`
}

type CreateURLResponse struct {
	ShortCode string `json:"shortCode"`
	ShortURL  string `json:"shortURL"`
}

func NewURLHandler(repo *repository.URLRepository) *URLHandler {
	return &URLHandler{repo: repo}
}

func (h *URLHandler) CreateShortURL(c *gin.Context) {
	var req CreateURLRequest

	// validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error 400 Bad Request"})
		return
	}

	// Check if URL exists
	existingURL, err := h.repo.CheckURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	if existingURL != nil {
		c.JSON(http.StatusOK, CreateURLResponse{
			ShortCode: existingURL.ShortCode,
			ShortURL:  "https://localhost:8080/" + existingURL.ShortCode,
		})
		return
	}

	// Generate short code
	shortCode := utils.GenerateShortcode()

	// Save to db with repo
	shortenedURL := &models.ShortenedURL{
		URL:       req.URL,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.repo.Create(shortenedURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create"})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, CreateURLResponse{
		ShortCode: shortCode,
		ShortURL:  "https://localhost:8080/" + shortCode,
	})
}

func (h *URLHandler) RetrieveOriginalURL(c *gin.Context) {
	id := c.Param("shortcode")

	shortenedURL, err := h.repo.Retrieve(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if shortenedURL == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Original URL not found"})
	}

	c.JSON(http.StatusOK, shortenedURL)
}
