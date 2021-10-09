package util

import (
	"path"
	"testing"

	"github.com/Theta-Dev/Talon/src/try"
	"github.com/Theta-Dev/Talon/tests"
	"github.com/stretchr/testify/assert"
)

func TestDoesFileExist(t *testing.T) {
	tests.CdProjectRoot()
	assert.True(t, DoesFileExist("go.sum"))
	assert.False(t, DoesFileExist("banana.txt"))
	assert.True(t, DoesFileExist(path.Join("src", "database", "models.go")))
	assert.False(t, DoesFileExist(path.Join("src", "database", "banana.txt")))
	assert.False(t, DoesFileExist(path.Join("src", "banana", "models.go")))
}

func TestHashPassword(t *testing.T) {
	hash := try.String(HashPassword("1234"))
	assert.True(t, CheckPasswordHash("1234", hash))
	assert.False(t, CheckPasswordHash("12345", hash))
}
