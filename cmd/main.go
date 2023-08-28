package main

import (
	"context"

	"go.uber.org/zap"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	initServers(ctx, sugar)
	initClients(cancelFunc, sugar)
	Execute()

	<-ctx.Done()
}
