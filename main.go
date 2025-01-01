package main

import (
	"log"
	"mcp-server/server"
)

func main() {
	// 创建串口传输，设置串口名称和波特率
	// Windows 上串口名称类似 "COM1"
	// Linux 上串口名称类似 "/dev/ttyUSB0"
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
