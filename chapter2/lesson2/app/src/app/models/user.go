package models

import (
	"app/src/app/models/concerns"
	"app/src/config"
)

type User struct {
	concerns.Base
	Name  string `gorm:"size:255" json:"name"`
	Email string `gorm:"size:255" json:"email"`
}

func GetAllUsers(users *[]User) {
	db := config.DbConnect()
	defer db.Close()

	db.Find(&users)
}

func GetUser(user *User, id string) {
	db := config.DbConnect()
	defer db.Close()

	db.First(&user, id)
}
