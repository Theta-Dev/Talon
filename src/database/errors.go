package database

import (
	"errors"
	"fmt"

	"code.thetadev.de/ThetaDev/gotry/try"
)

var (
	ErrEmptyDbUsername   = errors.New("empty db username")
	ErrEmptyDbPassword   = errors.New("empty db password")
	ErrEmptyDbName       = errors.New("empty db name")
	ErrUnknownSqlDialect = errors.New(
		"unknown dialect (allowed: sqlite, mysql, postgres)")
	ErrUsernameAlreadyExists    = errors.New("username already exists")
	ErrSitePathAlreadyExists    = errors.New("website path already exists")
	ErrVersionNameAlreadyExists = errors.New("version name already exists in website")
	ErrVersionFileAlreadyExists = errors.New("version file already exists in version")
	ErrEmptyUser                = errors.New("user empty")
	ErrEmptyWebsite             = errors.New("website empty")
	ErrEmptyVersion             = errors.New("version empty")
	ErrEmptyFile                = errors.New("file empty")
)

func newErrUnknownSqlDialect(dialect string) try.Err {
	err := try.FromErr(ErrUnknownSqlDialect)
	err.Annotate("dialect: " + dialect)
	return err
}

func newErrUsernameAlreadyExists(username string) try.Err {
	err := try.FromErr(ErrUsernameAlreadyExists)
	err.Annotate("username: " + username)
	return err
}

func newErrSitePathAlreadyExists(path string) try.Err {
	err := try.FromErr(ErrSitePathAlreadyExists)
	err.Annotate("site path: " + path)
	return err
}

func newErrVersionNameAlreadyExists(version string, websiteId uint) try.Err {
	err := try.FromErr(ErrVersionNameAlreadyExists)
	err.Annotate(fmt.Sprintf("version %s in website %d", version, websiteId))
	return err
}

func newErrVersionFileAlreadyExists(filePath string, versionId uint) try.Err {
	err := try.FromErr(ErrVersionFileAlreadyExists)
	err.Annotate(fmt.Sprintf("file path %s in version %d", filePath, versionId))
	return err
}
