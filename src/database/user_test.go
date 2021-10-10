package database_test

import (
	"testing"

	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/Theta-Dev/Talon/src/try"
	"github.com/stretchr/testify/assert"
)

func TestUserAdd(t *testing.T) {
	db := testdb.Open()

	params := []struct {
		name string
		user *database.User
	}{
		{
			name: "with_perm",
			user: &database.User{
				Name:         "Tris",
				PasswordHash: "hash",
				Permission: &database.Permission{
					IsAdmin:       false,
					CanCreate:     true,
					AllowedPaths:  "Talon",
					MaxSize:       5,
					MaxVersions:   3,
					MaxVisibility: 1,
				},
			},
		},
		{
			name: "empty_perm",
			user: &database.User{
				Name:         "Tris",
				PasswordHash: "hash",
				Permission:   &database.Permission{},
			},
		},
		{
			name: "nil_perm",
			user: &database.User{
				Name:         "Tris",
				PasswordHash: "hash",
			},
		},
	}

	for i, p := range params {
		t.Run(p.name, func(t *testing.T) {
			try.Check(db.UserAdd(p.user))

			assert.EqualValues(t, 4+i, p.user.ID)

			gotUser := try.X(db.UserByID(p.user.ID)).(*database.User)

			assert.EqualValues(t, 4+i, gotUser.ID)
			assert.Equal(t, "Tris", gotUser.Name)
			assert.Equal(t, "hash", gotUser.PasswordHash)
			assert.EqualValues(t, p.user.Permission, gotUser.Permission)
		})
	}
}

func TestUserByID(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		user := try.X(db.UserByID(1)).(*database.User)

		assert.EqualValues(t, 1, user.ID)
		assert.Equal(t, "ThetaDev", user.Name)
	})

	t.Run("not_found", func(t *testing.T) {
		noUser := try.X(db.UserByID(0)).(*database.User)
		assert.Nil(t, noUser)
	})
}

func TestUserByName(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		user := try.X(db.UserByName("ThetaDev")).(*database.User)

		assert.EqualValues(t, 1, user.ID)
		assert.Equal(t, "ThetaDev", user.Name)
	})

	t.Run("not_found", func(t *testing.T) {
		noUser := try.X(db.UserByName("XYZ")).(*database.User)
		assert.Nil(t, noUser)
	})
}
