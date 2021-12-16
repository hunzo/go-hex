package repository

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Firstname string
	LastName  string
}

type UserCreate struct {
	Firstname string
	LastName  string
}

type UsersRepository interface {
	GetUsers() ([]Users, error)
	GetUserById(int) (*Users, error)
	DeleteUserById(int) error
	CreateUser(UserCreate) error
}
