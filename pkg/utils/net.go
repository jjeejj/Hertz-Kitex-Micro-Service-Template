package utils

import (
	"net"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
)

// GetFreePort get a free port.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.FreePortAddress)
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP(consts.TCP, addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}
