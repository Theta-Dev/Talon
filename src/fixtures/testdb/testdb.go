package testdb

import (
	"log"
	"os"
	"path/filepath"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures"
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
		Path:    "assets/image.jpg",
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
	// 0. ThetaDev v0.1.0 index.html
	{Hash: "3b5f6bad5376897435def176d0fe77e5b9b4f0deafc7491fc27262650744ad68"},
	// 1. ThetaDev style.css
	{Hash: "356f131c825fbf604797c7e9c85352549d81db8af91fee834016d075110af026"},
	// 2. ThetaDev Logo
	{Hash: "9c37a2cb1230a9cbe7911d34404d4fb03b27552e56b2173683cf9fc52be7bc99"},
	// 3. ThetaDev v0.1.1 index.html
	{Hash: "a44816e6c3b650bdf88e6532659ba07ef187c2113ae311da9709e056aec8eadb"},
	// 4. ThetaDev v0.1.1 image.jpg
	{Hash: "901d291a47a8a9b55c06f84e5e5f82fd2dcee65cac1406d6e878b805d45c1e93"},
	// 5. ThetaDev v0.1.1 test.js
	{Hash: "b6ed35f5ae339a35a8babb11a91ff90c1a62ef250d30fa98e59500e8dbb896fa"},
	// 6. ThetaDev v0.1.1 example.txt
	{Hash: "bae6bdae8097c24f9a99028e04bfc8d5e0a0c318955316db0e7b955def9c1dbb"},
	// 7. Talon index.html
	{Hash: "6e99e2dbab5524a692616f208e7fbc54778dd67af6968468393a9912f51f707d"},
	// 8. Talon talon_style.css
	{Hash: "3362eb0f20a047541087bdb80f3932e6cb8234c418a3d91f316be2194191945e"},
	// 9. Talon Logo
	{Hash: "65356c31d40537b7e2ed6cc21b7070ee9440bb3fcb56b5866f42a58622b64c00"},
	// 10. GenderEx index.html
	{Hash: "6c5d37546616519e8973be51515b8a90898b4675f7b6d01f2d891edb686408a2"},
	// 11. GenderEx gex_style.css
	{Hash: "fc825b409a49724af8f5b3c4ad15e175e68095ea746237a7b46152d3f383f541"},
	// 12. GenderEx Logo
	{Hash: "9b35024aacebd74010ea595ef5d180f47f5ec822df100236dd6ac808497b64f6"},
	// Orphans
	{Hash: "ddc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
	{Hash: "edc712d4866413b299f80ca7a422df2a1afe278a5ee005dc9229f37dd8fbe8c6"},
}

func open() *database.Database {
	dialect := os.Getenv("DIALECT")
	dbfile := filepath.Join("talon_tmp", "test.db")

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

	return db
}

func Open() *database.Database {
	db := open()
	insertTestData(db)

	return db
}

func OpenWithoutFiles() *database.Database {
	db := open()
	insertTestDataWithoutFiles(db)

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

func insertTestData(db *database.Database) {
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

func insertTestDataWithoutFiles(db *database.Database) {
	for _, u := range Users {
		try.Check(db.UserAdd(u))
	}

	for _, w := range Websites {
		w.Logo = nil
		try.Check(db.WebsiteAdd(w))
	}

	for _, v := range Versions {
		try.Check(db.VersionAdd(v))
	}
}
