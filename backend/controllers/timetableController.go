package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func CreateTimetable() {

	database.DB.Create(&models.Timetable{Monday: "1", Tuesday: "2", Wednesday: "3", Thursday: "4", Friday: "5"})

}

func TimetableShow(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if AdminSearch(claims.Issuer) != 0 {
		var timetable models.Timetable
		database.DB.Where("owner_id = ? AND owner_type = ?", AdminSearch(claims.Issuer), "admins").First(&timetable)
		return c.JSON(timetable)
	} else if TeacherSearch(claims.Issuer) != 0 {
		var timetable models.Timetable
		database.DB.Where("owner_id = ? AND owner_type = ?", TeacherSearch(claims.Issuer), "teachers").First(&timetable)
		return c.JSON(timetable)
	} else if StudentSearch(claims.Issuer) != 0 {
		var timetable models.Timetable
		database.DB.Where("owner_id = ? AND owner_type = ?", StudentSearch(claims.Issuer), "students").First(&timetable)
		return c.JSON(timetable)
	}
	return nil
}
