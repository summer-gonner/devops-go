package mappers

import "devops-go/basicdata/models/entity"

type AuthMapper interface {
	QueryUserByUsernameAndPassword(username string, password string) *entity.User
}
