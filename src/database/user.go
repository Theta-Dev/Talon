package database

import (
	"github.com/Theta-Dev/Talon/src/try"
	"github.com/Theta-Dev/Talon/src/util"
)

func (db Database) AddUser(name, password string, perm Permission) (
	user User, tryErr error) {
	defer try.Returnf(&tryErr, "error adding user")

	pwdHash := try.String(util.HashPassword(password))

	user = User{
		Name:         name,
		PasswordHash: pwdHash,
		Permission:   perm,
	}
	db.orm.Create(&user)
	return
}
