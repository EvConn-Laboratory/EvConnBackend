package http

import (
	"context"
	"net/http"
	"time"

	"evconn/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	http   *http.Server
}

func NewServer(port string) *Server {
	router := gin.Default()

	return &Server{
		router: router,
		http: &http.Server{
			Addr:         ":" + port,
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

// Update the SetupMiddleware function
func (s *Server) SetupMiddleware() {
	s.router.Use(
		gin.Recovery(),
		middleware.CORS(),
		middleware.RequestID(),
		middleware.Logger(),
	)
}

// SetupRoutes method is defined in routes.go
