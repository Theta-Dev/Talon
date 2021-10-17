package util

import (
	"os"
	"path/filepath"
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestDoesFileExist(t *testing.T) {
	fixtures.CdProjectRoot()
	assert.True(t, DoesFileExist("go.sum"))
	assert.False(t, DoesFileExist("banana.txt"))
	assert.True(t, DoesFileExist(filepath.Join("src", "database", "database.go")))
	assert.False(t, DoesFileExist(filepath.Join("src", "database", "banana.txt")))
	assert.False(t, DoesFileExist(filepath.Join("src", "banana", "database.go")))
}

func TestHashPassword(t *testing.T) {
	hash := try.String(HashPassword("1234"))
	assert.True(t, CheckPasswordHash("1234", hash))
	assert.False(t, CheckPasswordHash("12345", hash))
}

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		name     string
		sitePath string
		want     string
	}{
		{
			name:     "uppercase",
			sitePath: "Test",
			want:     "test",
		},
		{
			name:     "slashes",
			sitePath: "/test/page/",
			want:     "test/page",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizePath(tt.sitePath)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTmpdir(t *testing.T) {
	td := try.String(NewTmpdir())

	tfile := filepath.Join(td, "test.txt")
	f := try.File(os.Create(tfile))
	try.Int(f.WriteString("Hello"))
	try.Check(f.Close())

	assert.FileExists(t, tfile)

	assert.Equal(t, 1, PurgeTmpdirs())
	assert.NoFileExists(t, tfile)
}

func TestFileSize(t *testing.T) {
	td := try.String(NewTmpdir())
	defer os.RemoveAll(td)

	tfile := filepath.Join(td, "test.txt")
	fixtures.WriteTestfile(tfile)

	size := try.Int64(FileSize(tfile))
	assert.EqualValues(t, 8, size)
}

func TestCopyFile(t *testing.T) {
	td := try.String(NewTmpdir())
	defer os.RemoveAll(td)

	tfile := filepath.Join(td, "test.txt")
	fixtures.WriteTestfile(tfile)

	cpfile := filepath.Join(td, "test2.txt")
	try.Check(CopyFile(tfile, cpfile))

	assert.FileExists(t, tfile)
	assert.FileExists(t, cpfile)

	size := try.Int64(FileSize(cpfile))
	assert.EqualValues(t, 8, size)
}
