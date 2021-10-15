package database

import (
	"database/sql"
	"log"
	"reflect"
	"time"

	"github.com/Theta-Dev/Talon/src/try"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	conn *Connection
	orm  *gorm.DB
}

func Open(conn *Connection, lgr *log.Logger) (db *Database, caught error) {
	defer try.Returnf(&caught, "error opening db")

	db = &Database{conn: conn}

	cfg := gorm.Config{
		Logger: logger.New(
			lgr,
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	}

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

func (db *Database) Migrate(printCmds bool) (caught error) {
	defer try.Returnf(&caught, "migration error")

	oldLogger := db.orm.Config.Logger

	if printCmds {
		db.orm.Config.Logger = db.orm.Config.Logger.LogMode(logger.Info)
	}

	try.Check(db.orm.AutoMigrate(AllModels...))

	db.orm.Config.Logger = oldLogger
	return
}

func (db *Database) GetDialect() string {
	return db.conn.Dialect
}

func (db *Database) IsDialect(dialect string) bool {
	return db.conn.Dialect == dialect
}

func (db *Database) Exec(sql string, values ...interface{}) (rows int, err error) {
	res := db.orm.Exec(sql, values...)
	return int(res.RowsAffected), res.Error
}

func (db *Database) GetVersion() (version string, caught error) {
	defer try.Returnf(&caught, "error getting version")

	var query, v string

	if db.IsDialect(DialectSqlite) {
		query = "SELECT sqlite_version()"
	} else {
		query = "SELECT version()"
	}

	db.orm.Raw(query).Scan(&v)
	return v, nil
}

func getRelId(idParam uint, sctParam interface{}) uint {
	if idParam > 0 {
		return idParam
	}

	idFromStruct := func() uint {
		defer func() {
			_ = recover()
		}()

		sct := reflect.ValueOf(sctParam)
		if sct.Kind() == reflect.Ptr {
			sct = sct.Elem()
		}
		id := sct.FieldByName("ID").Uint()
		return uint(id)
	}()

	return idFromStruct
}

func isRelSet(idParam uint, sctParam interface{}) bool {
	return getRelId(idParam, sctParam) > 0
}
