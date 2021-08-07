package iface

import "net"

type IConnection interface {
	Start()

	Stop()

	GetTCPConnection() *net.TCPConn

	GetConnID() uint32

	GetRemoteAddr() net.Addr

	Send(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
