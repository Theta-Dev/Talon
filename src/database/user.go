package database

import (
	"fmt"

	"code.thetadev.de/ThetaDev/gotry/try"
	"gorm.io/gorm"
)

func (db *Database) UserAdd(user *User) (caught try.Err) {
	defer try.Annotate(&caught, "error adding user")

	user.ID = 0
	try.Check(user.check(db))
	user.Permission.ID = 0

	try.Check(db.orm.Transaction(func(tx *gorm.DB) (err error) {
		try.ReturnStd(&err)

		tryORM(tx.Create(user.Permission))
		tryORM(tx.Create(user))
		return
	}))
	return
}

func (db *Database) UserUpdate(user *User) (caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error updating user %d", user.ID))

	try.Check(user.check(db))
	tryORM(db.orm.Save(user))
	return
}

func (db *Database) UserByID(id uint) (user *User, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting user %d", id))

	if tryORMIsEmpty(db.orm.Scopes(userFetchScope).First(&user, id)) {
		return nil, nil
	}
	return
}

func (db *Database) UserByName(name string) (user *User, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting user with name %s", name))

	if tryORMIsEmpty(db.orm.Scopes(userFetchScope).Where("name = ?", name).First(&user)) {
		return nil, nil
	}
	return
}

func (db *Database) UserNameExists(name string) (exists bool, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error checking username %s", name))

	c := try.Int64(db.UsersCount("name = ?", name))
	return c > 0, nil
}

func (db *Database) UsersGet(query ...interface{}) (users []*User, caught try.Err) {
	defer try.Annotate(&caught, "error getting users")

	if len(query) > 0 {
		tryORMIsEmpty(
			db.orm.Scopes(userFetchScope).Where(query[0], query[1:]...).Find(&users))
	} else {
		tryORMIsEmpty(db.orm.Scopes(userFetchScope).Find(&users))
	}
	return
}

func (db *Database) UsersCount(query ...interface{}) (count int64, caught try.Err) {
	defer try.Annotate(&caught, "error counting files")

	if len(query) > 0 {
		tryORM(db.orm.Model(User{}).Where(query[0], query[1:]...).Count(&count))
	} else {
		tryORM(db.orm.Model(User{}).Count(&count))
	}
	return
}

func (db *Database) UserDeleteByID(id uint) (caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error deleting user %d", id))

	tryORM(db.orm.Delete(&User{}, id))
	return
}

func userFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("Permission")
}

func (u *User) check(db *Database) try.Err {
	if try.Int64(db.UsersCount("name = ? AND id <> ?", u.Name, u.ID)) > 0 {
		return newErrUsernameAlreadyExists(u.Name)
	}

	if u.Permission == nil {
		u.PermissionID = 0
		u.Permission = &Permission{}
	}
	return nil
}
