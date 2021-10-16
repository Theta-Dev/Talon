package database

import (
	"fmt"

	"code.thetadev.de/ThetaDev/gotry/try"
	"gorm.io/gorm"
)

func (db *Database) VersionAdd(version *Version) (caught try.Err) {
	defer try.Annotate(&caught, "error adding version")

	version.ID = 0
	try.Check(version.check(db))
	tryORM(db.orm.Create(&version))
	return
}

func (db *Database) VersionUpdate(version *Version) (caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error updating version %d", version.ID))

	try.Check(version.check(db))
	tryORM(db.orm.Save(&version))
	return
}

func (db *Database) VersionByID(id uint) (version *Version, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting version %d", id))

	var f Version
	if tryORMIsEmpty(db.orm.Scopes(versionFetchScope).First(&f, id)) {
		return nil, nil
	}
	return &f, nil
}

func (db *Database) VersionsGet(query ...interface{}) (
	versions []*Version, caught try.Err) {

	defer try.Annotate(&caught, "error getting websites")

	var vs []*Version
	if len(query) > 0 {
		tryORM(db.orm.Scopes(versionFetchScope).Where(query[0], query[1:]...).Find(&vs))
	} else {
		tryORM(db.orm.Scopes(versionFetchScope).Find(&vs))
	}
	return vs, nil
}

func (db *Database) VersionsCount(query ...interface{}) (
	count int, caught try.Err) {

	defer try.Annotate(&caught, "error counting versions")

	var c int64
	if len(query) > 0 {
		tryORM(db.orm.Model(Version{}).Where(query[0], query[1:]...).Count(&c))
	} else {
		tryORM(db.orm.Model(Version{}).Count(&c))
	}
	return int(c), nil
}

func (db *Database) VersionDeleteByID(id uint) (caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error deleting version %d", id))

	tryORM(db.orm.Delete(&Version{}, id))
	return
}

func versionFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("User").Joins("Website").Preload("Files.File")
}

func (v *Version) check(db *Database) try.Err {
	if !isRelSet(v.UserID, v.User) {
		return try.FromErr(ErrEmptyUser)
	}

	wsid := getRelId(v.WebsiteID, v.Website)
	if wsid == 0 {
		return try.FromErr(ErrEmptyWebsite)
	}

	if try.Int(db.VersionsCount(
		"name = ? AND id <> ? AND website_id = ?", v.Name, v.ID, wsid)) > 0 {
		return newErrVersionNameAlreadyExists(v.Name, wsid)
	}

	return nil
}
