package test_database

import (
	"os"
	"path"

	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/try"
	"github.com/Theta-Dev/Talon/tests"
)

func OpenTestDB() database.Database {
	dialect := os.Getenv("DIALECT")
	dbfile := path.Join("tmp", "test.db")

	if dialect == "" || dialect == database.DialectSqlite {
		tests.CdProjectRoot()
		os.Remove(dbfile)
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
	try.Check(db.EmptyAllTables())
	try.Check(insertTestData(db))

	return db
}

func insertTestData(db database.Database) (tryErr error) {
	defer try.Returnf(&tryErr, "error inserting test data")

	db.AddUser("ThetaDev", "1234", database.Permission{
		AllowedPaths: "#",
		IsAdmin:      true,
		CanCreate:    true,
	})

	db.AddUser("Zoey", "5678", database.Permission{
		AllowedPaths: "Talon/#",
		CanCreate:    true,
	})

	db.AddUser("Izzy", "2020", database.Permission{
		AllowedPaths: "Talon/#",
	})

	return
}
