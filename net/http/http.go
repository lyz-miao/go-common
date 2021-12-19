package http

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	log2 "github.com/labstack/gommon/log"
	"github.com/lyz-miao/go-common/encode"
	"log"
	"net"
	"net/http"
)

type ServerConfig struct {
	DEBUG bool
	// Addr is grpc listen addr,default value is 0.0.0.0:8000
	Addr string
}

type Server struct {
	config *ServerConfig
	server *echo.Echo
}

var _defaultConfig = &ServerConfig{
	Addr: "0.0.0.0:9000",
}

func mereConfig(config *ServerConfig) {
	if config == nil {
		config = _defaultConfig
	}

	if config.Addr == "" {
		config.Addr = "0.0.0.0:8000"
	}
}

func NewServer(c *ServerConfig) *Server {
	mereConfig(c)

	app := echo.New()
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{c}
			return next(cc)
		}
	})
	app.Debug = c.DEBUG
	app.HideBanner = true
	app.Validator = &CustomValidator{validator: validator.New()}
	if c.DEBUG {
		app.Logger.SetLevel(log2.DEBUG)
	} else {
		app.Logger.SetLevel(log2.INFO)
	}
	app.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if c.DEBUG {
			err = encode.JSON(ctx, encode.InternalServerError, err.Error())
		} else {
			err = encode.JSON(ctx, encode.InternalServerError, nil)
		}
		if err != nil {
			ctx.Logger().Error(err)
		}

		ctx.Logger().Error(err)
	}

	l, err := net.Listen("tcp", c.Addr)
	if err != nil {
		panic(err)
	}
	app.Listener = l

	return &Server{
		config: c,
		server: app,
	}
}

func (s *Server) Start() error {
	defer log.Printf("http server closed\n")
	err := s.server.Start("")
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Server() *echo.Echo {
	return s.server
}

func (s *Server) Close() {
	s.server.Close()
}
