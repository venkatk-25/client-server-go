package http

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewServer(logger *zap.SugaredLogger) *Server {
	return &Server{Logger: logger, Address: "localhost:8000", TLSAddress: "localhost:8443"}
}

func NewServerCmd(ctx context.Context, logger *zap.SugaredLogger) *cobra.Command {
	server := NewServer(logger)
	serverCmd := &cobra.Command{
		Use:   "http",
		Short: "Start HTTP server",
		Long:  `Http Server which replies with a message`,
		Run: func(cmd *cobra.Command, args []string) {
			server.Logger.Info("starting HTTP server")
			go server.ServeTLS(ctx)

			// go server.Serve(ctx)
		},
	}
	return serverCmd
}
