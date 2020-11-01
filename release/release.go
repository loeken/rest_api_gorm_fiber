package release

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/loeken/rest_api_gorm_fiber/database"
)

type Release struct {
	gorm.Model
	Title string `json:"title"`
	Artist string `json:"artist"`
}

func GetReleases(c *fiber.Ctx) {
	db := database.DBConn
	var releases []Release
	db.Find(&releases)
	c.JSON(releases)
}
func GetRelease(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var release  Release
	db.Find(&release, id)
	c.JSON(release)
}
func NewRelease(c *fiber.Ctx) {
	db := database.DBConn
	
	release := new (Release)
	if err := c.BodyParser(release); err != nil {
		c.Status(503).Send(err)
	}
	db.Create(&release)
	c.JSON(release) 
}
func DeleteRelease(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var release Release
	db.First(&release, id)
	if release.Title == "" {
		c.Status(500).Send("no release found with given id")
		return
	}
	db.Delete(&release)
	c.Send("release successfully deleted")
}
func UpdateRelease(c *fiber.Ctx) {
	c.Send("Updates Release")
}