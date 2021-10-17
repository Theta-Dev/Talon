package database_test

import (
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/stretchr/testify/assert"
)

func TestVersionFileAdd(t *testing.T) {
	db := testdb.Open()

	t.Run("ok", func(t *testing.T) {
		vf := &database.VersionFile{
			Path:    "test.txt",
			Version: testdb.Versions[0],
			File:    testdb.Files[0],
		}
		try.Check(db.VersionFileAdd(vf))

		id := len(testdb.VersionFiles) + 1
		assert.EqualValues(t, id, vf.ID)

		gotVf := try.X(db.VersionFileByID(uint(id), true)).(*database.VersionFile)
		assert.Equal(t, "test.txt", gotVf.Path)
		assert.EqualValues(t, 1, gotVf.Version.ID)
		assert.EqualValues(t, 1, gotVf.File.ID)
	})

	t.Run("no_version", func(t *testing.T) {
		vf := &database.VersionFile{
			Path: "tmp.png",
			File: testdb.Files[1],
		}

		err := db.VersionFileAdd(vf)
		assert.ErrorIs(t, err, database.ErrEmptyVersion)
	})

	t.Run("no_file", func(t *testing.T) {
		vf := &database.VersionFile{
			Path:    "tmp.png",
			Version: testdb.Versions[0],
		}

		err := db.VersionFileAdd(vf)
		assert.ErrorIs(t, err, database.ErrEmptyFile)
	})

	t.Run("duplicate_path", func(t *testing.T) {
		vf := &database.VersionFile{
			Path:    "index.html",
			Version: testdb.Versions[0],
			File:    testdb.Files[10],
		}

		err := db.VersionFileAdd(vf)
		assert.ErrorIs(t, err, database.ErrVersionFileAlreadyExists)
	})
}

func TestVersionFileByID(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		v := try.X(db.VersionFileByID(2, true)).(*database.VersionFile)

		assert.EqualValues(t, 2, v.ID)
		assert.Equal(t, "style.css", v.Path)
		assert.Equal(t, testdb.Files[1].Hash, v.File.Hash)
		assert.Equal(t, "v0.1.0", v.Version.Name)
	})

	t.Run("not_found", func(t *testing.T) {
		noWs := try.X(db.VersionFileByID(0, false)).(*database.VersionFile)
		assert.Nil(t, noWs)
	})
}

func TestVersionFilesGet(t *testing.T) {
	db := testdb.Open()

	t.Run("all", func(t *testing.T) {
		vfiles := try.X(db.VersionFilesGet(true)).([]*database.VersionFile)

		for _, vf := range vfiles {
			i := vf.ID - 1
			assert.Equal(t, testdb.VersionFiles[i].Path, vf.Path)
			assert.NotEqualValues(t, 0, vf.File.ID)
			assert.NotEqualValues(t, 0, vf.Version.ID)
		}
	})

	t.Run("with_version", func(t *testing.T) {
		vfiles := try.X(
			db.VersionFilesGet(true, "version_id = ?", "1")).([]*database.VersionFile)
		assert.Len(t, vfiles, 2)

		vf := vfiles[0]
		if vf.ID != 1 {
			vf = vfiles[1]
		}

		assert.Equal(t, "index.html", vf.Path)
		assert.Equal(t, testdb.Files[0].Hash, vf.File.Hash)
	})

	t.Run("none", func(t *testing.T) {
		vfiles := try.X(db.VersionFilesGet(false, "id = 0")).([]*database.VersionFile)
		assert.Empty(t, vfiles)
	})
}

func TestVersionFilesCount(t *testing.T) {
	db := testdb.Open()

	tests := []struct {
		name   string
		query  []interface{}
		expect int
	}{
		{
			name:   "all",
			query:  []interface{}{},
			expect: len(testdb.VersionFiles),
		},
		{
			name:   "by_version",
			query:  []interface{}{"version_id = ?", "1"},
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := try.Int64(db.VersionFilesCount(tt.query...))
			assert.EqualValues(t, tt.expect, count)
		})
	}
}

func TestVersionFileByPath(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		vfile := try.X(db.VersionFileByPath(1, "style.css")).(*database.VersionFile)
		assert.Equal(t, testdb.Files[1].Hash, vfile.File.Hash)
	})

	t.Run("not_found", func(t *testing.T) {
		vfile := try.X(db.VersionFileByPath(1, "xyz.html")).(*database.VersionFile)
		assert.Nil(t, vfile)
	})
}
