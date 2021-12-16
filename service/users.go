package service

type UsersResponse struct {
	Firstname string
	LastName  string
	// FilesAttachements []Files
}

type Files struct {
	FileID    uint
	FileName  string
	FileType  string
	FielBytes []byte
}

type UserCreate struct {
	Firstname string
	LastName  string
	// FilesAttachements []Files
}

type UserSevice interface {
	GetUsers() ([]UsersResponse, error)
	CreateUser(UserCreate) error
}
