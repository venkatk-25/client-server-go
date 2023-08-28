package grpc

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewServer(logger *zap.SugaredLogger) *Server {
	return &Server{Logger: logger, Address: "localhost:8080", TLSAddress: "localhost:8081"}
}

func NewClient(logger *zap.SugaredLogger) *Client {

	return &Client{Logger: logger, Address: "localhost:8080"}
}

func NewClientCmd(cancelFunc context.CancelFunc, logger *zap.SugaredLogger) *cobra.Command {
	client := NewClient(logger)
	clientCmd := &cobra.Command{
		Use:   "grpc",
		Short: "Start  GRPC client",
		Long:  `GRPC client which replies with a message`,
		Run: func(cmd *cobra.Command, args []string) {
			client.Logger.Info("starting GRPC client")

			client.Connect(cancelFunc)
		},
	}
	return clientCmd
}
