package repository

import (
	"gorm.io/gorm"
)

type usersRepositoryDB struct {
	gdb *gorm.DB
}

func NewUsers(db *gorm.DB) usersRepositoryDB {
	db.AutoMigrate(&Users{})
	return usersRepositoryDB{gdb: db}
}

func (db usersRepositoryDB) GetUsers() ([]Users, error) {
	users := []Users{}
	db.gdb.Find(&users)

	return users, nil
}

func (db usersRepositoryDB) GetUserById(id int) (*Users, error) {
	user := Users{}
	db.gdb.Find(&user, id)

	return &user, nil
}

func (db usersRepositoryDB) DeleteUserById(id int) error {
	user := Users{}
	if chk := db.gdb.First(&user, "id = ?", id); chk.Error != nil {
		return chk.Error
	}
	if gdb := db.gdb.Unscoped().Delete(&user, id); gdb.Error != nil {
		return gdb.Error
	}
	return nil
}

func (db usersRepositoryDB) CreateUser(user UserCreate) error {
	u := Users{
		Firstname: user.Firstname,
		LastName:  user.LastName,
	}
	db.gdb.Create(&u)

	return nil
}
