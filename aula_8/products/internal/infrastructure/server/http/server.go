package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() {
	go func() {
		log.Printf("Product API running on %s!\n", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("ERROR: on ListenAndServe: %q\n", err)
		}
	}()
}

func (s *Server) Shutdown() {
	log.Println("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("Could not shutdown in 60s: %q\n", err)
		return
	}
	log.Println("Server gracefully stopped")
}
