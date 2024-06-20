package mappers

import (
	"devops-go/basicdata/common/global"
	"devops-go/basicdata/models/entity"
	"fmt"
	"log"
)

type AuthMapperSql struct {
}

// QueryUserByUsernameAndPassword 根据用户名和密码查询用户信息
func (ams AuthMapperSql) QueryUserByUsernameAndPassword(username string, password string) *entity.User {
	user := new(entity.User)
	has, err := global.DB.Where("username = ?", username).And("password=?", password).Get(user)
	if err != nil {
		fmt.Println("Error querying user:", err)
		return nil
	}
	if !has {
		fmt.Println("User not found")
		return nil
	}
	log.Printf("Found user: %s", user)
	return user
}
