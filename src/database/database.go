package database

import (
	"database/sql"
	"os"
	"path/filepath"
	"time"

	"github.com/Theta-Dev/Talon/src/try"
	"gorm.io/gorm"
)

type Database struct {
	conn Connection
	orm  *gorm.DB
}

func New(conn Connection) (db Database, err error) {
	defer try.Returnf(&err, "error opening db")

	db = Database{conn: conn}

	try.Check(db.createDirectoryIfSqlite())

	cfg := gorm.Config{}
	db.orm = try.DB(gorm.Open(conn.Open(), &cfg))

	sqldb := try.X(db.orm.DB()).(*sql.DB)

	if conn.getDialect() == DialectSqlite {
		sqldb.SetMaxOpenConns(1)
		db.orm.Exec("PRAGMA foreign_keys = ON")
	} else {
		sqldb.SetMaxOpenConns(10)
	}

	if conn.getDialect() == DialectMySql {
		// Mysql has a setting called wait_timeout, which defines the duration
		// after which a connection may not be used anymore.
		// The default for this setting on mariadb is 10 minutes.
		// See https://github.com/docker-library/mariadb/issues/113
		sqldb.SetConnMaxLifetime(9 * time.Minute)
	}
	return
}

func (db Database) Migrate() (err error) {
	defer try.Returnf(&err, "migration error")

	if db.conn.getDialect() != DialectSqlite {
		db.orm.Config.DisableForeignKeyConstraintWhenMigrating = true
		try.Check(db.orm.AutoMigrate(allModels...))
		db.orm.Config.DisableForeignKeyConstraintWhenMigrating = false
	}

	try.Check(db.orm.AutoMigrate(allModels...))
	return
}

func (db Database) createDirectoryIfSqlite() (err error) {
	defer try.Returnf(&err, "error creating sqlite dir")

	if db.conn.getDialect() == DialectSqlite {
		if _, err := os.Stat(filepath.Dir(db.conn.File)); os.IsNotExist(err) {
			try.Check(os.MkdirAll(filepath.Dir(db.conn.File), 0777))
		}
	}
	return
}

func (db Database) checkErr() {
	try.Check(db.orm.Error)
}

func (db Database) dropAllTables() (err error) {
	defer try.Returnf(&err, "drop tables error")

	var queries []string

	switch db.conn.getDialect() {
	case DialectMySql:
		queries = []string{"SET FOREIGN_KEY_CHECKS = 0"}

		for _, n := range tableNames {
			queries = append(queries, "DROP TABLE IF EXISTS "+n)
		}
		queries = append(queries, "SET FOREIGN_KEY_CHECKS = 1")
	case DialectPostgres:
		queries = []string{
			"DROP SCHEMA public CASCADE",
			"CREATE SCHEMA public",
			"GRANT ALL ON SCHEMA public TO test",
		}
	case DialectSqlite:
		queries = []string{"PRAGMA foreign_keys = OFF"}

		for _, n := range tableNames {
			queries = append(queries, "DROP TABLE IF EXISTS "+n)
		}
		queries = append(queries, "PRAGMA foreign_keys = ON")
	}

	for _, q := range queries {
		db.orm.Exec(q)
		db.checkErr()
	}
	return
}
