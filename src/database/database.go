package database

import (
	"database/sql"
	"time"

	"github.com/Theta-Dev/Talon/src/try"
	"gorm.io/gorm"
)

type Database struct {
	conn *Connection
	orm  *gorm.DB
}

func Open(conn *Connection) (db Database, caught error) {
	defer try.Returnf(&caught, "error opening db")

	db = Database{conn: conn}

	cfg := gorm.Config{}
	dialector := try.X(conn.Open()).(gorm.Dialector)
	db.orm = try.DB(gorm.Open(dialector, &cfg))

	sqldb := try.X(db.orm.DB()).(*sql.DB)

	if conn.Dialect == DialectSqlite {
		sqldb.SetMaxOpenConns(1)
		db.orm.Exec("PRAGMA foreign_keys = ON")
	} else {
		sqldb.SetMaxOpenConns(10)
	}

	if conn.Dialect == DialectMySql {
		// Mysql has a setting called wait_timeout, which defines the duration
		// after which a connection may not be used anymore.
		// The default for this setting on mariadb is 10 minutes.
		// See https://github.com/docker-library/mariadb/issues/113
		sqldb.SetConnMaxLifetime(9 * time.Minute)
	}
	return
}

func (db Database) Migrate() (caught error) {
	defer try.Returnf(&caught, "migration error")

	try.Check(db.orm.AutoMigrate(AllModels...))
	return
}

func (db Database) GetDialect() string {
	return db.conn.Dialect
}

func (db Database) IsDialect(dialect string) bool {
	return db.conn.Dialect == dialect
}

func (db Database) Exec(sql string, values ...interface{}) (rows int, err error) {
	res := db.orm.Exec(sql, values...)
	return int(res.RowsAffected), res.Error
}
