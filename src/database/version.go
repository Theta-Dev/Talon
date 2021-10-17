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

func (db *Database) VersionByID(id uint, deep bool) (version *Version, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting version %d", id))

	orm := db.orm
	if deep {
		orm = orm.Scopes(versionFetchScope)
	}

	if tryORMIsEmpty(orm.First(&version, id)) {
		return nil, nil
	}
	return
}

func (db *Database) VersionIDByWebsite(websiteId uint, versionName string) (
	versionId uint, caught try.Err) {

	defer try.Annotate(&caught, fmt.Sprintf("error getting version %s from website %d",
		versionName, websiteId))

	cond := "website_id = ?"
	if versionName != "" {
		cond += " AND name = ?"
	}

	tryORM(db.orm.Model(&Version{}).Select("id").
		Where(cond, websiteId, versionName).
		Order("id DESC").Limit(1).Scan(&versionId))
	return
}

func (db *Database) VersionsGet(deep bool, query ...interface{}) (
	versions []*Version, caught try.Err) {

	defer try.Annotate(&caught, "error getting versions")

	orm := db.orm
	if deep {
		orm = orm.Scopes(versionFetchScope)
	}

	if len(query) > 0 {
		tryORMIsEmpty(orm.Where(query[0], query[1:]...).Find(&versions))
	} else {
		tryORMIsEmpty(orm.Find(&versions))
	}
	return
}

func (db *Database) VersionsCount(query ...interface{}) (
	count int64, caught try.Err) {

	defer try.Annotate(&caught, "error counting versions")

	if len(query) > 0 {
		tryORM(db.orm.Model(Version{}).Where(query[0], query[1:]...).Count(&count))
	} else {
		tryORM(db.orm.Model(Version{}).Count(&count))
	}
	return
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

	if try.Int64(db.VersionsCount(
		"name = ? AND id <> ? AND website_id = ?", v.Name, v.ID, wsid)) > 0 {
		return newErrVersionNameAlreadyExists(v.Name, wsid)
	}

	return nil
}
