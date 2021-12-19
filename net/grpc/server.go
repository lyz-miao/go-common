package grpc

import (
	"errors"
	googleGrpc "google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *ServerConfig
	server *googleGrpc.Server
	listen net.Listener
}

type ServerConfig struct {
	// Addr is grpc listen addr,default value is 0.0.0.0:9000
	Addr string
}

var _defaultConfig = &ServerConfig{
	Addr: "0.0.0.0:9000",
}

func mereConfig(config *ServerConfig) {
	if config == nil {
		config = _defaultConfig
	}

	if config.Addr == "" {
		config.Addr = "0.0.0.0:9000"
	}
}

func NewServer(c *ServerConfig) *Server {
	mereConfig(c)

	s := googleGrpc.NewServer()

	l, err := net.Listen("tcp", c.Addr)
	if err != nil {
		panic(err)
	}

	return &Server{
		config: c,
		server: s,
		listen: l,
	}
}

func (g *Server) Start() error {
	log.Printf("gRPC server running in %v", g.listen.Addr().String())
	err := g.server.Serve(g.listen)
	if err != nil {
		if !errors.Is(err, googleGrpc.ErrServerStopped) {
			return err
		}
	}
	return nil
}

func (g *Server) Close() {
	log.Printf("gRPC server closed")
	g.server.GracefulStop()
}

func (g *Server) Server() *googleGrpc.Server {
	return g.server
}

func (g *Server) Listen() net.Listener {
	return g.listen
}
