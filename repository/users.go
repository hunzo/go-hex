package repository

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Firstname         string
	LastName          string
	FilesAttachements []Files `gorm:"foreignKey:FileID"`
}

type Files struct {
	FileID    uint
	FileName  string
	FileType  string
	FielBytes []byte
}

type UsersRepository interface {
	GetUsers() ([]Users, error)
}
