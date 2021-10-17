package database

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/util"
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
	Dialect string
	File    string
	Host    string
	User    string
	Pass    string
	DbName  string
}

func splitHostUrl(url string, defaultPort string) (host string, port string) {
	pattern := regexp.MustCompile(`(.*):(\d+)`)
	match := pattern.FindStringSubmatch(url)

	if match == nil {
		return url, defaultPort
	}

	return match[1], match[2]
}

func (c *Connection) prepare() (caught try.Err) {
	defer try.Annotate(&caught, "error with connection data")

	c.Dialect = strings.ToLower(c.Dialect)

	if c.Dialect == "" {
		c.Dialect = DialectSqlite
	} else if c.Dialect != DialectSqlite &&
		c.Dialect != DialectMySql && c.Dialect != DialectPostgres {
		return newErrUnknownSqlDialect(c.Dialect)
	}

	if c.Dialect == DialectSqlite {
		try.Check(c.prepareFileDB())
	} else {
		try.Check(c.prepareExternalDB())
	}
	return
}

func (c *Connection) prepareFileDB() (caught try.Err) {
	defer try.Return(&caught)

	if c.File == "" {
		c.File = "database.db"
	}

	try.Check(util.CreateDirIfNotExists(filepath.Dir(c.File)))

	c.Host = ""
	c.User = ""
	c.Pass = ""
	c.DbName = filepath.Base(c.File)

	return
}

func (c *Connection) prepareExternalDB() error {
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.User == "" {
		return ErrEmptyDbUsername
	}
	if c.Pass == "" {
		return ErrEmptyDbPassword
	}
	if c.DbName == "" {
		return ErrEmptyDbName
	}

	c.File = ""

	return nil
}

func (c *Connection) getDsn() string {
	switch c.Dialect {
	case DialectSqlite:
		return c.File
	case DialectMySql:
		host, port := splitHostUrl(c.Host, "3306")

		return fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
			c.User, c.Pass, host, port, c.DbName)

	case DialectPostgres:
		host, port := splitHostUrl(c.Host, "5432")

		return fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable`,
			host, c.User, c.Pass, c.DbName, port)
	}
	return ""
}

func (c *Connection) Open() (d gorm.Dialector, caught try.Err) {
	defer try.Return(&caught)

	try.Check(c.prepare())
	dsn := c.getDsn()

	switch c.Dialect {
	case DialectSqlite:
		d = sqlite.Open(dsn)
	case DialectMySql:
		d = mysql.Open(dsn)
	case DialectPostgres:
		d = postgres.Open(dsn)
	}
	return
}
