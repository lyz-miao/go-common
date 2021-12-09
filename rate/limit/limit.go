package limit

import (
    "golang.org/x/time/rate"
    "sync"
)

type limit interface {
    Add(key string) *rate.Limiter
    Get(key string) *rate.Limiter
}

type Limiter struct {
    keys map[string]*rate.Limiter
    mu   *sync.RWMutex
    r    rate.Limit
    b    int
}

func New(r rate.Limit, b int) *Limiter {
    return &Limiter{
        keys: map[string]*rate.Limiter{},
        mu:   &sync.RWMutex{},
        r:    r,
        b:    b,
    }
}

func (l *Limiter) Add(key string) *rate.Limiter {
    l.mu.Lock()
    defer l.mu.Unlock()
    limiter := rate.NewLimiter(l.r, l.b)
    l.keys[key] = limiter
    return limiter
}

func (l *Limiter) Get(key string) *rate.Limiter {
    l.mu.Lock()
    limiter, ok := l.keys[key]

    if !ok {
        l.mu.Unlock()
        return l.Add(key)
    }

    l.mu.Unlock()
    return limiter
}
