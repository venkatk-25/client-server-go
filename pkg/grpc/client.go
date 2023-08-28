package grpc

import (
	"context"

	"github.com/venkatk-25/client-server-go/pkg/grpc/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Logger  *zap.SugaredLogger
	Address string
}

func (c *Client) Connect(cancelFunc context.CancelFunc) {
	defer cancelFunc()
	conn, err := grpc.Dial(c.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.Logger.Fatal("GRPC client connection failed err: ", err)
	}
	client := pb.NewChatServiceClient(conn)
	msg, err := client.SayHello(context.Background(), &pb.Message{Body: "World"})
	if err != nil {
		c.Logger.Fatal("GRPC client connection failed err: ", err)
	}
	c.Logger.Info("GRPC client received Message: ", msg.Body)
}
