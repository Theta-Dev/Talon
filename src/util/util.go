package util

import (
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func DoesFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NormalizePath(sitePath string) string {
	p := strings.ToLower(sitePath)
	p = strings.Trim(p, "/")
	return p
}
