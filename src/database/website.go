package database

import (
	"fmt"
	"strings"

	"github.com/Theta-Dev/Talon/src/try"
	"github.com/Theta-Dev/Talon/src/util"
	"gorm.io/gorm"
)

func (db *Database) WebsiteAdd(website *Website) (caught error) {
	defer try.Returnf(&caught, "error adding website")

	website.ID = 0
	try.Check(website.check(db))
	try.ORM(db.orm.Create(&website))
	return
}

func (db *Database) WebsiteUpdate(website *Website) (caught error) {
	defer try.Returnf(&caught, "error updating website %d", website.ID)

	try.Check(website.check(db))
	try.ORM(db.orm.Save(&website))
	return
}

func (db *Database) WebsiteByID(id uint) (website *Website, caught error) {
	defer try.Returnf(&caught, "error getting website %d", id)

	var w Website
	if try.ORMIsEmpty(db.orm.Scopes(websiteFetchScope).First(&w, id)) {
		return nil, nil
	}
	return &w, nil
}

func (db *Database) WebsiteByPath(sitePath string) (website *Website, caught error) {
	defer try.Returnf(&caught, "error getting website at path %s", sitePath)

	normPath := util.NormalizePath(sitePath)

	var w Website
	if try.ORMIsEmpty(db.orm.Scopes(websiteFetchScope).Where(
		"path_lower = ?", normPath).First(&w)) {
		return nil, nil
	}
	return &w, nil
}

func (db *Database) WebsitePathExists(sitePath string) (exists bool, caught error) {
	defer try.Returnf(&caught, "error checking website path")

	normPath := util.NormalizePath(sitePath)

	c := try.Int(db.WebsitesCount("path_lower = ?", normPath))
	return c > 0, nil
}

func (db *Database) WebsitesGet(query ...interface{}) (
	websites []*Website, caught error) {

	defer try.Returnf(&caught, "error getting websites")

	var ws []*Website
	if len(query) > 0 {
		try.ORM(db.orm.Scopes(websiteFetchScope).Where(query[0], query[1:]...).Find(&ws))
	} else {
		try.ORM(db.orm.Scopes(websiteFetchScope).Find(&ws))
	}
	return ws, nil
}

func (db *Database) WebsitesCount(query ...interface{}) (
	count int, caught error) {

	defer try.Returnf(&caught, "error counting websites")

	var c int64
	if len(query) > 0 {
		try.ORM(db.orm.Model(Website{}).Where(query[0], query[1:]...).Count(&c))
	} else {
		try.ORM(db.orm.Model(Website{}).Count(&c))
	}
	return int(c), nil
}

func (db *Database) WebsiteDeleteByID(id uint) (caught error) {
	defer try.Returnf(&caught, "error deleting website %d", id)

	try.ORM(db.orm.Delete(&Website{}, id))
	return
}

func websiteFetchScope(db *gorm.DB) *gorm.DB {
	return db.Joins("Logo").Joins("User").Preload("Versions")
}

func (w *Website) check(db *Database) error {
	w.Path = strings.Trim(w.Path, "/")
	w.PathLower = strings.ToLower(w.Path)

	if !isRelSet(w.UserID, w.User) {
		return fmt.Errorf("no user")
	}

	if try.Int(db.WebsitesCount("path_lower = ? AND id <> ?", w.PathLower, w.ID)) > 0 {
		return fmt.Errorf("website path %s already exists", w.PathLower)
	}
	return nil
}
