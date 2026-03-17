// Package server
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes() {
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"worked": "hello"})
	})

	s.router.POST("/api/v1/urls", s.handler.CreateShortURL)
	s.router.GET("/api/v1/shorten/:shortcode", s.handler.RetrieveOriginalURL)
	s.router.DELETE("/api/v1/shorten/:shortcode", s.handler.DeleteURL)
}
