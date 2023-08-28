/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/venkatk-25/client-server-go/pkg/grpc"
	"github.com/venkatk-25/client-server-go/pkg/tcp"
	"go.uber.org/zap"
)

// clientCmd represents the server command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start client",
	Long:  `start client for given protocol`,
}

func initClients(cancelFunc context.CancelFunc, l *zap.SugaredLogger) {
	clientCmd.AddCommand(tcp.NewClientCmd(cancelFunc, l.Named("tcp")))
	clientCmd.AddCommand(grpc.NewClientCmd(cancelFunc, l.Named("grpc")))
	rootCmd.AddCommand(clientCmd)
}
