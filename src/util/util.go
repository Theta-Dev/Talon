package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"code.thetadev.de/ThetaDev/gotry/try"
	"golang.org/x/crypto/bcrypt"
)

func DoesFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func CreateDirIfNotExists(dirpath string) (caught try.Err) {
	defer try.Annotate(&caught, "error creating directory "+dirpath)

	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		try.Check(os.MkdirAll(dirpath, 0o777))
	}
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TrimPath(sitePath string) string {
	return strings.Trim(sitePath, "/")
}

func NormalizePath(sitePath string) string {
	p := strings.ToLower(sitePath)
	return TrimPath(p)
}

func NewTmpdir() (tmpdir string, caught try.Err) {
	defer try.Annotate(&caught, "error creating tmpdir")

	for {
		bts := make([]byte, 16)
		try.Int(rand.Read(bts))
		tmpdir = filepath.Join(os.TempDir(), fmt.Sprintf("talon_%x", bts))

		if !DoesFileExist(tmpdir) {
			break
		}
	}

	try.Check(CreateDirIfNotExists(tmpdir))

	return
}

func PurgeTmpdirs() (count int) {
	dirs, _ := os.ReadDir(os.TempDir())

	for _, de := range dirs {
		if !de.IsDir() {
			continue
		}
		if strings.HasPrefix(de.Name(), "talon_") {
			err := os.RemoveAll(filepath.Join(os.TempDir(), de.Name()))
			if err == nil {
				count++
			}
		}
	}
	return
}

func FileSize(filePath string) (size int64, caught try.Err) {
	defer try.Annotate(&caught, "error getting size of file "+filePath)

	fileStat := try.X(os.Stat(filePath)).(fs.FileInfo)
	return fileStat.Size(), nil
}

func CopyFile(src, dst string) (caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error coping file %s to %s", src, dst))

	sourceFileStat := try.X(os.Stat(src)).(fs.FileInfo)

	if !sourceFileStat.Mode().IsRegular() {
		try.Check(ErrNoRegularFile)
	}

	source := try.File(os.Open(src))
	defer source.Close()

	destination := try.File(os.Create(dst))
	defer destination.Close()

	try.Int64(io.Copy(destination, source))
	return
}
