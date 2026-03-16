// Package server
package server

import (
	"net/http"
	"time"

	"github.com/Viggie1/gourl/internal/handlers"
	"github.com/Viggie1/gourl/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes() {
	urlList := [4]models.ShortenedURL{
		{ID: 1, URL: "hello.com", ShortCode: "hello", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, URL: "hello.net", ShortCode: "net", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, URL: "hello.go", ShortCode: "go", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 4, URL: "hello.xyz", ShortCode: "xyz", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	s.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, urlList)
	})
	s.router.POST("/urls", handlers.CreateShortURL)
}
