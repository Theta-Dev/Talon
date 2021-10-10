package testdb

import (
	"testing"

	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/try"
	"github.com/stretchr/testify/assert"
)

func TestOpenDB(t *testing.T) {
	db := Open()
	u1 := &database.User{
		Name:         "Katniss",
		PasswordHash: "1234",
	}
	u2 := &database.User{
		Name:         "Katniss",
		PasswordHash: "1234",
	}

	try.Check(db.UserAdd(u1))
	EmptyAllTables(db)
	try.Check(db.UserAdd(u2))

	assert.EqualValues(t, 4, u1.ID)
	// After emptying the table, IDs should start with 1 again
	assert.EqualValues(t, 1, u2.ID)
}
