package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateTeacher(data map[string]string, c *fiber.Ctx) interface{} {
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	teacher := models.Teacher{
		Type:       data["type"],
		Name:       data["name"],
		Login:      data["login"],
		Lastname:   data["lastname"],
		Pantonymic: data["pantonymic"],
		Password:   password,
	}
	database.DB.Create(&teacher)
	return c.JSON(teacher)
}

func TeacherSearch(login string) uint {
	var teacher models.Teacher
	database.DB.Where("login = ?", login).Find(&teacher)
	if teacher.Id == 0 {
		return 0
	}
	return teacher.Id
}
