package service

import (
	"errors"
	"go-http-service/internal/db"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUserDetails(userID string) (User, error) {
	// 这里进行业务逻辑处理，例如从数据库获取用户信息
	user, err := db.GetUserFromDB(userID)
	if err != nil {
		return User{}, errors.New("无法获取用户信息")
	}
	return user, nil
}
