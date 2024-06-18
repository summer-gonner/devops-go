package services

type ServicesGroup struct {
	AuthService AuthService
	UserService UserService
}

var ServiceGroupApp = new(ServicesGroup)
