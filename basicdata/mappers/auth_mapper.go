package mappers

import "devops-go/models/entity"

type AuthMapper interface {
	QueryUserByUsernameAndPassword(username string, password string) *entity.User
}
