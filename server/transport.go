package server

import "net"

// Transport 接口继承 net.Listener 接口
type Transport interface {
	net.Listener
	Listen() error
}
