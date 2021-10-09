package testdb

import (
	"testing"

	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/try"
	"github.com/stretchr/testify/assert"
)

func TestOpenDB(t *testing.T) {
	db := Open()
	u1 := try.X(db.AddUser("A", "1234", &database.Permission{})).(*database.User)
	EmptyAllTables(db)
	u2 := try.X(db.AddUser("A", "1234", &database.Permission{})).(*database.User)

	assert.EqualValues(t, 4, u1.ID)
	assert.EqualValues(t, 1, u2.ID)
}
