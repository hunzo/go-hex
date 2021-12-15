package repository

import (
	"gorm.io/gorm"
)

type payloadRepositoryDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) payloadRepositoryDB {
	db.AutoMigrate(&Payload{}, &Attch{})
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
func (r payloadRepositoryDB) GetById(id int) ([]Payload, error) {
	ret := []Payload{}
	return ret, nil
}
func (r payloadRepositoryDB) DeleteById(id int) error {
	if err := r.db.Delete(&[]Payload{}, id).Error; err != nil {
		return err
	}
	return nil
}
