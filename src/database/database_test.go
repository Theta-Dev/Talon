package database

import (
	"testing"

	"github.com/Theta-Dev/Talon/src/try"
	"gorm.io/gorm/logger"
)

func TestNew_sqlite(t *testing.T) {
	conn := Connection{
		Dialect: DialectSqlite,
		File:    "",
	}

	db := try.X(New(conn)).(Database)
	db.dropAllTables()

	db.orm.Logger.LogMode(logger.Info)
	try.Check(db.Migrate())
}

func TestNew_mariadb(t *testing.T) {
	conn := Connection{
		Dialect: DialectMySql,
		DbName:  "talon",
		User:    "test",
		Pass:    "=bLa@23-x=",
	}

	db := try.X(New(conn)).(Database)
	db.dropAllTables()

	db.orm.Logger.LogMode(logger.Info)
	try.Check(db.Migrate())
}

func TestNew_postgres(t *testing.T) {
	conn := Connection{
		Dialect: DialectPostgres,
		DbName:  "talon",
		User:    "test",
		Pass:    "=bLa@23-x=",
	}

	db := try.X(New(conn)).(Database)
	db.dropAllTables()

	db.orm.Logger.LogMode(logger.Info)
	try.Check(db.Migrate())
}
