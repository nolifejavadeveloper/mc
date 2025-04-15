package network

import (
	"github.com/rs/zerolog"
	"net"
)

type Listener struct {
	logger    zerolog.Logger
	closeChan chan struct{}
}

func NewListener(logger zerolog.Logger) *Listener {
	return &Listener{
		logger:    logger,
		closeChan: make(chan struct{}),
	}
}

func (l *Listener) StartListening(addr string) {
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

				go l.handleConnection(conn)
			}
		}
	}()
}

func (l *Listener) handleConnection(conn net.Conn) {

}

func (l *Listener) StopListening() {
	l.closeChan <- struct{}{}
}
