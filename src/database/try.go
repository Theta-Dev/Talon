package database

import (
	"errors"

	"code.thetadev.de/ThetaDev/gotry/try"
	"gorm.io/gorm"
)

func tryGormDB(v *gorm.DB, err error) *gorm.DB {
	try.Check(err)
	return v
}

func tryORM(v *gorm.DB) {
	try.Check(v.Error)
}

func tryORMIsEmpty(v *gorm.DB) bool {
	if errors.Is(v.Error, gorm.ErrRecordNotFound) {
		return true
	}
	try.Check(v.Error)
	return false
}
