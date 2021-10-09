package database

import (
	"github.com/Theta-Dev/Talon/src/try"
	"github.com/Theta-Dev/Talon/src/util"
	"gorm.io/gorm"
)

func (db Database) AddUser(name, password string, perm *Permission) (
	user *User, caught error) {
	defer try.Returnf(&caught, "error adding user")

	pwdHash := try.String(util.HashPassword(password))

	try.Check(db.orm.Transaction(func(tx *gorm.DB) (tcaught error) {
		defer try.Return(&tcaught)

		try.ORM(tx.Create(perm))

		user = &User{
			Name:         name,
			PasswordHash: pwdHash,
			Permission:   perm,
		}
		try.ORM(tx.Create(user))
		return
	}))
	return
}

func (db Database) GetUserByID(id uint) (user *User, caught error) {
	defer try.Returnf(&caught, "error getting user")

	var u User
	try.ORM(db.orm.First(&u, id))
	return &u, nil
}

func (db Database) GetUserByName(name string) (user *User, caught error) {
	defer try.Returnf(&caught, "error getting user")

	var u User
	if try.ORMIsEmpty(db.orm.Where("name = ?", name).First(&u)) {
		return nil, nil
	}
	return &u, nil
}

func (db Database) LoginUser(name, password string) (user *User, caught error) {
	defer try.Returnf(&caught, "error logging in")
	user = try.X(db.GetUserByName(name)).(*User)

	if user == nil || !util.CheckPasswordHash(password, user.PasswordHash) {
		return nil, util.ErrWrongUserPass
	}

	return
}

func (db Database) UpdateUserPassword(id uint, oldPwd, newPwd string) (caught error) {
	defer try.Returnf(&caught, "error changing password")

	user := try.X(db.GetUserByID(id)).(*User)

	if !util.CheckPasswordHash(oldPwd, user.PasswordHash) {
		panic(util.ErrWrongUserPass)
	}

	user.PasswordHash = try.String(util.HashPassword(newPwd))
	try.ORM(db.orm.Save(user))

	return
}
