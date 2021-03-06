package database_test

import (
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/stretchr/testify/assert"
)

func TestUserAdd(t *testing.T) {
	db := testdb.Open()

	tests := []struct {
		name string
		user database.User
	}{
		{
			name: "with_perm",
			user: database.User{
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
			user: database.User{
				Name:         "Max",
				PasswordHash: "hash",
				Permission:   &database.Permission{},
			},
		},
		{
			name: "nil_perm",
			user: database.User{
				Name:         "Lynn",
				PasswordHash: "hash",
			},
		},
	}

	for i, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			try.Check(db.UserAdd(&p.user))

			assert.EqualValues(t, len(testdb.Users)+1+i, p.user.ID)

			gotUser := try.X(db.UserByID(p.user.ID)).(*database.User)

			assert.EqualValues(t, p.user.ID, gotUser.ID)
			assert.Equal(t, p.user.Name, gotUser.Name)
			assert.Equal(t, "hash", gotUser.PasswordHash)
			assert.EqualValues(t, p.user.Permission, gotUser.Permission)
		})
	}

	t.Run("duplicate_username", func(t *testing.T) {
		u := &database.User{
			Name:         "ThetaDev",
			PasswordHash: "hash",
		}

		err := db.UserAdd(u)
		assert.ErrorIs(t, err, database.ErrUsernameAlreadyExists)
	})
}

func TestUserUpdate(t *testing.T) {
	db := testdb.Open()

	user := try.X(db.UserByID(1)).(*database.User)

	t.Run("ok", func(t *testing.T) {
		user.Name = "Eric"
		try.Check(db.UserUpdate(user))

		gotUser := try.X(db.UserByName("Eric")).(*database.User)
		assert.EqualValues(t, 1, gotUser.ID)
		assert.Equal(t, "Eric", gotUser.Name)
	})

	t.Run("duplicate_name", func(t *testing.T) {
		user.Name = "Izzy"
		err := db.UserUpdate(user)
		assert.ErrorIs(t, err, database.ErrUsernameAlreadyExists)
	})
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

func TestUserNameExists(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		found := try.Bool(db.UserNameExists("ThetaDev"))
		assert.True(t, found)
	})

	t.Run("not_found", func(t *testing.T) {
		found := try.Bool(db.UserNameExists("XYZ"))
		assert.False(t, found)
	})
}

func TestUsersGet(t *testing.T) {
	db := testdb.Open()

	t.Run("all", func(t *testing.T) {
		users := try.X(db.UsersGet()).([]*database.User)

		for _, u := range users {
			i := u.ID - 1
			assert.Equal(t, testdb.Users[i].Name, u.Name)
			assert.Equal(t, testdb.Users[i].Permission.AllowedPaths, u.Permission.AllowedPaths)
		}
	})

	t.Run("with_name", func(t *testing.T) {
		users := try.X(db.UsersGet("name = ?", "Izzy")).([]*database.User)
		assert.Len(t, users, 1)
		assert.Equal(t, "Izzy", users[0].Name)
		assert.Equal(t, "tests/*", users[0].Permission.AllowedPaths)
	})

	t.Run("none", func(t *testing.T) {
		users := try.X(db.UsersGet("Name = 'XYZ'")).([]*database.User)
		assert.Empty(t, users)
	})
}

func TestUsersCount(t *testing.T) {
	db := testdb.Open()

	t.Run("all", func(t *testing.T) {
		count := try.Int64(db.UsersCount())
		assert.EqualValues(t, len(testdb.Users), count)
	})

	t.Run("with_name", func(t *testing.T) {
		count := try.Int64(db.UsersCount("name = ?", "Izzy"))
		assert.EqualValues(t, 1, count)
	})
}

func TestUserDeleteByID(t *testing.T) {
	db := testdb.Open()

	try.Check(db.UserDeleteByID(2))

	gotUser := try.X(db.UserByID(2)).(*database.User)
	assert.Nil(t, gotUser)
}
