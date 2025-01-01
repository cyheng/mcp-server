package server

import (
	"io"
	"net"
	"time"

	"go.bug.st/serial"
)

type SerialTransport struct {
	port     string
	baudRate int
	conn     io.ReadWriteCloser
}

// 创建串口 transport
func NewSerialTransport(port string, baudRate int) *SerialTransport {
	return &SerialTransport{
		port:     port,
		baudRate: baudRate,
	}
}

func (s *SerialTransport) Listen() error {
	mode := &serial.Mode{
		BaudRate: s.baudRate,
	}

	port, err := serial.Open(s.port, mode)
	if err != nil {
		return err
	}

	s.conn = port
	return nil
}

func (s *SerialTransport) Accept() (net.Conn, error) {
	// 串口不需要 Accept,直接返回一个包装了串口连接的 net.Conn
	return &serialConn{s.conn}, nil
}

func (s *SerialTransport) Close() error {
	if s.conn != nil {
		return s.conn.(io.Closer).Close()
	}
	return nil
}

// 包装串口连接为 net.Conn 接口
type serialConn struct {
	io.ReadWriteCloser
}

func (s *serialConn) LocalAddr() net.Addr                { return &addr{"serial", s.RemoteAddr().String()} }
func (s *serialConn) RemoteAddr() net.Addr               { return &addr{"serial", "remote"} }
func (s *serialConn) SetDeadline(_ time.Time) error      { return nil }
func (s *serialConn) SetReadDeadline(_ time.Time) error  { return nil }
func (s *serialConn) SetWriteDeadline(_ time.Time) error { return nil }

type addr struct {
	network string
	address string
}

func (a *addr) Network() string { return a.network }
func (a *addr) String() string  { return a.address }

// 添加 Addr 方法实现 net.Listener 接口
func (s *SerialTransport) Addr() net.Addr {
	return &addr{
		network: "serial",
		address: s.port,
	}
}
