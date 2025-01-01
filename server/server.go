package server

import (
	"mcp-server/easytcp"
)

type Server struct {
	srv       *easytcp.Server
	transport Transport
}

// NewServer 创建一个新的服务器实例
func NewServer(transport Transport) *Server {
	return &Server{
		transport: transport,
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	// 先启动传输层
	if err := s.transport.Listen(); err != nil {
		return err
	}

	// 创建 TCP 服务器
	s.srv = easytcp.NewServer(&easytcp.ServerOption{})

	// 开始接受连接
	go func() {
		for {
			conn, err := s.transport.Accept()
			if err != nil {
				// TODO: 处理错误，可以添加日志
				continue
			}
			s.srv.HandleConn(conn)
		}
	}()

	return nil
}

// Stop 停止服务器
func (s *Server) Stop() error {
	if s.srv != nil {
		s.srv.Stop()
	}
	return s.transport.Close()
}