package mappers

type AuthMapper interface {
	QueryUserByUsernameAndPassword(username string, password string) *User
}
