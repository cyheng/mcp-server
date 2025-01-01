package main

import (
	"log"
	"mcp-server/server"
)

func main() {

	serialTransport := server.NewSerialTransport("/dev/ttyUSB0", 115200)

	// 创建服务器实例
	srv := server.NewServer(serialTransport)

	// 启动服务器
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// 等待信号以优雅退出
	waitForSignal()

	// 停止服务器
	if err := srv.Stop(); err != nil {
		log.Printf("Error stopping server: %v", err)
	}
}

func waitForSignal() {
	// 实现信号处理逻辑，等待 Ctrl+C 等信号
	select {}
}
