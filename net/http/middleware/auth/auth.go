package auth

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"library/database/redis"
	"library/encode"
)

type Config struct {
	JWTSecret string
	JWTClaims jwt.Claims
	// CacheKey Default access_token
	CacheKey string
	Redis    *redis.Config
}

type Auth struct {
	config *Config
	redis  *redis2.Client
}

func mereConfig(c *Config)  {
    if c.CacheKey == "" {
        c.CacheKey = "access_token"
    }
}

func New(c *Config) *Auth {
    mereConfig(c)

	return &Auth{
		config: c,
		redis:  redis.New(c.Redis),
	}
}

func (a *Auth) JWTAuth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(a.config.JWTSecret),
		Claims:     a.config.JWTClaims,
		ContextKey: "uid",
		ErrorHandlerWithContext: func(err error, ctx echo.Context) error {
			return encode.JSON(ctx, encode.Unauthorized, nil)
		},
	})
}

func (a *Auth) TokenAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tk := ""
		tk = ctx.Request().Header.Get("Authorization")
		if tk == "" {
			tk = ctx.Request().URL.Query().Get("Authorization")
		} else if tk == "" {
			return encode.JSON(ctx, encode.Unauthorized, nil)
		}

		k := fmt.Sprintf("%v:%v", a.config.CacheKey, tk)
		b, err := a.redis.Get(context.Background(), k).Uint64()
		if err != nil || b == 0 {
			return encode.JSON(ctx, encode.Unauthorized, nil)
		}

		ctx.Set("uid", b)

		return next(ctx)
	}
}
