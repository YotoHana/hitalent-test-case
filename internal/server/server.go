package server

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	mux *http.ServeMux
}

func (s *Server) Start() error {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}

func NewServer(cfg *Config) *Server {
	mux := http.NewServeMux()

	return &Server{
		mux: mux,
		httpServer: &http.Server{
			Addr: cfg.host + cfg.port,
			Handler: mux,
			ReadTimeout: cfg.readTimeout,
			WriteTimeout: cfg.writeTimeout,
			IdleTimeout: cfg.idleTimeout,
		},
	}
}