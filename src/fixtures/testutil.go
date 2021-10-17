package fixtures

import (
	"os"
	"path/filepath"

	"code.thetadev.de/ThetaDev/gotry/try"
)

func doesFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func getProjectRoot() string {
	p := try.String(os.Getwd())

	for i := 0; i < 10; i++ {
		if doesFileExist(filepath.Join(p, "go.mod")) {
			return p
		}
		p = filepath.Join(p, "..")
	}

	panic("Could not find project root")
}

func CdProjectRoot() {
	root := getProjectRoot()
	try.Check(os.Chdir(root))
}

func GetTestfilesDir() string {
	CdProjectRoot()
	return filepath.Join("src", "fixtures", "testfiles")
}

func WriteTestfile(tfile string) {
	f := try.File(os.Create(tfile))
	defer f.Close()
	try.Int(f.WriteString("HelloTST"))
}
