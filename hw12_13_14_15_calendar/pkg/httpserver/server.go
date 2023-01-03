package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(host string, port int, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second, // todo implement in config
		WriteTimeout: 10 * time.Second, // todo implement in config
	}

	server := &Server{
		server: httpServer,
	}
	return server
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	os.Exit(1)
	return nil
}
