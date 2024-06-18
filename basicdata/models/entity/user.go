package entity

type User struct {
	Id       int64  `xorm:"'id' pk autoincr"`
	Username string `xorm:"'username' varchar(100)"`
	Password string `xorm:"'password' varchar(100)"`
}

func (User) TableName() string {
	return "sys_user"
}
