package service

type UsersResponse struct {
	Firstname string
	LastName  string
}

type UserCreate struct {
	Firstname string
	LastName  string
}

type UserSevice interface {
	GetUsers() ([]UsersResponse, error)
	CreateUser(UserCreate) error
	DeleteUserById(int) error
}
