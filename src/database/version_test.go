package database_test

import (
	"testing"

	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/Theta-Dev/Talon/src/try"
	"github.com/stretchr/testify/assert"
)

func TestVersionAdd(t *testing.T) {
	db := testdb.Open()

	t.Run("ok", func(t *testing.T) {
		version := &database.Version{
			Name:    "#2",
			Tags:    `{"test": {"v": "1"}}`,
			Website: testdb.Websites[1],
			User:    testdb.Users[1],
		}
		try.Check(db.VersionAdd(version))

		id := len(testdb.Versions) + 1
		assert.EqualValues(t, id, version.ID)

		gotV := try.X(db.VersionByID(uint(id))).(*database.Version)
		assert.Equal(t, "#2", gotV.Name)
		assert.Equal(t, `{"test": {"v": "1"}}`, gotV.Tags)
		assert.EqualValues(t, 2, gotV.Website.ID)
		assert.EqualValues(t, 2, gotV.User.ID)
	})

	t.Run("duplicate_name", func(t *testing.T) {
		version := &database.Version{
			Name:    "v0.1.1",
			Website: testdb.Websites[0],
			User:    testdb.Users[0],
		}
		err := db.VersionAdd(version)

		assert.EqualError(t, err,
			"error adding version: version name v0.1.1 already exists in website 1")
	})

	t.Run("no_user", func(t *testing.T) {
		version := &database.Version{
			Name:    "#3",
			Website: testdb.Websites[1],
		}
		err := db.VersionAdd(version)

		assert.EqualError(t, err, "error adding version: no user")
	})

	t.Run("no_website", func(t *testing.T) {
		version := &database.Version{
			Name: "#4",
			User: testdb.Users[1],
		}
		err := db.VersionAdd(version)

		assert.EqualError(t, err, "error adding version: no website")
	})
}

func TestVersionUpdate(t *testing.T) {
	db := testdb.Open()

	version := try.X(db.VersionByID(1)).(*database.Version)

	t.Run("ok", func(t *testing.T) {
		version.Name = "TestV"
		try.Check(db.VersionUpdate(version))

		gotVersion := try.X(db.VersionByID(1)).(*database.Version)
		assert.EqualValues(t, 1, gotVersion.ID)
		assert.Equal(t, "TestV", gotVersion.Name)
	})

	t.Run("duplicate_name", func(t *testing.T) {
		version.Name = "v0.1.1"
		err := db.VersionUpdate(version)
		assert.EqualError(t, err,
			"error updating version 1: version name v0.1.1 already exists in website 1")
	})
}

func TestVersionByID(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		v := try.X(db.VersionByID(1)).(*database.Version)

		assert.EqualValues(t, 1, v.ID)
		assert.Equal(t, "v0.1.0", v.Name)
		assert.Equal(t, "ThetaDev", v.Website.Name)
		assert.Equal(t, "ThetaDev", v.User.Name)
		assert.Len(t, v.Files, 2)
	})

	t.Run("not_found", func(t *testing.T) {
		noWs := try.X(db.VersionByID(0)).(*database.Version)
		assert.Nil(t, noWs)
	})
}

func TestVersionsGet(t *testing.T) {
	db := testdb.Open()

	t.Run("all", func(t *testing.T) {
		versions := try.X(db.VersionsGet()).([]*database.Version)

		for i, v := range versions {
			assert.Equal(t, testdb.Versions[i].Name, v.Name)
			assert.NotEqualValues(t, 0, v.User.ID)
			assert.NotEqualValues(t, 0, v.Website.ID)
			assert.NotEmpty(t, v.Files)
		}
	})

	t.Run("with_name", func(t *testing.T) {
		versions := try.X(db.VersionsGet("versions.name = ?", "v0.1.0")).([]*database.Version)
		assert.Len(t, versions, 1)

		v := versions[0]
		assert.Equal(t, "v0.1.0", v.Name)
		assert.Equal(t, "ThetaDev", v.Website.Name)
		assert.Equal(t, "ThetaDev", v.User.Name)
		assert.Len(t, v.Files, 2)
	})
}

func TestVersionsCount(t *testing.T) {
	db := testdb.Open()

	params := []struct {
		name   string
		query  []interface{}
		expect int
	}{
		{
			name:   "all",
			query:  []interface{}{},
			expect: 4,
		},
		{
			name:   "with_name",
			query:  []interface{}{"versions.name = ?", "#1"},
			expect: 2,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			count := try.Int(db.VersionsCount(p.query...))
			assert.Equal(t, p.expect, count)
		})
	}
}

func TestVersionDeleteByID(t *testing.T) {
	db := testdb.Open()

	try.Check(db.VersionDeleteByID(2))

	gotVersion := try.X(db.VersionByID(2)).(*database.Version)
	assert.Nil(t, gotVersion)
}
