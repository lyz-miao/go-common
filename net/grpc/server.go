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
	return g.server.Serve(g.listen)
}

func (g *Server) Close() {
	log.Printf("gRPC server closed")
	g.server.Stop()
}

func (g *Server) Server() *googleGrpc.Server {
	return g.server
}

// GetServerOutsideListenAddr
// 用于获取外部应用访问服务时的地址
func (g *Server) GetServerOutsideListenAddr() (*net.TCPAddr, error) {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	outsideAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return nil, errors.New("can not find listen addr")
	}

	n := &net.TCPAddr{
		IP:   outsideAddr.IP,
		Port: g.listen.Addr().(*net.TCPAddr).Port,
	}

	return n, nil
}
