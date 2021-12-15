package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payload struct {
	gorm.Model
	Uid         uuid.UUID `gorm:"type:uuid"`
	FileName    string
	FileType    string
	FileContent []byte
	FileList    []byte
	AttachFile  []Attch `gorm:"foreignKey:FileID"`
}

type Attch struct {
	FileID    uuid.UUID `gorm:"type:uuid"`
	FileBytes []byte
}

type payloadRepository interface {
	GetData() ([]Payload, error)
	GetById(int) (Payload, error)
	CreateData(Payload) error
	DeleteById(int) error
}
