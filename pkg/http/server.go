package http

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	Logger     *zap.SugaredLogger
	Address    string
	TLSAddress string
}

func (s *Server) Serve(ctx context.Context) {
	http.HandleFunc("/", s.root)

	s.Logger.Info("starting HTTP server at ", s.Address)
	err := http.ListenAndServe(s.Address, nil)
	if err != nil {
		s.Logger.Error(err)
	}
}

func (s *Server) ServeTLS(ctx context.Context) {
	s.Logger.Info("starting HTTPS server at ", s.TLSAddress)
	err := http.ListenAndServeTLS(s.TLSAddress, "scripts/server-cert.pem", "scripts/server-key.pem", nil)
	if err != nil {
		s.Logger.Error(err)
	}
}

func (s *Server) root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}
