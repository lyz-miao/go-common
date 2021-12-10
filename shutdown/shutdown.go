package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

type shutdown struct {
	ch chan os.Signal
}

func New() *shutdown {
	return &shutdown{
		ch: make(chan os.Signal, 1),
	}
}

func (s *shutdown) Start() {
	signal.Notify(s.ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer close(s.ch)
	switch <-s.ch {
	default:
		return
	}
}

func (s *shutdown) Stop() {
	s.ch <- syscall.SIGINT
}
