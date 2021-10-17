package database_test

import (
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
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

		gotVDeep := try.X(db.VersionByID(uint(id), true)).(*database.Version)
		assert.Equal(t, "#2", gotVDeep.Name)
		assert.Equal(t, `{"test": {"v": "1"}}`, gotVDeep.Tags)
		assert.EqualValues(t, 2, gotVDeep.Website.ID)
		assert.EqualValues(t, 2, gotVDeep.User.ID)
	})

	t.Run("duplicate_name", func(t *testing.T) {
		version := &database.Version{
			Name:    "v0.1.1",
			Website: testdb.Websites[0],
			User:    testdb.Users[0],
		}
		err := db.VersionAdd(version)

		assert.ErrorIs(t, err, database.ErrVersionNameAlreadyExists)
	})

	t.Run("no_user", func(t *testing.T) {
		version := &database.Version{
			Name:    "#3",
			Website: testdb.Websites[1],
		}
		err := db.VersionAdd(version)

		assert.ErrorIs(t, err, database.ErrEmptyUser)
	})

	t.Run("no_website", func(t *testing.T) {
		version := &database.Version{
			Name: "#4",
			User: testdb.Users[1],
		}
		err := db.VersionAdd(version)

		assert.ErrorIs(t, err, database.ErrEmptyWebsite)
	})
}

func TestVersionUpdate(t *testing.T) {
	db := testdb.Open()

	version := try.X(db.VersionByID(1, false)).(*database.Version)

	t.Run("ok", func(t *testing.T) {
		version.Name = "TestV"
		try.Check(db.VersionUpdate(version))

		gotVersion := try.X(db.VersionByID(1, false)).(*database.Version)
		assert.EqualValues(t, 1, gotVersion.ID)
		assert.Equal(t, "TestV", gotVersion.Name)
	})

	t.Run("duplicate_name", func(t *testing.T) {
		version.Name = "v0.1.1"
		err := db.VersionUpdate(version)
		assert.ErrorIs(t, err, database.ErrVersionNameAlreadyExists)
	})
}

func TestVersionByID(t *testing.T) {
	db := testdb.Open()

	t.Run("found_deep", func(t *testing.T) {
		v := try.X(db.VersionByID(1, true)).(*database.Version)

		assert.EqualValues(t, 1, v.ID)
		assert.Equal(t, "v0.1.0", v.Name)
		assert.Equal(t, "ThetaDev", v.Website.Name)
		assert.Equal(t, "ThetaDev", v.User.Name)
		assert.Len(t, v.Files, 2)
	})

	t.Run("found", func(t *testing.T) {
		v := try.X(db.VersionByID(1, false)).(*database.Version)

		assert.EqualValues(t, 1, v.ID)
		assert.Equal(t, "v0.1.0", v.Name)
		assert.Nil(t, v.Website)
		assert.Nil(t, v.User)
	})

	t.Run("not_found", func(t *testing.T) {
		noWs := try.X(db.VersionByID(0, false)).(*database.Version)
		assert.Nil(t, noWs)
	})
}

func TestVersionsGet(t *testing.T) {
	db := testdb.Open()

	t.Run("all", func(t *testing.T) {
		versions := try.X(db.VersionsGet(true)).([]*database.Version)

		for i, v := range versions {
			assert.Equal(t, testdb.Versions[i].Name, v.Name)
			assert.NotEqualValues(t, 0, v.User.ID)
			assert.NotEqualValues(t, 0, v.Website.ID)
			assert.NotEmpty(t, v.Files)
		}
	})

	t.Run("with_name", func(t *testing.T) {
		versions := try.X(
			db.VersionsGet(true, "versions.name = ?", "v0.1.0")).([]*database.Version)
		assert.Len(t, versions, 1)

		v := versions[0]
		assert.Equal(t, "v0.1.0", v.Name)
		assert.Equal(t, "ThetaDev", v.Website.Name)
		assert.Equal(t, "ThetaDev", v.User.Name)
		assert.Len(t, v.Files, 2)
	})

	t.Run("none", func(t *testing.T) {
		versions := try.X(db.VersionsGet(false, "versions.id = 0")).([]*database.Version)
		assert.Empty(t, versions)
	})
}

func TestVersionsCount(t *testing.T) {
	db := testdb.Open()

	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := try.Int64(db.VersionsCount(tt.query...))
			assert.EqualValues(t, tt.expect, count)
		})
	}
}

func TestVersionDeleteByID(t *testing.T) {
	db := testdb.Open()

	try.Check(db.VersionDeleteByID(2))

	gotVersion := try.X(db.VersionByID(2, false)).(*database.Version)
	assert.Nil(t, gotVersion)
}

func TestVersionIDByWebsitex(t *testing.T) {
	db := testdb.Open()

	tests := []struct {
		name        string
		websiteId   uint
		versionName string
		expect      uint
	}{
		{
			name:        "latest",
			websiteId:   1,
			versionName: "",
			expect:      2,
		},
		{
			name:        "named",
			websiteId:   1,
			versionName: "v0.1.0",
			expect:      1,
		},
		{
			name:        "none",
			websiteId:   1,
			versionName: "XYZ",
			expect:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vid := try.Uint(db.VersionIDByWebsite(tt.websiteId, tt.versionName))
			assert.Equal(t, tt.expect, vid)
		})
	}
}
