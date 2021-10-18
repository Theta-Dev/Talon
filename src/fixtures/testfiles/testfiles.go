package testfiles

import (
	"os"
	"path/filepath"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/Theta-Dev/Talon/src/storage"
)

func Open() (*storage.Storage, *database.Database) {
	tfPath := fixtures.GetTestfilesDir()
	_ = os.RemoveAll("talon_tmp")

	db := testdb.OpenWithoutFiles()

	s := try.X(storage.New("talon_tmp", db)).(*storage.Storage)

	storeFile(s, 1, tfPath, "ThetaDev0", "index.html")
	storeFile(s, 1, tfPath, "ThetaDev0", "style.css")
	storeFile(s, 1, tfPath, "ThetaDev0", "thetadev-blue.svg")

	storeFile(s, 2, tfPath, "ThetaDev1", "index.html")
	storeFile(s, 2, tfPath, "ThetaDev1", "assets/style.css")
	storeFile(s, 2, tfPath, "ThetaDev1", "assets/image.jpg")
	storeFile(s, 2, tfPath, "ThetaDev1", "assets/test.js")
	storeFile(s, 2, tfPath, "ThetaDev1", "assets/example.txt")
	storeFile(s, 2, tfPath, "ThetaDev1", "assets/thetadev-blue.svg")

	storeFile(s, 3, tfPath, "Talon", "index.html")
	storeFile(s, 3, tfPath, "Talon", "talon_style.css")
	storeFile(s, 3, tfPath, "Talon", "logo.svg")

	storeFile(s, 4, tfPath, "GenderEx", "index.html")
	storeFile(s, 4, tfPath, "GenderEx", "gex_style.css")
	storeFile(s, 4, tfPath, "GenderEx", "logo.svg")

	return s, db
}

func storeFile(s *storage.Storage, vid uint, tfPath, dir, fname string) {
	try.Check(s.StoreFile(vid, filepath.Join(tfPath, dir, fname), fname))
}
