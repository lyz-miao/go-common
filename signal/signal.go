package signal

import (
	"os"
	"os/signal"
)

type Signal struct {
	ch     chan os.Signal
	signal []os.Signal
}

func New(signal ...os.Signal) *Signal {
	return &Signal{
		ch:     make(chan os.Signal, 1),
		signal: signal,
	}
}

func (s *Signal) Start() {
	signal.Notify(s.ch, s.signal...)
	defer close(s.ch)
	switch <-s.ch {
	default:
		return
	}
}
