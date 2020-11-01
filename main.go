package main


import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/loeken/rest_api_gorm_fiber/release"
	"github.com/loeken/rest_api_gorm_fiber/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World");
}
func setupRoutes(app *fiber.App) {

	app.Get("/api/v1/release", release.GetReleases)
	app.Get("/api/v1/release/:id", release.GetRelease)
	app.Post("/api/v1/release", release.NewRelease)
	app.Delete("/api/v1/release/:id", release.DeleteRelease)
	app.Put("/api/v1/release/:id", release.UpdateRelease)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "releases.db")
	if err != nil  {
		panic("Failed to connect to database")
	}
	fmt.Println("database connection succefully opened")

	database.DBConn.AutoMigrate(&release.Release{})
	fmt.Println("Database migrated")
}
func main() {
	app := fiber.New()
	
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)


	app.Listen(3000)
}