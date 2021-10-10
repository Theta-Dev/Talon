package util

import "errors"

var (
	ErrUnknownSqlDialect = errors.New("unknown dialect (allowed: sqlite, mysql, postgres)")
	ErrWrongUserPass     = errors.New("username/password wrong")
)
