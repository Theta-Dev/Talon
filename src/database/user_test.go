package database_test

import (
	"testing"

	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/Theta-Dev/Talon/src/try"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	db := testdb.Open()

	createdUser := try.X(db.AddUser("Tris", "mypw", &database.Permission{
		CanCreate:     true,
		AllowedPaths:  "/dauntless",
		MaxSize:       5,
		MaxVersions:   3,
		MaxVisibility: 1,
	})).(*database.User)

	assert.EqualValues(t, 4, createdUser.ID)

	gotUser := try.X(db.GetUserByID(createdUser.ID)).(*database.User)

	assert.EqualValues(t, 4, gotUser.ID)
	assert.Equal(t, "Tris", gotUser.Name)
}

func TestGetUserByID(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		user := try.X(db.GetUserByID(1)).(*database.User)

		assert.EqualValues(t, 1, user.ID)
		assert.Equal(t, "ThetaDev", user.Name)
	})

	t.Run("not_found", func(t *testing.T) {
		noUser, err := db.GetUserByID(0)
		assert.Nil(t, noUser)
		assert.Equal(t, "error getting user: record not found", err.Error())
	})
}

func TestGetUserByName(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		user := try.X(db.GetUserByName("ThetaDev")).(*database.User)

		assert.EqualValues(t, 1, user.ID)
		assert.Equal(t, "ThetaDev", user.Name)
	})

	t.Run("not_found", func(t *testing.T) {
		noUser := try.X(db.GetUserByName("XYZ")).(*database.User)
		assert.Nil(t, noUser)
	})
}

func TestLoginUser(t *testing.T) {
	db := testdb.Open()

	t.Run("success", func(t *testing.T) {
		user := try.X(db.LoginUser("ThetaDev", "1234")).(*database.User)

		assert.EqualValues(t, 1, user.ID)
		assert.Equal(t, "ThetaDev", user.Name)
	})

	t.Run("wrong_username", func(t *testing.T) {
		user, err := db.LoginUser("XYZ", "1234")

		assert.Nil(t, user)
		assert.Equal(t, "error logging in: username/password wrong", err.Error())
	})

	t.Run("wrong_password", func(t *testing.T) {
		user, err := db.LoginUser("ThetaDev", "bananas")

		assert.Nil(t, user)
		assert.Equal(t, "error logging in: username/password wrong", err.Error())
	})
}

func TestUpdateUserPassword(t *testing.T) {
	db := testdb.Open()

	t.Run("success", func(t *testing.T) {
		try.Check(db.UpdateUserPassword(1, "1234", "bananas"))

		user, _ := db.LoginUser("ThetaDev", "1234")
		assert.Nil(t, user)
		user = try.X(db.LoginUser("ThetaDev", "bananas")).(*database.User)
		assert.EqualValues(t, 1, user.ID)
	})

	t.Run("fail", func(t *testing.T) {
		err := db.UpdateUserPassword(1, "1234", "bananas")
		assert.EqualValues(t, "error changing password: username/password wrong", err.Error())
	})
}
