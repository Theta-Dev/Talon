package database

import (
	"fmt"

	"github.com/Theta-Dev/Talon/src/try"
	"gorm.io/gorm"
)

func (db *Database) VersionAdd(version *Version) (caught error) {
	defer try.Returnf(&caught, "error adding version")

	version.ID = 0
	try.Check(version.check(db))
	try.ORM(db.orm.Create(&version))
	return
}

func (db *Database) VersionUpdate(version *Version) (caught error) {
	defer try.Returnf(&caught, "error updating version %d", version.ID)

	try.Check(version.check(db))
	try.ORM(db.orm.Save(&version))
	return
}

func (db *Database) VersionByID(id uint) (version *Version, caught error) {
	defer try.Returnf(&caught, "error getting version %d", id)

	var f Version
	if try.ORMIsEmpty(db.orm.Scopes(versionFetchScope).First(&f, id)) {
		return nil, nil
	}
	return &f, nil
}

func (db *Database) VersionsGet(query ...interface{}) (
	versions []*Version, caught error) {

	defer try.Returnf(&caught, "error getting websites")

	var vs []*Version
	if len(query) > 0 {
		try.ORM(db.orm.Scopes(versionFetchScope).Where(query[0], query[1:]...).Find(&vs))
	} else {
		try.ORM(db.orm.Scopes(versionFetchScope).Find(&vs))
	}
	return vs, nil
}

func (db *Database) VersionsCount(query ...interface{}) (
	count int, caught error) {

	defer try.Returnf(&caught, "error counting versions")

	var c int64
	if len(query) > 0 {
		try.ORM(db.orm.Model(Version{}).Where(query[0], query[1:]...).Count(&c))
	} else {
		try.ORM(db.orm.Model(Version{}).Count(&c))
	}
	return int(c), nil
}

func (db *Database) VersionDeleteByID(id uint) (caught error) {
	defer try.Returnf(&caught, "error deleting version %d", id)

	try.ORM(db.orm.Delete(&Version{}, id))
	return
}

func versionFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("User").Joins("Website").Preload("Files.File")
}

func (v *Version) check(db *Database) error {
	if !isRelSet(v.UserID, v.User) {
		return fmt.Errorf("no user")
	}

	wsid := getRelId(v.WebsiteID, v.Website)
	if wsid == 0 {
		return fmt.Errorf("no website")
	}

	if try.Int(db.VersionsCount(
		"name = ? AND id <> ? AND website_id = ?", v.Name, v.ID, wsid)) > 0 {
		return fmt.Errorf("version name %s already exists in website %d", v.Name, wsid)
	}

	return nil
}
