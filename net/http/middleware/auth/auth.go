package auth

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/lyz-miao/account-service/api"
	"github.com/lyz-miao/go-common/encode"
	"github.com/lyz-miao/go-common/net/http"
)

type Config struct {
    AuthFunc AuthFunc
}

type auth struct {
	accountService pb.AccountClient
    authFunc AuthFunc
}

type AuthFunc func (token string) (*pb.AccessTokenData, error)

func New() *auth {
	return &auth{
		accountService: pb.NewGRPCClient(context.Background()),
	}
}

func NewWithConfig(c *Config) *auth {
	return &auth{
		accountService: pb.NewGRPCClient(context.Background()),
        authFunc: c.AuthFunc,
	}
}

func (a *auth) TokenAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		c := ctx.(*http.Context)
		tk := ""
		tk = c.Request().Header.Get("Authorization")
		if tk == "" {
			tk = c.Request().URL.Query().Get("Authorization")
		} else if tk == "" {
			return encode.JSON(c, encode.Unauthorized, nil)
		}

        var data *pb.AccessTokenData
        var err error

        if a.authFunc == nil{
            r, e := a.accountService.GetAccessTokenData(context.Background(), &pb.GetAccessTokenDataReq{Token: tk})
            data = r.Data
            err = e
        }else {
            r, e := a.authFunc(tk)
            err = e
            data = r
        }
        if err != nil {
            ctx.Logger().Error(err)
            return encode.JSON(c, encode.Unauthorized, nil)
        }
        if data == nil {
            return encode.JSON(c, encode.Unauthorized, nil)
        }

		c.SetAccessTokenData(data)

		return next(c)
	}
}
