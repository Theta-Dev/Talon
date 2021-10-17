package fixtures

import (
	"os"
	"path/filepath"
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectRoot(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		root := getProjectRoot()
		assert.True(t, doesFileExist(filepath.Join(root, "go.sum")))
	})

	t.Run("subdir", func(t *testing.T) {
		root1 := getProjectRoot()
		try.Check(os.Chdir(filepath.Join(root1, "src/database")))

		root := getProjectRoot()
		assert.True(t, doesFileExist(filepath.Join(root, "go.sum")))
	})
}

func TestCdProjectRoot(t *testing.T) {
	CdProjectRoot()
	try.Check(os.Chdir("src/database"))
	CdProjectRoot()
	assert.True(t, doesFileExist("go.sum"))
}
