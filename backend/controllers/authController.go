package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	if AdminSearch(data["login"]) != 0 || TeacherSearch(data["login"]) != 0 || StudentSearch(data["login"]) != 0 {
		return c.JSON(fiber.Map{
			"message": "login already used",
		})
	}

	if data["type"] == "admin" {
		CreateAdmin(data, c)
	} else if data["type"] == "teacher" {
		CreateTeacher(data, c)
	} else if data["type"] == "student" {
		CreateStudent(data, c)
	}
	return err
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	if AdminSearch(data["login"]) != 0 {
		var admin models.Admin
		database.DB.Where("login = ?", data["login"]).First(&admin)

		if err := bcrypt.CompareHashAndPassword(admin.Password, []byte(data["password"])); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "incorrect password",
			})
		}

		t, _ := jwt.ParseTime(time.Now().Add(time.Hour * 2))
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    string(admin.Login),
			ExpiresAt: t,
		})

		token, err := claims.SignedString([]byte(SecretKey))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "could not login",
			})
		}

		cockie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 2),
			HTTPOnly: true,
		}
		c.Cookie(&cockie)
		return c.JSON(fiber.Map{
			"message": "SuccessA",
		})
	} else if TeacherSearch(data["login"]) != 0 {
		var teacher models.Teacher
		database.DB.Where("login = ?", data["login"]).First(&teacher)

		if err := bcrypt.CompareHashAndPassword(teacher.Password, []byte(data["password"])); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "incorrect password",
			})
		}

		t, _ := jwt.ParseTime(time.Now().Add(time.Hour * 2))
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    string(teacher.Login),
			ExpiresAt: t,
		})

		token, err := claims.SignedString([]byte(SecretKey))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "could not login",
			})
		}

		cockie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 2),
			HTTPOnly: true,
		}
		c.Cookie(&cockie)
		return c.JSON(fiber.Map{
			"message": "SuccessT",
		})
	} else if StudentSearch(data["login"]) != 0 {
		var student models.Student
		database.DB.Where("login = ?", data["login"]).First(&student)

		if err := bcrypt.CompareHashAndPassword(student.Password, []byte(data["password"])); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "incorrect password",
			})
		}

		t, _ := jwt.ParseTime(time.Now().Add(time.Hour * 2))
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    string(student.Login),
			ExpiresAt: t,
		})

		token, err := claims.SignedString([]byte(SecretKey))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "could not login",
			})
		}

		cockie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 2),
			HTTPOnly: true,
		}
		c.Cookie(&cockie)
		return c.JSON(fiber.Map{
			"message": "SuccessS",
		})
	} else {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

}

func User(c *fiber.Ctx) error {
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
		var admin models.Admin
		database.DB.Where("login = ?", claims.Issuer).First(&admin)
		return c.JSON(admin)
	} else if TeacherSearch(claims.Issuer) != 0 {
		var teacher models.Teacher
		database.DB.Where("login= ?", claims.Issuer).First(&teacher)
		return c.JSON(teacher)
	} else if StudentSearch(claims.Issuer) != 0 {
		var student models.Student
		database.DB.Where("login = ?", claims.Issuer).First(&student)
		return c.JSON(student)
	}
	return nil
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func DeleteAccount(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {

		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unautheficated",
		})
	}
	var admin models.Admin
	claims := token.Claims.(*jwt.StandardClaims)

	database.DB.Where("id = ?", claims.Issuer).First(&admin)
	database.DB.Delete(models.Admin{}, admin.Id)

	return c.JSON(fiber.Map{
		"message": "succes",
	})
}

func Test(c *fiber.Ctx) error {

	return c.SendString("Hello, World 👋!")
}
