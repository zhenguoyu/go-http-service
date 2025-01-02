package service

import (
	"errors"
	"go-http-service/internal/db"
)

type Order struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
}

func GetRecentOrders(userID string, limit int) ([]Order, error) {
	// 这里进行业务逻辑处理，例如从数据库获取订单信息
	orders, err := db.GetOrdersFromDB(userID, limit)
	if err != nil {
		return nil, errors.New("无法获取订单信息")
	}
	return orders, nil
}
