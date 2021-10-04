package database

import (
	"fmt"
	"regexp"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DialectSqlite   = "sqlite"
	DialectMySql    = "mysql"
	DialectPostgres = "postgres"
)

type Connection struct {
	Dialect string `def:"sqlite"`
	File    string `def:"database.db"`
	Host    string `def:"127.0.0.1"`
	User    string `def:"user"`
	Pass    string `def:"pass"`
	DbName  string `def:"db"`
}

func (c Connection) getDialect() string {
	if c.Dialect != DialectSqlite && c.Dialect != DialectMySql && c.Dialect != DialectPostgres {
		return DialectSqlite
	}
	return c.Dialect
}

func (c Connection) getFile() string {
	if c.File == "" {
		return "database.db"
	}
	return c.File
}

func (c Connection) getHost() string {
	if c.Host == "" {
		return "127.0.0.1"
	}
	return c.Host
}
func (c Connection) getUser() string {
	if c.User == "" {
		return "user"
	}
	return c.User
}
func (c Connection) getPass() string {
	if c.Pass == "" {
		return "pass"
	}
	return c.Pass
}

func (c Connection) getDbName() string {
	if c.DbName == "" {
		return "db"
	}
	return c.DbName
}

func splitHostUrl(url string, defaultPort string) (host string, port string) {
	pattern := regexp.MustCompile(`(.*):(\d+)`)
	match := pattern.FindStringSubmatch(url)

	if match == nil {
		return url, defaultPort
	}

	return match[1], match[2]
}

func (c Connection) getDsn() string {
	switch c.getDialect() {
	case DialectSqlite:
		return c.getFile()
	case DialectMySql:
		host, port := splitHostUrl(c.getHost(), "3306")

		return fmt.Sprintf(`%s:"%s"@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
			c.getUser(), c.getPass(), host, port, c.getDbName())
	case DialectPostgres:
		host, port := splitHostUrl(c.getHost(), "5432")

		return fmt.Sprintf(`host=%s user=%s password="%s" dbname=%s port=%s sslmode=disable`,
			host, c.getUser(), c.getPass(), c.getDbName(), port)
	}
	return ""
}

func (c Connection) Open() gorm.Dialector {
	dsn := c.getDsn()

	switch c.getDialect() {
	case DialectSqlite:
		return sqlite.Open(dsn)
	case DialectMySql:
		return mysql.Open(dsn)
	case DialectPostgres:
		return postgres.Open(dsn)
	}
	return nil
}
