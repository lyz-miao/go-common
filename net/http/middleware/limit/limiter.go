package limit

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "library/rate/limit"
    "library/util"
    "net/http"
    "time"
)

type Limit struct {
    limiter *limit.Limiter
}

func New() *Limit {
    return &Limit{
        limiter: limit.New(20, 120),
    }
}

func (l *Limit) LimiterWithEcho(next echo.HandlerFunc) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        userAgent := ctx.Request().UserAgent()
        ip := ctx.RealIP()
        hash := util.SumMd5([]byte(fmt.Sprintf("%v\n%v\n", userAgent, ip)))
        m := l.limiter.Get(hash)

        if !m.AllowN(time.Now(), 3) {
            return ctx.HTML(http.StatusTooManyRequests, "<h1>Too many requests</h1>")
        }

        return next(ctx)
    }
}
