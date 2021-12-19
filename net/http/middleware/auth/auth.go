package auth

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/lyz-miao/account-service/api"
	"github.com/lyz-miao/go-common/encode"
	"github.com/lyz-miao/go-common/net/http"
)

type Auth struct {
	accountService pb.AccountClient
}

func New() *Auth {
	return &Auth{
		accountService: pb.NewGRPCClient(context.Background()),
	}
}

func (a *Auth) TokenAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		c := ctx.(*http.Context)
		tk := ""
		tk = c.Request().Header.Get("Authorization")
		if tk == "" {
			tk = c.Request().URL.Query().Get("Authorization")
		} else if tk == "" {
			return encode.JSON(c, encode.Unauthorized, nil)
		}

		rs, err := a.accountService.GetAccessTokenData(context.Background(), &pb.GetAccessTokenDataReq{Token: tk})
		if err != nil {
			ctx.Logger().Error(err)
			return encode.JSON(c, encode.Unauthorized, nil)
		}
		if rs.Data == nil {
			return encode.JSON(c, encode.Unauthorized, nil)
		}

		c.SetAccessTokenData(rs.Data)

		return next(c)
	}
}
