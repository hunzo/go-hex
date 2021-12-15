package repository

import (
	"gorm.io/gorm"
)

type payloadRepositoryDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) payloadRepositoryDB {
	db.AutoMigrate(&Payload{})
	return payloadRepositoryDB{db: db}
}

func (r payloadRepositoryDB) GetData() ([]Payload, error) {
	p := []Payload{}
	r.db.Find(&p)
	return p, nil
}

func (r *payloadRepositoryDB) CreateData(req Payload) error {
	r.db.Create(&req)
	return nil
}
