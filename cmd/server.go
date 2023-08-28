/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/venkatk-25/client-server-go/pkg/grpc"
	httpServer "github.com/venkatk-25/client-server-go/pkg/http"
	"github.com/venkatk-25/client-server-go/pkg/tcp"
	"go.uber.org/zap"
)

func initServers(ctx context.Context, l *zap.SugaredLogger) {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start server",
		Long:  `Start servers for different protocols.`,
		Run:   runServers(ctx, l),
	}

	rootCmd.AddCommand(serverCmd)
}

func runServers(ctx context.Context, l *zap.SugaredLogger) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		go httpServer.NewServer(l.Named("http")).Serve(ctx)
		go httpServer.NewServer(l.Named("https")).ServeTLS(ctx)
		go tcp.NewServer(l.Named("tcp")).Serve(ctx)
		go grpc.NewServer(l.Named("grpc")).Serve(ctx)
		go grpc.NewServer(l.Named("grpc-tls")).ServeTLS(ctx)
	}
}
