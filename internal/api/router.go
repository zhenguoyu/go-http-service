package api

import (
	"go-http-service/internal/log"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.LoggerWithWriter(log.InfoLogger.Writer()))
	router.GET("/user/:id", HandleUser)
	return router
}
