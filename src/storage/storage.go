package storage

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/util"
)

const (
	dirFiles = "files"
	// dirCompressed = "compressed"
)

var fileCandidates = []string{"", "index.html", "index.htm"}

type Storage struct {
	path string
	db   *database.Database
}

func New(path string, db *database.Database) (s *Storage, caught try.Err) {
	defer try.Annotate(&caught, "error creating storage at "+path)

	try.Check(util.CreateDirIfNotExists(filepath.Join(path, dirFiles)))
	// util.CreateDirIfNotExists(filepath.Join(path, dirCompressed))

	return &Storage{path: path, db: db}, nil
}

func getFileHash(filePath string) (hash string, caught try.Err) {
	defer try.Annotate(&caught, "error getting file hash for "+filePath)

	file := try.File(os.Open(filePath))
	defer file.Close()

	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			try.Int(sha256.Write(buf[:n]))
		}

		if err == io.EOF {
			break
		}

		try.Check(err)
	}

	sum := sha256.Sum(nil)

	return fmt.Sprintf("%x", sum), nil
}

func (s *Storage) getFilePath(hash, subdir string) (fpath, dirpath string) {
	dirpath = filepath.Join(s.path, subdir, hash[:2])
	fpath = filepath.Join(dirpath, hash)
	return
}

func (s *Storage) getVersionFileHash(vid uint, versionFilePath string) (
	hash string, caught try.Err) {
	defer try.Return(&caught)

	for _, c := range fileCandidates {
		fp := path.Join(versionFilePath, c)
		vfile := try.X(s.db.VersionFileByPath(vid, fp)).(*database.VersionFile)

		if vfile != nil && vfile.File != nil {
			return vfile.File.Hash, nil
		}
	}
	try.Check(ErrFileNotFound)
	return
}

func (s *Storage) GetFile(sitePath string) (filePath string, caught try.Err) {
	defer try.Annotate(&caught, "error getting file at url "+sitePath)

	trimmedpath := util.TrimPath(sitePath)
	pathsegs := strings.Split(trimmedpath, "/")

	var websitePath, versionFilePath string
	var website *database.Website

	for i := len(pathsegs); i >= 0; i-- {
		websitePath = strings.Join(pathsegs[0:i], "/")
		versionFilePath = strings.Join(pathsegs[i:], "/")

		website = try.X(s.db.WebsiteByPath(websitePath, false)).(*database.Website)
		if website != nil {
			break
		}
	}

	if website == nil {
		try.Check(ErrFileNotFound)
		return
	}

	vid := try.Uint(s.db.VersionIDByWebsite(website.ID, ""))
	if vid == 0 {
		try.Check(ErrFileNotFound)
		return
	}

	hash := try.String(s.getVersionFileHash(vid, versionFilePath))

	filePath, _ = s.getFilePath(hash, "files")

	return
}

func (s *Storage) StoreFile(versionId uint, sourcePath string, sitePath string) (
	caught try.Err) {
	defer try.Annotate(&caught, "error storing file "+sourcePath)

	/*
		TODO: add html file preparation
		ext := filepath.Ext(sourcePath)
		if ext == "html" {
			try.Check(prepareHtmlFile(sourcePath))
		}*/

	hash := try.String(getFileHash(sourcePath))
	storePath, storeDir := s.getFilePath(hash, "files")

	file := &database.File{Hash: hash}
	try.Check(s.db.FileAdd(file))

	if sitePath != "" {
		vfile := &database.VersionFile{
			Path:      sitePath,
			VersionID: versionId,
			File:      file,
		}
		try.Check(s.db.VersionFileAdd(vfile))
	}

	try.Check(util.CreateDirIfNotExists(storeDir))
	try.Check(util.CopyFile(sourcePath, storePath))
	return
}
