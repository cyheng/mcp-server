package main

import (
	"mcp-server/easytcp"
	"mcp-server/logger"
	"mcp-server/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	serialTransport := server.NewSerialTransport("COM1", 9600)

	// 创建服务器实例

	srv := server.NewServer(serialTransport, &easytcp.ServerOption{})
	srv.Srv.Use(logger.RecoverMiddleware(logger.Ins()), logger.LogMiddleware)
	easytcp.SetLogger(logger.Ins())
	// register a route

	// 启动服务器
	logger.Ins().Infoln("start server...")
	if err := srv.Start(); err != nil {
		logger.Ins().Fatalf("Failed to start server: %v", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	if err := srv.Stop(); err != nil {
		logger.Ins().Errorf("server stopped err: %s", err)
	}
	time.Sleep(time.Second * 3)
}
