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

		if n != "talon_infos" {
			if db.IsDialect(database.DialectMySql) {
				queries = append(queries, "ALTER TABLE "+n+" AUTO_INCREMENT=1")
			} else if db.IsDialect(database.DialectPostgres) {
				queries = append(queries, "ALTER SEQUENCE "+n+"_id_seq RESTART WITH 1")
			}
		}
	}

	for _, q := range queries {
		try.Int(db.Exec(q))
	}
}

func InsertTestData(db database.Database) {
	try.Check(db.UserAdd(&database.User{
		Name:         "ThetaDev",
		PasswordHash: "$2a$10$psQvPDk7kDGBf5khGTohRuQajdmwGrY1OFb9c2b5pNiexuII.HMyO", // 1234
		Permission: &database.Permission{
			AllowedPaths: "#",
			IsAdmin:      true,
			CanCreate:    true,
		},
	}))

	try.Check(db.UserAdd(&database.User{
		Name:         "Zoey",
		PasswordHash: "$2a$10$732AemL9NzKCqT/QrJvpx.3UD/v/YmdM9aY.YjohgmgzAB70k0Jx6", // 5678
		Permission: &database.Permission{
			AllowedPaths: "Talon/#",
			CanCreate:    true,
		},
	}))

	try.Check(db.UserAdd(&database.User{
		Name:         "Izzy",
		PasswordHash: "$2a$10$AsSei6htRq68e4U3x3sni.QwMUTtYIrz7qgdEKX2nm79.Or8HtIii", // 2020
		Permission: &database.Permission{
			AllowedPaths: "Talon",
		},
	}))
}
