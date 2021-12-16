package repository

import "gorm.io/gorm"

type usersRepositoryDB struct {
	gdb *gorm.DB
}

func NewUsers(db *gorm.DB) usersRepositoryDB {
	db.AutoMigrate(&Users{}, &Files{})
	return usersRepositoryDB{gdb: db}
}

func (db usersRepositoryDB) GetUsers() ([]Users, error) {
	users := []Users{}
	db.gdb.Find(&users)
	return users, nil
}
