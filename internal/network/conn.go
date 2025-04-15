package network

import (
	"github.com/rs/zerolog"
	"io"
	"mc-server/pkg/buffer"
	"net"
)

const BufferLength = 4096

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

func (c *MinecraftConn) Read() {
	buf := buffer.MakeBuffer(make([]byte, 0, BufferLength))
	_, err := c.conn.Read(buf.Data)
	if err != nil {
		if err == io.EOF {
			c.logger.Debug().Msg("EOF: Connection closed")
			c.Close()
		}
		c.logger.Warn().Err(err).Msg("Error reading from connection")
	}

	packetLen, err := buf.ReadVarInt()
	if err != nil {
		c.logger.Warn().Err(err).Msg("Error reading packet length")
		return
	}

	if packetLen > BufferLength {
		targetLength := packetLen - BufferLength
		newData := make([]byte, 0, targetLength)

		for len(newData) != int(targetLength) {
			_, err = c.conn.Read(newData)
			if err != nil {
				if err == io.EOF {
					c.logger.Debug().Msg("EOF: Connection closed")
					c.Close()
					return
				}
				c.logger.Warn().Err(err).Msg("Error reading from connection")
			}
		}

		buf.Data = append(buf.Data, newData...)
	}

	// TODO: add handling for multiple packets in one read
}

func (c *MinecraftConn) Close() {
	err := c.conn.Close()
	if err != nil {
		c.logger.Warn().Err(err).Msg("Error closing connection")
	}
}
