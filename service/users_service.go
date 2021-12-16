package service

import (
	"fmt"
	"go-hex/repository"
)

type userSevice struct {
	userRepo repository.UsersRepository
}

func NewUserService(ur repository.UsersRepository) userSevice {
	return userSevice{
		userRepo: ur,
	}
}

func (s userSevice) GetUsers() ([]UsersResponse, error) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	userResponses := []UsersResponse{}

	for _, user := range users {
		userRespone := UsersResponse{
			Firstname: user.Firstname,
			LastName:  user.LastName,
		}

		userResponses = append(userResponses, userRespone)
	}

	return userResponses, nil
}

func (s userSevice) DeleteUserById(id int) error {
	if err := s.userRepo.DeleteUserById(id); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s userSevice) CreateUser(user UserCreate) error {

	payload := repository.UserCreate{
		Firstname: user.Firstname,
		LastName:  user.LastName,
	}

	fmt.Printf("services: %v\n", payload)

	if err := s.userRepo.CreateUser(payload); err != nil {
		return err
	}

	return nil
}
