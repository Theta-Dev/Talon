package try

// gotry auto-generated type definitions. DO NOT EDIT.

import (
	"gorm.io/gorm"
)

// DB is a helper method to handle errors of
// func() (*gorm.DB, error) functions.
func DB(v *gorm.DB, err error) *gorm.DB {
	Check(err)
	return v
}
