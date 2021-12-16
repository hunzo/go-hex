package service

type UsersResponse struct {
	Firstname         string
	LastName          string
	FilesAttachements []Files
}

type Files struct {
	FileID    uint
	FileName  string
	FileType  string
	FielBytes []byte
}

type UserSevice interface {
	GetUsers() ([]UsersResponse, error)
}
