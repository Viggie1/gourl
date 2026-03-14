package server

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func New() *Server {
	r := gin.Default()

	s := &Server{
		router: r,
	}

	s.registerRoutes()
	return s
}

func (s *Server) Run() error {
	return s.router.Run(":8080")
}
