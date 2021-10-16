package database

import (
	"fmt"

	"code.thetadev.de/ThetaDev/gotry/try"
	"gorm.io/gorm"
)

func (db *Database) VersionFileAdd(versionFile *VersionFile) (caught try.Err) {
	defer try.Annotate(&caught, "error adding versionfile")

	versionFile.ID = 0
	try.Check(versionFile.check())
	tryORM(db.orm.Create(&versionFile))
	return
}

func (db *Database) VersionFileByID(id uint) (
	versionFile *VersionFile, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting versionfile %d", id))

	var f VersionFile
	if tryORMIsEmpty(db.orm.Scopes(versionFileFetchScope).First(&f, id)) {
		return nil, nil
	}
	return &f, nil
}

func (db *Database) VersionFilesGet(query ...interface{}) (
	versionFiles []*VersionFile, caught try.Err) {

	defer try.Annotate(&caught, "error getting versionfiles")

	var vfs []*VersionFile
	if len(query) > 0 {
		tryORM(db.orm.Scopes(versionFileFetchScope).Where(query[0], query[1:]...).Find(&vfs))
	} else {
		tryORM(db.orm.Scopes(versionFileFetchScope).Find(&vfs))
	}
	return vfs, nil
}

func (db *Database) VersionFilesCount(query ...interface{}) (
	count int, caught try.Err) {

	defer try.Annotate(&caught, "error counting versionfiles")

	var c int64
	if len(query) > 0 {
		tryORM(db.orm.Model(VersionFile{}).Where(
			query[0], query[1:]...).Count(&c))
	} else {
		tryORM(db.orm.Model(VersionFile{}).Count(&c))
	}
	return int(c), nil
}

func versionFileFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("Version").Joins("File")
}

func (vf *VersionFile) check() error {
	if !isRelSet(vf.VersionID, vf.Version) {
		return ErrEmptyVersion
	}

	if !isRelSet(vf.FileID, vf.File) {
		return ErrEmptyFile
	}
	return nil
}
