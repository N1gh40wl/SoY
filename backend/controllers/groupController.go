package controllers

import (
	"backend/database"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func IsTeacherGroup(teacher string, group string) bool {
	return false
}

func TeacherGroupShow(c *fiber.Ctx) error {
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

	if TeacherSearch(claims.Issuer) != 0 {
		teacherId := TeacherSearch(claims.Issuer)
		type Result struct {
			Name      string
			teacherId int
		}
		var result []Result
		database.DB.Table("groups").Select("name", "teacher_id").Where("teacher_id = ?", teacherId).Scan(&result)
		return c.JSON(result)
	} else {
		return c.JSON(fiber.Map{
			"message": "You are not a teacher",
		})
	}

}
