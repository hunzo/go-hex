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
	FileContent string
}

type payloadRepository interface {
	GetData() ([]Payload, error)
	CreateData(Payload) error
}
