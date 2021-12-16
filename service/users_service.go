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
	u, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	userResponses := []UsersResponse{}

	for _, user := range u {
		userRespone := UsersResponse{
			Firstname: user.Firstname,
			LastName:  user.LastName,
			// FilesAttachements: ,
		}

		userResponses = append(userResponses, userRespone)
	}

	return userResponses, nil
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
