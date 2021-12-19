package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

type shutdown struct {
	ch   chan os.Signal
	stop chan struct{}
	f    func()
}

func New(f func()) *shutdown {
	s := &shutdown{
		ch:   make(chan os.Signal, 1),
		stop: make(chan struct{}),
		f:    f,
	}

	go func() {
		for {
			select {
			case <-s.stop:
				s.f()
			}
		}
	}()

	return s
}

func (s *shutdown) watch() {
	defer close(s.stop)
	switch <-s.stop {
	default:
		if s.f != nil {
			s.f()
		}
		return
	}
}

func (s *shutdown) Start() {
	signal.Notify(s.ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer close(s.ch)
	switch <-s.ch {
	default:
		if s.f != nil {
			s.f()
		}
		return
	}
}

func (s *shutdown) Stop() {
	s.ch <- syscall.SIGINT
}
