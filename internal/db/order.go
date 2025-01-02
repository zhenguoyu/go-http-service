package db

type Order struct {
	ID     string
	Amount float64
	Date   string
}

func GetOrdersFromDB(userID string, limit int) ([]Order, error) {
	// 模拟从数据库获取订单信息
	// 实际代码中，这里会有数据库查询逻辑
	orders := []Order{
		{ID: "1", Amount: 100.0, Date: "2023-01-01"},
		{ID: "2", Amount: 200.0, Date: "2023-01-02"},
		// 添加更多订单
	}
	return orders, nil
}
