package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Viggie1/gourl/internal/database"
	"github.com/Viggie1/gourl/internal/handlers"
	"github.com/Viggie1/gourl/internal/repository"
)

type Server struct {
	router  *gin.Engine
	handler *handlers.URLHandler
}

func New() (*Server, error) {
	r := gin.Default()
	db, err := database.New()
	if err != nil {
		return nil, err
	}
	repo := repository.NewURLRepository(db)
	handler := handlers.NewURLHandler(repo)

	s := &Server{
		router:  r,
		handler: handler,
	}

	s.registerRoutes()
	return s, nil
}

func (s *Server) Run() error {
	return s.router.Run(":8080")
}
