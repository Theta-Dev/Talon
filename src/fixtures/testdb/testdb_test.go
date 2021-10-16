package testdb

import (
	"regexp"
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/stretchr/testify/assert"
)

func TestEmptyAllTables(t *testing.T) {
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

	assert.EqualValues(t, len(Users)+1, u1.ID)
	// After emptying the table, IDs should start with 1 again
	assert.EqualValues(t, 1, u2.ID)
}

func TestExec(t *testing.T) {
	db := Open()

	n := try.Int(db.Exec("DELETE FROM users"))
	assert.Equal(t, len(Users), n)
}

func TestGetVersion(t *testing.T) {
	db := Open()

	v := try.String(db.GetVersion())
	reg := regexp.MustCompile(`\d\.\d`)
	assert.Regexp(t, reg, v)
}
