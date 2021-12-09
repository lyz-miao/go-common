package encode

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Content interface{} `json:"content,omitempty"`
}

type Code struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func JSON(ctx echo.Context, code Code, content interface{}) error {
	data := &response{
		Code:    code.Code,
		Message: code.Message,
		Content: content,
	}

	return ctx.JSON(http.StatusOK, data)
}
