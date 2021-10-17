package database

import (
	"fmt"
	"strings"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/util"
	"gorm.io/gorm"
)

func (db *Database) WebsiteAdd(website *Website) (caught try.Err) {
	defer try.Annotate(&caught, "error adding website")

	website.ID = 0
	try.Check(website.check(db))
	tryORM(db.orm.Create(&website))
	return
}

func (db *Database) WebsiteUpdate(website *Website) (caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error updating website %d", website.ID))

	try.Check(website.check(db))
	tryORM(db.orm.Save(&website))
	return
}

func (db *Database) WebsiteByID(id uint, deep bool) (website *Website, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting website %d", id))

	orm := db.orm
	if deep {
		orm = orm.Scopes(websiteFetchScope)
	}

	if tryORMIsEmpty(orm.First(&website, id)) {
		return nil, nil
	}
	return
}

func (db *Database) WebsiteByPath(sitePath string, deep bool) (
	website *Website, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting website at path %s", sitePath))

	normPath := util.NormalizePath(sitePath)

	orm := db.orm
	if deep {
		orm = orm.Scopes(websiteFetchScope)
	}

	if tryORMIsEmpty(orm.Where("path_lower = ?", normPath).First(&website)) {
		return nil, nil
	}
	return
}

func (db *Database) WebsitePathExists(sitePath string) (exists bool, caught try.Err) {
	defer try.Annotate(&caught, "error checking website path")

	normPath := util.NormalizePath(sitePath)

	c := try.Int(db.WebsitesCount("path_lower = ?", normPath))
	return c > 0, nil
}

func (db *Database) WebsitesGet(deep bool, query ...interface{}) (
	websites []*Website, caught try.Err) {

	defer try.Annotate(&caught, "error getting websites")

	orm := db.orm
	if deep {
		orm = orm.Scopes(websiteFetchScope)
	}

	if len(query) > 0 {
		tryORMIsEmpty(orm.Where(query[0], query[1:]...).Find(&websites))
	} else {
		tryORMIsEmpty(orm.Find(&websites))
	}
	return
}

func (db *Database) WebsitesCount(query ...interface{}) (
	count int, caught try.Err) {

	defer try.Annotate(&caught, "error counting websites")

	var c int64
	if len(query) > 0 {
		tryORM(db.orm.Model(Website{}).Where(query[0], query[1:]...).Count(&c))
	} else {
		tryORM(db.orm.Model(Website{}).Count(&c))
	}
	return int(c), nil
}

func (db *Database) WebsiteDeleteByID(id uint) (caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error deleting website %d", id))

	tryORM(db.orm.Delete(&Website{}, id))
	return
}

func websiteFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("Logo").Joins("User").Preload("Versions")
}

func (w *Website) check(db *Database) try.Err {
	w.Path = util.TrimPath(w.Path)
	w.PathLower = strings.ToLower(w.Path)

	if !isRelSet(w.UserID, w.User) {
		return try.FromErr(ErrEmptyUser)
	}

	if try.Int(db.WebsitesCount("path_lower = ? AND id <> ?", w.PathLower, w.ID)) > 0 {
		return newErrSitePathAlreadyExists(w.PathLower)
	}
	return nil
}
