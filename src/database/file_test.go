package database_test

import (
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/stretchr/testify/assert"
)

func TestFileAdd(t *testing.T) {
	db := testdb.Open()

	t.Run("ok", func(t *testing.T) {
		file := &database.File{Hash: "testHash"}
		try.Check(db.FileAdd(file))

		gotFile := try.X(db.FileByID(file.ID)).(*database.File)
		assert.Equal(t, file.ID, gotFile.ID)
		assert.Equal(t, "testHash", gotFile.Hash)
	})

	t.Run("duplicate", func(t *testing.T) {
		file := &database.File{Hash: "testHash"}
		err := db.FileAdd(file)
		assert.ErrorIs(t, err, database.ErrFileHashAlreadyExists)
	})
}

func TestFileByID(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		file := try.X(db.FileByID(1)).(*database.File)

		assert.EqualValues(t, 1, file.ID)
		assert.Equal(t, testdb.Files[0].Hash, file.Hash)
	})

	t.Run("not_found", func(t *testing.T) {
		noFile := try.X(db.FileByID(0)).(*database.File)
		assert.Nil(t, noFile)
	})
}

func TestFileByHash(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		file := try.X(db.FileByHash(testdb.Files[0].Hash)).(*database.File)

		assert.EqualValues(t, 1, file.ID)
		assert.Equal(t, testdb.Files[0].Hash, file.Hash)
	})

	t.Run("not_found", func(t *testing.T) {
		noFile := try.X(db.FileByHash("XYZ")).(*database.File)
		assert.Nil(t, noFile)
	})
}

func TestFileHashExists(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		found := try.Bool(db.FileHashExists(testdb.Files[0].Hash))
		assert.True(t, found)
	})

	t.Run("not_found", func(t *testing.T) {
		found := try.Bool(db.FileHashExists("XYZ"))
		assert.False(t, found)
	})
}

func TestFilesGet(t *testing.T) {
	db := testdb.Open()

	t.Run("all", func(t *testing.T) {
		files := try.X(db.FilesGet()).([]*database.File)

		for _, f := range files {
			i := f.ID - 1
			assert.Equal(t, testdb.Files[i].Hash, f.Hash)
		}
	})

	t.Run("with_hash", func(t *testing.T) {
		files := try.X(db.FilesGet("hash = ?", testdb.Files[0].Hash)).([]*database.File)
		assert.Len(t, files, 1)
		assert.Equal(t, testdb.Files[0].Hash, files[0].Hash)
	})
}

func TestFilesCount(t *testing.T) {
	db := testdb.Open()

	t.Run("all", func(t *testing.T) {
		count := try.Int(db.FilesCount())
		assert.Equal(t, len(testdb.Files), count)
	})

	t.Run("with_hash", func(t *testing.T) {
		count := try.Int(db.FilesCount("hash = ?", testdb.Files[0].Hash))
		assert.Equal(t, 1, count)
	})
}

func TestFilesGetOrphans(t *testing.T) {
	db := testdb.Open()

	orphans := try.X(db.FilesGetOrphans()).([]*database.File)
	assert.Len(t, orphans, 2)

	orphanHashes := []string{orphans[0].Hash, orphans[1].Hash}
	assert.Contains(t, orphanHashes, testdb.Files[13].Hash)
	assert.Contains(t, orphanHashes, testdb.Files[14].Hash)
}

func TestFilesDeleteOrphans(t *testing.T) {
	db := testdb.Open()

	try.Check(db.FilesDeleteOrphans())
	assert.Equal(t, len(testdb.Files)-2, try.Int(db.FilesCount()))

	orphans := try.X(db.FilesGetOrphans()).([]*database.File)
	assert.Len(t, orphans, 0)
}
