package mappers

type AuthMapper interface {
	QueryUserByUsernameAndPassword(username string) *User
}
