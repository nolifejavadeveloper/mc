package network

import (
	"github.com/rs/zerolog"
	"net"
)

type MinecraftConn interface {
}

type minecraftConn struct {
	logger zerolog.Logger
	conn   net.Conn
}

func MakeMinecraftConn(logger zerolog.Logger, conn net.Conn) MinecraftConn {
	return &minecraftConn{
		logger: logger,
		conn:   conn,
	}
}
