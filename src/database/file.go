package database

import (
	"fmt"

	"code.thetadev.de/ThetaDev/gotry/try"
)

func (db *Database) FileAdd(file *File) (caught try.Err) {
	defer try.Annotate(&caught, "error adding file")

	if try.Bool(db.FileHashExists(file.Hash)) {
		try.Check(newErrFileHashAlreadyExists(file.Hash))
	}

	file.ID = 0
	tryORM(db.orm.Create(&file))
	return
}

func (db *Database) FileByID(id uint) (file *File, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting file %d", id))

	if tryORMIsEmpty(db.orm.First(&file, id)) {
		return nil, nil
	}
	return
}

func (db *Database) FileByHash(hash string) (file *File, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error getting file with hash %s", hash))

	if tryORMIsEmpty(db.orm.Where("hash = ?", hash).First(&file)) {
		return nil, nil
	}
	return
}

func (db *Database) FileHashExists(hash string) (exists bool, caught try.Err) {
	defer try.Annotate(&caught, fmt.Sprintf("error checking file hash %s", hash))

	c := try.Int64(db.FilesCount("hash = ?", hash))
	return c > 0, nil
}

func (db *Database) FilesGet(query ...interface{}) (files []*File, caught try.Err) {
	defer try.Annotate(&caught, "error getting files")

	if len(query) > 0 {
		tryORMIsEmpty(db.orm.Where(query[0], query[1:]...).Find(&files))
	} else {
		tryORMIsEmpty(db.orm.Find(&files))
	}
	return
}

func (db *Database) FilesCount(query ...interface{}) (count int64, caught try.Err) {
	defer try.Annotate(&caught, "error counting files")

	if len(query) > 0 {
		tryORM(db.orm.Model(File{}).Where(query[0], query[1:]...).Count(&count))
	} else {
		tryORM(db.orm.Model(File{}).Count(&count))
	}
	return
}

func (db *Database) FilesGetOrphans() (files []*File, caught try.Err) {
	defer try.Annotate(&caught, "error getting orphans")

	var f []*File
	tryORM(db.orm.Raw("SELECT f.id, f.hash FROM files f LEFT JOIN version_files vf " +
		"ON f.id = vf.file_id LEFT JOIN websites w ON f.id = w.logo_id " +
		"WHERE vf.id IS NULL AND w.id IS NULL").Scan(&f))

	return f, nil
}

func (db *Database) FilesDeleteOrphans() (caught try.Err) {
	defer try.Annotate(&caught, "error deleting orphans")

	tryORM(db.orm.Exec("DELETE FROM files WHERE id IN (" +
		"SELECT f.id FROM files f LEFT JOIN version_files vf " +
		"ON f.id = vf.file_id LEFT JOIN websites w ON f.id = w.logo_id " +
		"WHERE vf.id IS NULL AND w.id IS NULL )"))

	return
}
