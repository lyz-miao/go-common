package limit

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lyz-miao/go-common/rate/limit"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type Limit struct {
	limiter *limit.Limiter
}

func New(r rate.Limit, b int) *Limit {
	return &Limit{
		limiter: limit.New(r, b),
	}
}

func (l *Limit) LimiterWithEcho(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userAgent := ctx.Request().UserAgent()
		ip := ctx.RealIP()
		hash := fmt.Sprintf("%v_%v\n", userAgent, ip)
		m := l.limiter.Get(hash)

		if !m.AllowN(time.Now(), 3) {
			return ctx.HTML(http.StatusTooManyRequests, "<h1>Too many requests</h1>")
		}

		return next(ctx)
	}
}
