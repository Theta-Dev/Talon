package database

import (
	"fmt"

	"github.com/Theta-Dev/Talon/src/try"
	"gorm.io/gorm"
)

func (db *Database) UserAdd(user *User) (caught error) {
	defer try.Returnf(&caught, "error adding user")

	user.ID = 0
	try.Check(user.check(db))
	user.Permission.ID = 0

	try.Check(db.orm.Transaction(func(tx *gorm.DB) (tcaught error) {
		defer try.Return(&tcaught)

		try.ORM(tx.Create(user.Permission))
		try.ORM(tx.Create(user))
		return
	}))
	return
}

func (db *Database) UserUpdate(user *User) (caught error) {
	defer try.Returnf(&caught, "error updating user %d", user.ID)

	try.Check(user.check(db))
	try.ORM(db.orm.Save(user))
	return
}

func (db *Database) UserByID(id uint) (user *User, caught error) {
	defer try.Returnf(&caught, "error getting user %d", id)

	var u User
	if try.ORMIsEmpty(db.orm.Scopes(userFetchScope).First(&u, id)) {
		return nil, nil
	}
	return &u, nil
}

func (db *Database) UserByName(name string) (user *User, caught error) {
	defer try.Returnf(&caught, "error getting user with name %s", name)

	var u User
	if try.ORMIsEmpty(db.orm.Scopes(userFetchScope).Where("name = ?", name).First(&u)) {
		return nil, nil
	}
	return &u, nil
}

func (db *Database) UserNameExists(name string) (exists bool, caught error) {
	defer try.Returnf(&caught, "error checking username %s", name)

	c := try.Int(db.UsersCount("name = ?", name))
	return c > 0, nil
}

func (db *Database) UsersGet(query ...interface{}) (users []*User, caught error) {
	defer try.Returnf(&caught, "error getting users")

	var us []*User
	if len(query) > 0 {
		try.ORM(db.orm.Scopes(userFetchScope).Where(query[0], query[1:]...).Find(&us))
	} else {
		try.ORM(db.orm.Scopes(userFetchScope).Find(&us))
	}
	return us, nil
}

func (db *Database) UsersCount(query ...interface{}) (count int, caught error) {
	defer try.Returnf(&caught, "error counting files")

	var c int64
	if len(query) > 0 {
		try.ORM(db.orm.Model(User{}).Where(query[0], query[1:]...).Count(&c))
	} else {
		try.ORM(db.orm.Model(User{}).Count(&c))
	}
	return int(c), nil
}

func (db *Database) UserDeleteByID(id uint) (caught error) {
	defer try.Returnf(&caught, "error deleting user %d", id)

	try.ORM(db.orm.Delete(&User{}, id))
	return
}

func userFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("Permission")
}

func (u *User) check(db *Database) error {
	if try.Int(db.UsersCount("name = ? AND id <> ?", u.Name, u.ID)) > 0 {
		return fmt.Errorf("username %s already exists", u.Name)
	}

	if u.Permission == nil {
		u.PermissionID = 0
		u.Permission = &Permission{}
	}
	return nil
}
