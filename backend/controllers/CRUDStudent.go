package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateStudent(data map[string]string, c *fiber.Ctx) interface{} {
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	student := models.Student{
		Type:       data["type"],
		Name:       data["name"],
		Login:      data["login"],
		Lastname:   data["lastname"],
		Pantonymic: data["pantonymic"],
		Password:   password,
	}
	database.DB.Create(&student)
	return c.JSON(student)
}

func StudentSearch(login string) uint {
	var student models.Student
	database.DB.Where("login = ?", login).Find(&student)
	if student.Id == 0 {
		return 0
	}
	return student.Id
}

func StudentClass(studentId uint) uint {
	var student models.Student
	database.DB.Where("id = ?", studentId).Find(&student)
	return student.GroupId
}
