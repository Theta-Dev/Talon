package fixtures

import (
	"os"
	"path"
	"testing"

	"github.com/Theta-Dev/Talon/src/try"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectRoot(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		root := GetProjectRoot()
		assert.True(t, doesFileExist(path.Join(root, "go.sum")))
	})

	t.Run("subdir", func(t *testing.T) {
		root1 := GetProjectRoot()
		try.Check(os.Chdir(path.Join(root1, "src/database")))

		root := GetProjectRoot()
		assert.True(t, doesFileExist(path.Join(root, "go.sum")))
	})
}

func TestCdProjectRoot(t *testing.T) {
	CdProjectRoot()
	try.Check(os.Chdir("src/database"))
	CdProjectRoot()
	assert.True(t, doesFileExist("go.sum"))
}
