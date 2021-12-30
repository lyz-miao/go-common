package http

import (
	"github.com/labstack/echo/v4"
	"github.com/lyz-miao/account-service/api"
)

type Context struct {
	echo.Context
}

const _accessTokenStoreKey = "_ACCESS_TOKEN_DATA_"

func (c *Context) SetAccessTokenData(data *pb.AccessTokenData) {
	c.Set(_accessTokenStoreKey, data)
}

func (c *Context) GetAccessTokenData() *pb.AccessTokenData {
	v, ok := c.Get(_accessTokenStoreKey).(*pb.AccessTokenData)
	if ok && v != nil {
		return v
	}
	return nil
}
