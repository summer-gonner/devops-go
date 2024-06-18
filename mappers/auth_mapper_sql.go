package mappers

import (
	"devops-go/common/global"
	"fmt"
	"log"
)

type AuthMapperSql struct {
}
type User struct {
	Id       int64  `xorm:"'id' pk autoincr"`
	Username string `xorm:"'username' varchar(100)"`
	Password string `xorm:"'password' varchar(100)"`
}

func (User) TableName() string {
	return "sys_user"
}

// QueryUserByUsernameAndPassword 根据用户名和密码查询用户信息
func (ams AuthMapperSql) QueryUserByUsernameAndPassword(username string) *User {
	user := new(User)
	_, err := global.DB.Where("username = ?", username).Get(user)
	if err != nil {
		fmt.Println("Error querying user:", err)
		return nil
	}
	log.Printf("user: %s", user)
	return user
}
