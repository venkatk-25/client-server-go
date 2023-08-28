package grpc

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/venkatk-25/client-server-go/pkg/grpc/pb"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	Logger     *zap.SugaredLogger
	Address    string
	TLSAddress string

	pb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	s.Logger.Info("GRPC server received Message: ", msg.Body)
	return &pb.Message{Body: "Hello!"}, nil
}

func (s *Server) Serve(context.Context) {
	s.Logger.Info("Starting grpc server at ", s.Address)
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		panic(err)
	}

	server := g.NewServer()
	pb.RegisterChatServiceServer(server, s)
	if err := server.Serve(listener); err != nil {
		s.Logger.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) ServeTLS(context.Context) {
	s.Logger.Info("Starting grpc server at ", s.TLSAddress)
	listener, err := net.Listen("tcp", s.TLSAddress)
	if err != nil {
		panic(err)
	}

	server := g.NewServer(g.Creds(credentials.NewTLS(&tls.Config{})))
	pb.RegisterChatServiceServer(server, s)
	if err := server.Serve(listener); err != nil {
		s.Logger.Fatalf("failed to serve: %v", err)
	}
}
