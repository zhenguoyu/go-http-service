package db

type User struct {
	ID    string
	Name  string
	Email string
}

func GetUserFromDB(userID string) (User, error) {
	// 模拟从数据库获取用户信息
	// 实际代码中，这里会有数据库查询逻辑
	return User{ID: userID, Name: "John Doe", Email: "john.doe@example.com"}, nil
}
