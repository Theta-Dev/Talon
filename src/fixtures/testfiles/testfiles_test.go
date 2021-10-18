package testfiles

import (
	"path/filepath"
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	storage, db := Open()

	files := try.X(db.FilesGet()).([]*database.File)
	assert.Len(t, files, 13)

	for _, f := range files {
		i := f.ID - 1
		assert.Equal(t, testdb.Files[i].Hash, f.Hash)
	}

	sitePaths := map[string]string{
		// ThetaDev
		"":                         testdb.Files[3].Hash,
		"assets/style.css":         testdb.Files[1].Hash,
		"assets/thetadev-blue.svg": testdb.Files[2].Hash,
		"assets/test.js":           testdb.Files[5].Hash,
		"assets/example.txt":       testdb.Files[6].Hash,
		// Talon
		"Talon":                 testdb.Files[7].Hash,
		"Talon/talon_style.css": testdb.Files[8].Hash,
		"Talon/logo.svg":        testdb.Files[9].Hash,
		// GenderEx
		"Spotify-Gender-Ex":               testdb.Files[10].Hash,
		"Spotify-Gender-Ex/gex_style.css": testdb.Files[11].Hash,
		"Spotify-Gender-Ex/logo.svg":      testdb.Files[12].Hash,
	}

	storagePath := filepath.Join("talon_tmp", "files")

	for p, h := range sitePaths {
		filePath := try.String(storage.GetFile(p))
		assert.Equal(t, filepath.Join(storagePath, h[:2], h), filePath)
		assert.FileExists(t, filePath)
	}
}
