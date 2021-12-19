package ipx

import (
	"errors"
	"net"
)

func PublicIPWithInterface() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if n, ok := addr.(*net.IPNet); ok && !n.IP.IsLoopback() {
			if n.IP.To4() != nil {
				return n.IP.String(), nil
			}
		}
	}

	return "", nil
}

func PublicIPWithDial() (string, error) {
	conn, err := net.Dial("udp", "1.0.0.1:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	addr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return "", errors.New("can not find listen addr")
	}

	return addr.IP.String(), nil
}
