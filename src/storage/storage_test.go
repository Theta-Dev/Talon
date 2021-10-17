package storage

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/Theta-Dev/Talon/src/util"
	"github.com/stretchr/testify/assert"
)

const test_hash = "6e99e2dbab5524a692616f208e7fbc54778dd67af6968468393a9912f51f707d"

func TestGetFileHash(t *testing.T) {
	file := filepath.Join(fixtures.GetTestfilesDir(), "Talon", "index.html")

	hash := try.String(getFileHash(file))

	assert.Equal(t, test_hash, hash)
}

func TestGetFilePath(t *testing.T) {
	storage := Storage{path: "/tst/path"}
	fpath, dirpath := storage.getFilePath(test_hash, dir_files)

	assert.Equal(t,
		"/tst/path/files/6e/6e99e2dbab5524a692616f208e7fbc54778dd67af6968468393a9912f51f707d",
		fpath)

	assert.Equal(t, "/tst/path/files/6e", dirpath)
}

func TestGetFile(t *testing.T) {
	db := testdb.Open()
	tmpdir := try.String(util.NewTmpdir())
	defer os.RemoveAll(tmpdir)

	storage := try.X(New(tmpdir, db)).(*Storage)

	filePath := try.String(storage.GetFile("index.html"))
	assert.True(t, strings.HasSuffix(filePath, "/files/53/"+testdb.Files[5].Hash))
}

func TestStoreFile(t *testing.T) {
	db := testdb.Open()
	tmpdir := try.String(util.NewTmpdir())
	testfile := filepath.Join(fixtures.GetTestfilesDir(), "Talon", "index.html")
	defer os.RemoveAll(tmpdir)

	version := &database.Version{
		Name:      "TestVersion",
		WebsiteID: 2,
		UserID:    2,
	}
	try.Check(db.VersionAdd(version))

	storage := try.X(New(tmpdir, db)).(*Storage)
	try.Check(storage.StoreFile(version.ID, testfile, "index.html"))

	file := try.X(db.FileByHash(test_hash)).(*database.File)
	assert.Equal(t, test_hash, file.Hash)

	versionFile := try.X(
		db.VersionFileByPath(version.ID, "index.html")).(*database.VersionFile)
	assert.Equal(t, file.ID, versionFile.FileID)
}
