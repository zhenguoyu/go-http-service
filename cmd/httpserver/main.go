package main

import (
	"go-http-service/internal/api"
	"go-http-service/internal/config"
	"go-http-service/internal/logger"
)

func main() {
	// 载入配置文件
	config.LoadConfig("config/config.yaml")

	// 初始化路由
	router := api.NewRouter()

	// 启动HTTP服务
	logger.InfoLogger("启动HTTP服务，监听端口8080...")
	if err := router.Run(":8080"); err != nil {
		logger.ErrorLogger("无法启动HTTP服务: %v", err)
	}
}
