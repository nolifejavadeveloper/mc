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

func makeMinecraftConn(logger zerolog.Logger, conn net.Conn) MinecraftConn {
	return &minecraftConn{
		logger: logger.With().Str("ip", conn.RemoteAddr().String()).Logger(),
		conn:   conn,
	}
}

func (mc *minecraftConn) Read() {

}

var _ MinecraftConn = (*minecraftConn)(nil)
