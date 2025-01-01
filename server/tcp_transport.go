package server

import (
	"net"
)

type TCPTransport struct {
	addr     string
	listener net.Listener
}

// 实现 net.Listener 接口
func (t *TCPTransport) Accept() (net.Conn, error) {
	return t.listener.Accept()
}

func (t *TCPTransport) Close() error {
	if t.listener != nil {
		return t.listener.Close()
	}
	return nil
}

func (t *TCPTransport) Addr() net.Addr {
	if t.listener != nil {
		return t.listener.Addr()
	}
	return nil
}

// 创建 TCP transport
func NewTCPTransport(addr string) *TCPTransport {
	return &TCPTransport{
		addr: addr,
	}
}

// Listen 初始化底层的 net.Listener
func (t *TCPTransport) Listen() error {
	listener, err := net.Listen("tcp", t.addr)
	if err != nil {
		return err
	}
	t.listener = listener
	return nil
}
