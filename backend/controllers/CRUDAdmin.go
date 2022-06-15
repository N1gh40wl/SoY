package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(data map[string]string, c *fiber.Ctx) interface{} {
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	admin := models.Admin{
		Type:     data["type"],
		Name:     data["name"],
		Login:    data["login"],
		Password: password,
	}
	database.DB.Create(&admin)
	return c.JSON(admin)
}

func Create(c *fiber.Ctx) error {
	password, _ := bcrypt.GenerateFromPassword([]byte("www"), 14)
	database.DB.Create(&models.Admin{Name: "www", Login: "www", Password: password, Timetable: models.Timetable{Monday: "1", Tuesday: "2", Wednesday: "3", Thursday: "4", Friday: "5"}})
	return nil
}

func AdminSearch(login string) uint {
	var admin models.Admin
	database.DB.Where("login = ?", login).Find(&admin)
	if admin.Id == 0 {
		return 0
	}
	return admin.Id
}
