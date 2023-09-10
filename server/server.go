package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		Handler:      handler,
	}
	return s.httpServer.ListenAndServe()
}
