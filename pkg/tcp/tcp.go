package tcp

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewServer(logger *zap.SugaredLogger) *Server {
	return &Server{Logger: logger, Address: "localhost:3636"}
}

func NewServerCmd(ctx context.Context, logger *zap.SugaredLogger) *cobra.Command {
	server := NewServer(logger)
	serverCmd := &cobra.Command{
		Use:   "tcp",
		Short: "Start TCP server",
		Long:  `TCP Server which replies with a message`,
		Run: func(cmd *cobra.Command, args []string) {
			server.Logger.Info("starting TCP server")

			go server.Serve(ctx)
		},
	}
	return serverCmd
}

func NewClient(logger *zap.SugaredLogger) *Client {
	return &Client{Logger: logger, Address: "localhost:3636"}
}

func NewClientCmd(cancelFunc context.CancelFunc, logger *zap.SugaredLogger) *cobra.Command {
	client := NewClient(logger)
	serverCmd := &cobra.Command{
		Use:   "tcp",
		Short: "Start TCP server",
		Long:  `TCP Server which replies with a message`,
		Run: func(cmd *cobra.Command, args []string) {
			client.Logger.Info("starting TCP client")

			client.Connect(cancelFunc)
		},
	}
	return serverCmd
}
