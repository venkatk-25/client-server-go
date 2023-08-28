package tcp

import (
	"context"
	"net"
	"os"

	"go.uber.org/zap"
)

type Client struct {
	Logger  *zap.SugaredLogger
	Address string
}

func (c *Client) Connect(cancelFunc context.CancelFunc) {
	defer cancelFunc()
	tcpServer, err := net.ResolveTCPAddr("tcp", c.Address)
	if err != nil {
		c.Logger.Error("Resolve address failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		c.Logger.Error("Dial failed:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte("This is a message"))
	if err != nil {
		c.Logger.Error("Write data failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	size, err := conn.Read(received)
	if err != nil {
		c.Logger.Error("Read data failed:", err.Error())
		os.Exit(1)
	}
	c.Logger.Info("Received message:", string(received[:size]))
}
