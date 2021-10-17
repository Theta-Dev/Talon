package database

import (
	"fmt"

	"code.thetadev.de/ThetaDev/gotry/try"
	"gorm.io/gorm"
)

func (db *Database) VersionFileAdd(versionFile *VersionFile) (caught try.Err) {
	defer try.Annotate(&caught, "error adding versionfile")

	versionFile.ID = 0
	try.Check(versionFile.check(db))
	tryORM(db.orm.Create(&versionFile))
	return
}

func (db *Database) VersionFileByID(id uint, deep bool) (
	versionFile *VersionFile, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting versionfile %d", id))

	orm := db.orm
	if deep {
		orm = orm.Scopes(versionFileFetchScope)
	}

	if tryORMIsEmpty(orm.First(&versionFile, id)) {
		return nil, nil
	}
	return
}

func (db *Database) VersionFileByPath(versionId uint, filePath string) (
	versionFile *VersionFile, caught try.Err) {

	defer try.Annotate(&caught, fmt.Sprintf("error getting version file %s for version %d",
		filePath, versionId))

	if tryORMIsEmpty(db.orm.Joins("File").
		Where("version_id = ? AND path = ?", versionId, filePath).First(&versionFile)) {
		return nil, nil
	}
	return
}

func (db *Database) VersionFilesGet(deep bool, query ...interface{}) (
	versionFiles []*VersionFile, caught try.Err) {

	defer try.Annotate(&caught, "error getting versionfiles")

	orm := db.orm
	if deep {
		orm = orm.Scopes(versionFileFetchScope)
	}

	if len(query) > 0 {
		tryORMIsEmpty(orm.Where(query[0], query[1:]...).Find(&versionFiles))
	} else {
		tryORMIsEmpty(orm.Find(&versionFiles))
	}
	return
}

func (db *Database) VersionFilesCount(query ...interface{}) (
	count int64, caught try.Err) {

	defer try.Annotate(&caught, "error counting versionfiles")

	if len(query) > 0 {
		tryORM(db.orm.Model(VersionFile{}).Where(
			query[0], query[1:]...).Count(&count))
	} else {
		tryORM(db.orm.Model(VersionFile{}).Count(&count))
	}
	return
}

func versionFileFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("Version").Joins("File")
}

func (vf *VersionFile) check(db *Database) try.Err {
	vid := getRelId(vf.VersionID, vf.Version)
	if vid == 0 {
		return try.FromErr(ErrEmptyVersion)
	}

	if !isRelSet(vf.FileID, vf.File) {
		return try.FromErr(ErrEmptyFile)
	}

	if try.Int64(db.VersionFilesCount("version_id = ? AND path = ? AND id <> ?",
		vid, vf.Path, vf.ID)) > 0 {
		return newErrVersionFileAlreadyExists(vf.Path, vid)
	}
	return nil
}
