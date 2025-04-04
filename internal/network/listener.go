package network

import "github.com/rs/zerolog"

type Listener interface {
	StartListening(addr string) error
	StopListening() error
}

type listener struct {
	logger    zerolog.Logger
	closeChan chan struct{}
}

func (l listener) StartListening(addr string) error {
	//TODO implement me
	panic("implement me")
}

func (l listener) StopListening() error {
	//TODO implement me
	panic("implement me")
}

var _ Listener = (*listener)(nil)
