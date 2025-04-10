package network

import (
	"github.com/rs/zerolog"
	"net"
)

type MinecraftConn struct {
	logger zerolog.Logger
	conn   net.Conn
}

func makeMinecraftConn(logger zerolog.Logger, conn net.Conn) *MinecraftConn {
	return &MinecraftConn{
		logger: logger.With().Str("addr", conn.RemoteAddr().String()).Logger(),
		conn:   conn,
	}
}

func (mc *MinecraftConn) Read() {
	
}
