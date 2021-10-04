package main

func main() {
	/*
	app := fiber.New()

	app.Use(compress.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Static("/assets", "./assets")

	app.Listen(":8000")*/

	/*
	conn := database.Connection{
		Dialect: database.DialectMySql,
		User: "test",
		Pass: "1234",
		DbName: "talon",
	}

	db, err := gorm.Open(conn.Open(), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}


	db.AutoMigrate(&model.Website{})
	db.AutoMigrate(&model.Version{})
	db.AutoMigrate(&model.VersionFile{})
	db.AutoMigrate(&model.File{})

	ws1 := model.Website{
		Name: "Talon",
		Path: "talon",
		Color: sql.NullString{String: "#7935df"},
		Visibility: "featured",
	}
	result := db.Create(&ws1)
	fmt.Printf("%d rows affected", result.RowsAffected)*/
}
