package service

import "go-hex/repository"

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
