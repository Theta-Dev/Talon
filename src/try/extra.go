package try

import (
	"errors"

	"gorm.io/gorm"
)

// ORM is a helper method to handle errors of GORM operations.
func ORM(v *gorm.DB) *gorm.DB {
	Check(v.Error)
	return v
}

// ORMIsEmpty is a helper method to handle errors of GORM operations
// and check if the result is empty.
func ORMIsEmpty(v *gorm.DB) bool {
	if errors.Is(v.Error, gorm.ErrRecordNotFound) {
		return true
	}
	Check(v.Error)
	return false
}
