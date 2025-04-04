package network

import (
	"github.com/rs/zerolog"
	"net"
)

type Listener interface {
	StartListening(addr string)
	StopListening()
}

type listener struct {
	logger    zerolog.Logger
	closeChan chan struct{}
}

func (l listener) StartListening(addr string) {
	go func() {
		tcpListener, err := net.Listen("tcp", addr)
		if err != nil {
			l.logger.Panic().Err(err).Msg("Error starting TCP listener")
			return
		}

		for {
			select {
			case l.closeChan <- struct{}{}:
				err := tcpListener.Close()
				if err != nil {
					l.logger.Error().Err(err).Msg("Error closing listener")
					return
				}
			default:
				conn, err := tcpListener.Accept()
				if err != nil {
					l.logger.Error().Err(err).Msg("Error accepting TCP connection")
					continue
				}

				l.handleConnection(conn)
			}
		}
	}()
}

func (l listener) handleConnection(conn net.Conn) {

}

func (l listener) StopListening() {
	//TODO implement me
	panic("implement me")
}

var _ Listener = (*listener)(nil)
