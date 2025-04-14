package network

import (
	"github.com/rs/zerolog"
	"net"
)

const BufferLength = 4096;

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
	buffer := makeBuffer(make([]byte, 0, BufferLength))
	_, err := c.conn.Read(buffer.data)
	if err != nil {
		c.logger.Warn().Err(err).Msg("Error reading from connection")
	}


	packetLen, err := buffer.ReadVarInt()
	if err != nil {
		c.logger.Warn().Err(err).Msg("Error reading packet length")
		return;
	}

	if packetLen > BufferLength {
		targetLength := packetLen - BufferLength
		newData := make([]byte, 0, targetLength)

		for len(newData) != int(targetLength) {
			c.conn.Read(newData)
		}

		buffer.data = append(buffer.data, newData...)
	}



	// TODO: add handling for multiple packets in one read


}
