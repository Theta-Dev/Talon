package tests

import (
	"os"
	"path"

	"github.com/Theta-Dev/Talon/src/try"
)

func doesFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func GetProjectRoot() string {
	p := try.String(os.Getwd())

	for i := 0; i < 10; i++ {
		if doesFileExist(path.Join(p, "go.mod")) {
			return p
		}
		p = path.Join(p, "..")
	}

	panic("Could not find project root")
}

func CdProjectRoot() {
	root := GetProjectRoot()
	try.Check(os.Chdir(root))
}
