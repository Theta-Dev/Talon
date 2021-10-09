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

func Open(conn *Connection) (db Database, tryErr error) {
	defer try.Returnf(&tryErr, "error opening db")

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

func (db Database) Migrate() (tryErr error) {
	defer try.Returnf(&tryErr, "migration error")

	try.Check(db.orm.AutoMigrate(allModels...))
	return
}

func (db Database) DropAllTables() (tryErr error) {
	defer try.Returnf(&tryErr, "drop tables error")

	var queries []string

	switch db.conn.Dialect {
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
		try.ORM(db.orm.Exec(q))
	}
	return
}

func (db Database) EmptyAllTables() (tryErr error) {
	defer try.Returnf(&tryErr, "empty tables error")

	var queries []string

	for _, n := range tableNames {
		queries = append(queries, "DELETE FROM "+n)

		if db.conn.Dialect != DialectSqlite {
			queries = append(queries, "ALTER TABLE "+n+" AUTO_INCREMENT=1")
		}
	}

	for _, q := range queries {
		try.ORM(db.orm.Exec(q))
	}
	return
}
