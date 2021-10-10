package database

import (
	"github.com/Theta-Dev/Talon/src/try"
	"gorm.io/gorm"
)

func (db Database) UserAdd(user *User) (caught error) {
	defer try.Returnf(&caught, "error adding user")

	if user.Permission == nil {
		user.Permission = &Permission{}
	}

	try.Check(db.orm.Transaction(func(tx *gorm.DB) (tcaught error) {
		defer try.Return(&tcaught)

		try.ORM(tx.Create(user.Permission))
		try.ORM(tx.Create(user))
		return
	}))
	return
}

func (db Database) UserByID(id uint) (user *User, caught error) {
	defer try.Returnf(&caught, "error getting user")

	var u User
	if try.ORMIsEmpty(db.orm.Joins("Permission").First(&u, id)) {
		return nil, nil
	}
	return &u, nil
}

func (db Database) UserByName(name string) (user *User, caught error) {
	defer try.Returnf(&caught, "error getting user")

	var u User
	if try.ORMIsEmpty(db.orm.Joins("Permission").Where("name = ?", name).First(&u)) {
		return nil, nil
	}
	return &u, nil
}
