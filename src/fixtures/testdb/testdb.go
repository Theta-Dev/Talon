package testdb

import (
	"log"
	"os"

	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures"
	"github.com/Theta-Dev/Talon/src/try"
)

var Users []*database.User = []*database.User{
	{
		Name:         "ThetaDev",
		PasswordHash: "$2a$10$psQvPDk7kDGBf5khGTohRuQajdmwGrY1OFb9c2b5pNiexuII.HMyO", // 1234
		Permission: &database.Permission{
			AllowedPaths: "#",
			IsAdmin:      true,
			CanCreate:    true,
		},
	},
	{
		Name:         "Zoey",
		PasswordHash: "$2a$10$732AemL9NzKCqT/QrJvpx.3UD/v/YmdM9aY.YjohgmgzAB70k0Jx6", // 5678
		Permission: &database.Permission{
			AllowedPaths: "Talon/#",
			CanCreate:    true,
		},
	},
	{
		Name:         "Izzy",
		PasswordHash: "$2a$10$AsSei6htRq68e4U3x3sni.QwMUTtYIrz7qgdEKX2nm79.Or8HtIii", // 2020
		Permission: &database.Permission{
			AllowedPaths: "tests/*",
		},
	},
}

var Websites []*database.Website = []*database.Website{
	{
		Name:       "ThetaDev",
		Path:       "",
		Logo:       Files[2],
		Color:      "#1f91ee",
		Visibility: 3,
		User:       Users[0],
		SourceUrl:  "https://github.com/Theta-Dev/Talon",
		SourceType: "github",
	},
	{
		Name:       "Talon",
		Path:       "Talon",
		Logo:       Files[4],
		Color:      "#7935df",
		Visibility: 3,
		User:       Users[1],
		SourceUrl:  "https://github.com/Theta-Dev/Talon",
		SourceType: "github",
	},
	{
		Name:       "Spotify-Gender-Ex",
		Path:       "Spotify-Gender-Ex",
		Logo:       Files[9],
		Color:      "#1DB954",
		Visibility: 2,
		User:       Users[0],
		SourceUrl:  "https://github.com/Theta-Dev/Talon",
		SourceType: "github",
	},
	{
		Name:       "TestA",
		Path:       "tests/A",
		Color:      "#ff0000",
		Visibility: 1,
		User:       Users[2],
	},
}

var Versions []*database.Version = []*database.Version{
	// ThetaDev
	{
		Name: "v0.1.0",
		Tags: `{"Deployed by": {"v": "GitHub Actions",` +
			`"u": "https://github.com/Theta-Dev/Talon/actions"}, "Build": {"v": "1"}}`,
		Website: Websites[0],
		User:    Users[0],
	},
	{
		Name: "v0.1.1",
		Tags: `{"Deployed by": {"v": "GitHub Actions",` +
			`"u": "https://github.com/Theta-Dev/Talon/actions"}, "Build": {"v": "1"}}`,
		Website: Websites[0],
		User:    Users[0],
	},
	// Talon
	{
		Name:    "#1",
		Website: Websites[1],
		User:    Users[1],
	},
	// Spotify-Gender-Ex
	{
		Name:    "#1",
		Website: Websites[2],
		User:    Users[0],
	},
}

var VersionFiles []*database.VersionFile = []*database.VersionFile{
	// ThetaDev v0.1.0
	{
		Path:    "index.html",
		Version: Versions[0],
		File:    Files[0],
	},
	{
		Path:    "style.css",
		Version: Versions[0],
		File:    Files[1],
	},
	// ThetaDev v0.1.1
	{
		Path:    "index.html",
		Version: Versions[1],
		File:    Files[5],
	},
	{
		Path:    "assets/style.css",
		Version: Versions[1],
		File:    Files[1],
	},
	{
		Path:    "assets/image.png",
		Version: Versions[1],
		File:    Files[10],
	},
	{
		Path:    "assets/test.js",
		Version: Versions[1],
		File:    Files[11],
	},
	{
		Path:    "data/example.txt",
		Version: Versions[1],
		File:    Files[12],
	},
	// Talon
	{
		Path:    "index.html",
		Version: Versions[2],
		File:    Files[3],
	},
	{
		Path:    "talon_style.css",
		Version: Versions[2],
		File:    Files[6],
	},
	// Spotify-Gender-Ex
	{
		Path:    "index.html",
		Version: Versions[3],
		File:    Files[7],
	},
	{
		Path:    "gex_style.css",
		Version: Versions[3],
		File:    Files[8],
	},
}

var Files []*database.File = []*database.File{
	// ThetaDev v0.1.0 index.html
	{Hash: "03806aba61268a03bcbd86998b76a8f7de12c010e5ec362057ad8dbef06075b7"},
	// ThetaDev style.css
	{Hash: "153f818f7297fd7d627331f6ab0a1713d53f66724b16653ef8c10ab912244e0f"},
	// ThetaDev Logo
	{Hash: "22c8011399c94d20527c0b85fd03441ad89885215d73b33af9e70ce5af130bc9"},
	// Talon index.html
	{Hash: "3dc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	// Talon Logo
	{Hash: "453b01efb5af334729659ce57df263c1c956996f3589aa1c21c803c13e2eb2b5"},
	// ThetaDev v0.1.1 index.html
	{Hash: "5336deebadf5cfe16195666916817cff73fd7032bf4b861ce0a34417b88e7dd4"},
	// Talon talon_style.css
	{Hash: "6dc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	// GenderEx index.html
	{Hash: "7dc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	// GenderEx gex_style.css
	{Hash: "8dc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	// GenderEx Logo
	{Hash: "9dc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	// ThetaDev v0.1.1 image.png
	{Hash: "adc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	// ThetaDev v0.1.1 test.js
	{Hash: "bdc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	// ThetaDev v0.1.1 example.txt
	{Hash: "cdc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	// Orphans
	{Hash: "ddc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	{Hash: "edc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
}

func Open() *database.Database {
	dialect := os.Getenv("DIALECT")
	dbfile := "test.db"

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

	lgr := log.New(os.Stdout, "\n", log.LstdFlags)
	db := try.X(database.Open(conn, lgr)).(*database.Database)

	try.Check(db.Migrate(false))

	if !db.IsDialect(database.DialectSqlite) {
		EmptyAllTables(db)
	}
	InsertTestData(db)

	return db
}

func DropAllTables(db *database.Database) {
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

func EmptyAllTables(db *database.Database) {
	var queries []string

	for _, n := range database.TableNames {
		queries = append(queries, "DELETE FROM "+n)

		if db.IsDialect(database.DialectMySql) {
			queries = append(queries, "ALTER TABLE "+n+" AUTO_INCREMENT=1")
		} else if db.IsDialect(database.DialectPostgres) {
			queries = append(queries, "ALTER SEQUENCE "+n+"_id_seq RESTART WITH 1")
		}
	}

	for _, q := range queries {
		try.Int(db.Exec(q))
	}
}

func InsertTestData(db *database.Database) {
	for _, u := range Users {
		try.Check(db.UserAdd(u))
	}

	for _, f := range Files {
		try.Check(db.FileAdd(f))
	}

	for _, w := range Websites {
		try.Check(db.WebsiteAdd(w))
	}

	for _, v := range Versions {
		try.Check(db.VersionAdd(v))
	}

	for _, vf := range VersionFiles {
		try.Check(db.VersionFileAdd(vf))
	}
}
