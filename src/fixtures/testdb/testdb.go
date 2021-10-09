package testdb

import (
	"os"
	"path"

	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures"
	"github.com/Theta-Dev/Talon/src/try"
)

func Open() database.Database {
	dialect := os.Getenv("DIALECT")
	dbfile := path.Join("tmp", "test.db")

	if dialect == "" || dialect == database.DialectSqlite {
		fixtures.CdProjectRoot()
		_ = os.Remove(dbfile)
	}

	conn := &database.Connection{
		Dialect: dialect,
		File:    dbfile,
		DbName:  "talon",
		User:    "test",
		Pass:    "1234",
	}

	db := try.X(database.Open(conn)).(database.Database)

	try.Check(db.Migrate())
	EmptyAllTables(db)
	InsertTestData(db)

	return db
}

func DropAllTables(db database.Database) {
	var queries []string

	switch db.GetDialect() {
	case database.DialectMySql:
		queries = []string{"SET FOREIGN_KEY_CHECKS = 0"}

		for _, n := range database.TableNames {
			queries = append(queries, "DROP TABLE IF EXISTS "+n)
		}
		queries = append(queries, "SET FOREIGN_KEY_CHECKS = 1")
	case database.DialectPostgres:
		queries = []string{
			"DROP SCHEMA public CASCADE",
			"CREATE SCHEMA public",
			"GRANT ALL ON SCHEMA public TO test",
		}
	case database.DialectSqlite:
		queries = []string{"PRAGMA foreign_keys = OFF"}

		for _, n := range database.TableNames {
			queries = append(queries, "DROP TABLE IF EXISTS "+n)
		}
		queries = append(queries, "PRAGMA foreign_keys = ON")
	}

	for _, q := range queries {
		try.Int(db.Exec(q))
	}
}

func EmptyAllTables(db database.Database) {
	var queries []string

	for _, n := range database.TableNames {
		queries = append(queries, "DELETE FROM "+n)

		if !db.IsDialect(database.DialectSqlite) {
			queries = append(queries, "ALTER TABLE "+n+" AUTO_INCREMENT=1")
		}
	}

	for _, q := range queries {
		try.Int(db.Exec(q))
	}
}

func InsertTestData(db database.Database) {
	try.X(db.AddUser("ThetaDev", "1234", &database.Permission{
		AllowedPaths: "#",
		IsAdmin:      true,
		CanCreate:    true,
	}))

	try.X(db.AddUser("Zoey", "5678", &database.Permission{
		AllowedPaths: "Talon/#",
		CanCreate:    true,
	}))

	try.X(db.AddUser("Izzy", "2020", &database.Permission{
		AllowedPaths: "Talon/#",
	}))
}
