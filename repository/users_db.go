package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type usersRepositoryDB struct {
	gdb *gorm.DB
}

type userController struct {
	user usersRepositoryDB
}

func NewUsers(db *gorm.DB) usersRepositoryDB {
	db.AutoMigrate()
	return usersRepositoryDB{gdb: db}
}

func (db usersRepositoryDB) GetUsers() ([]Users, error) {
	users := []Users{}
	db.gdb.Find(&users)
	return users, nil
}

func (db *userController) CreateUser(user UserCreate) error {

	u := UserCreate{
		Firstname: user.Firstname,
		LastName:  user.LastName,
		// FilesAttachements: [],
	}

	fmt.Printf("repository: %v\n", u)

	err := db.user.gdb.Create(&UserCreate{
		Firstname: "testsetset",
		LastName:  "Last ndfnsfmnsmfnsdf",
	}).Error

	if err != nil {
		panic(err)
	}

	return nil
}
