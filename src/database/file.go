package database

import (
	"fmt"

	"github.com/Theta-Dev/Talon/src/try"
)

func (db *Database) FileAdd(file *File) (caught error) {
	defer try.Returnf(&caught, "error adding file")

	if try.Bool(db.FileHashExists(file.Hash)) {
		try.Check(fmt.Errorf("file hash %s already exists", file.Hash))
	}

	file.ID = 0
	try.ORM(db.orm.Create(&file))
	return
}

func (db *Database) FileByID(id uint) (file *File, caught error) {
	defer try.Returnf(&caught, "error getting file %d", id)

	var f File
	if try.ORMIsEmpty(db.orm.First(&f, id)) {
		return nil, nil
	}
	return &f, nil
}

func (db *Database) FileByHash(hash string) (file *File, caught error) {
	defer try.Returnf(&caught, "error getting file with hash %s", hash)

	var f File
	if try.ORMIsEmpty(db.orm.Where("hash = ?", hash).First(&f)) {
		return nil, nil
	}
	return &f, nil
}

func (db *Database) FileHashExists(hash string) (exists bool, caught error) {
	defer try.Returnf(&caught, "error checking file hash %s", hash)

	c := try.Int(db.FilesCount("hash = ?", hash))
	return c > 0, nil
}

func (db *Database) FilesGet(query ...interface{}) (files []*File, caught error) {
	defer try.Returnf(&caught, "error getting files")

	var fs []*File
	if len(query) > 0 {
		try.ORM(db.orm.Where(query[0], query[1:]...).Find(&fs))
	} else {
		try.ORM(db.orm.Find(&fs))
	}
	return fs, nil
}

func (db *Database) FilesCount(query ...interface{}) (count int, caught error) {
	defer try.Returnf(&caught, "error counting files")

	var c int64
	if len(query) > 0 {
		try.ORM(db.orm.Model(File{}).Where(query[0], query[1:]...).Count(&c))
	} else {
		try.ORM(db.orm.Model(File{}).Count(&c))
	}
	return int(c), nil
}

func (db *Database) FilesGetOrphans() (files []*File, caught error) {
	defer try.Returnf(&caught, "error getting orphans")

	var f []*File
	try.ORM(db.orm.Raw("SELECT f.id, f.hash FROM files f LEFT JOIN version_files vf " +
		"ON f.id = vf.file_id LEFT JOIN websites w ON f.id = w.logo_id " +
		"WHERE vf.id IS NULL AND w.id IS NULL").Scan(&f))

	return f, nil
}

func (db *Database) FilesDeleteOrphans() (caught error) {
	defer try.Returnf(&caught, "error deleting orphans")

	try.ORM(db.orm.Exec("DELETE FROM files WHERE id IN (" +
		"SELECT f.id FROM files f LEFT JOIN version_files vf " +
		"ON f.id = vf.file_id LEFT JOIN websites w ON f.id = w.logo_id " +
		"WHERE vf.id IS NULL AND w.id IS NULL )"))

	return
}
