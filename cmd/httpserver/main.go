package main

import (
	"go-http-service/internal/api"
	"log"
)

func main() {
	router := api.NewRouter()
	log.Println("启动HTTP服务，监听端口8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("无法启动HTTP服务: %v", err)
	}
}
