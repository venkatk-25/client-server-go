package tcp

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	Logger  *zap.SugaredLogger
	Address string
}

func (s *Server) Serve(ctx context.Context) {
	s.Logger.Info("Starting TCP server at ", s.Address)
	listen, err := net.Listen("tcp", s.Address)
	if err != nil {
		s.Logger.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			s.Logger.Fatal(err)
			os.Exit(1)
		}
		go s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	size, err := conn.Read(buffer)
	if err != nil {
		s.Logger.Fatal(err)
	}

	receivedStr := string(buffer[:size])
	s.Logger.Info("Received message: ", receivedStr)
	// write data to response
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", receivedStr, time)
	conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
}
