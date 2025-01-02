package api

import (
	"go-http-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUser(c *gin.Context) {
	userID := c.Param("id")

	// 调用service层的业务逻辑获取用户信息
	user, err := service.GetUserDetails(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 调用service层的业务逻辑获取用户的最近订单信息
	orders, err := service.GetRecentOrders(userID, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回用户信息和订单信息
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"orders": orders,
	})
}
