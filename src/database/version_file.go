package database

import (
	"fmt"

	"github.com/Theta-Dev/Talon/src/try"
	"gorm.io/gorm"
)

func (db *Database) VersionFileAdd(versionFile *VersionFile) (caught error) {
	defer try.Returnf(&caught, "error adding versionfile")

	versionFile.ID = 0
	try.Check(versionFile.check())
	try.ORM(db.orm.Create(&versionFile))
	return
}

func (db *Database) VersionFileByID(id uint) (versionFile *VersionFile, caught error) {
	defer try.Returnf(&caught, "error getting versionfile %d", id)

	var f VersionFile
	if try.ORMIsEmpty(db.orm.Scopes(versionFileFetchScope).First(&f, id)) {
		return nil, nil
	}
	return &f, nil
}

func (db *Database) VersionFilesGet(query ...interface{}) (
	versionFiles []*VersionFile, caught error) {

	defer try.Returnf(&caught, "error getting versionfiles")

	var vfs []*VersionFile
	if len(query) > 0 {
		try.ORM(db.orm.Scopes(versionFileFetchScope).Where(query[0], query[1:]...).Find(&vfs))
	} else {
		try.ORM(db.orm.Scopes(versionFileFetchScope).Find(&vfs))
	}
	return vfs, nil
}

func (db *Database) VersionFilesCount(query ...interface{}) (
	count int, caught error) {

	defer try.Returnf(&caught, "error counting versionfiles")

	var c int64
	if len(query) > 0 {
		try.ORM(db.orm.Model(VersionFile{}).Where(
			query[0], query[1:]...).Count(&c))
	} else {
		try.ORM(db.orm.Model(VersionFile{}).Count(&c))
	}
	return int(c), nil
}

func versionFileFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("Version").Joins("File")
}

func (vf *VersionFile) check() error {
	if !isRelSet(vf.VersionID, vf.Version) {
		return fmt.Errorf("no version")
	}

	if !isRelSet(vf.FileID, vf.File) {
		return fmt.Errorf("no file")
	}
	return nil
}
