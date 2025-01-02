package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func HandleRequest(c *gin.Context) {
    // 处理特定的API请求
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello, World!",
    })
}