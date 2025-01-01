package server

import (
	"errors"
	"io"
	"net"
	"time"

	"go.bug.st/serial"
)

type SerialTransport struct {
	port     string
	baudRate int
	conn     serial.Port
	connChan chan net.Conn // 添加连接通道
}

func (s *SerialTransport) Accept() (net.Conn, error) {
	// 阻塞等待连接通道的数据
	conn, ok := <-s.connChan
	if !ok {
		return nil, errors.New("connection channel closed")
	}
	return conn, nil
}

// 在 Listen 中初始化
func (s *SerialTransport) Listen() error {
	mode := &serial.Mode{
		BaudRate: s.baudRate,
	}

	port, err := serial.Open(s.port, mode)
	if err != nil {
		return err
	}

	s.conn = port
	s.connChan = make(chan net.Conn, 1)

	// 立即放入一个连接
	s.connChan <- &serialConn{s.conn}

	return nil
}

// 创建串口 transport
func NewSerialTransport(port string, baudRate int) *SerialTransport {
	return &SerialTransport{
		port:     port,
		baudRate: baudRate,
	}
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
