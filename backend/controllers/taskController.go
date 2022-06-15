package controllers

import (
	"backend/database"
	"backend/models"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func DonwloadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("document")
	if err != nil {
		return err
	}

	data := new(models.Parse)
	err = c.BodyParser(data)
	if err != nil {
		return err
	}
	// Check for errors:
	log.Println(data.Test)
	// 👷 Save file inside uploads folder under current working directory:

	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", data.Test))
	if err != nil {
		log.Println(err)

		return c.JSON(fiber.Map{
			"message": "Fail to store file",
		})
	} else {

		return c.JSON(fiber.Map{
			"message": "Success",
		})
	}

}

func CreateTask(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

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
		var Id []string
		database.DB.Table("students").Select("id").Where("group_id = ?", data["group_id"]).Scan(&Id)
		for i := 0; i < len(Id); i++ {
			sId, err := strconv.ParseUint(Id[i], 10, 32)
			if err != nil {
				return c.JSON(fiber.Map{
					"message": "Error",
				})
			}
			task := models.Task{
				Condition:   "given",
				Name:        data["name"],
				Description: data["description"],
				TeacherId:   TeacherSearch(claims.Issuer),
				Subject:     data["subject"],
				TaskLink:    "",
				StudentId:   uint(sId),
			}
			database.DB.Create(&task)
			database.DB.Model(&task).Select("task_link").Updates(models.Task{TaskLink: "./uploads/" + strconv.FormatUint(uint64(TeacherSearch(claims.Issuer)), 10) + "_" + data["name"] + "_teacher.pdf"})
			log.Println(task.Id)

		}

	} else {
		return c.JSON(fiber.Map{
			"message": "You are not a teacher",
		})
	}

	return c.JSON(fiber.Map{

		"message": "Success",
	})
}

func AnswerTask(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

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
	if StudentSearch(claims.Issuer) != 0 {

		var task models.Task
		database.DB.Where("id = ?", data["id"]).First(&task)
		if StudentSearch(claims.Issuer) == task.StudentId {
			database.DB.Model(&task).Select("answer_link", "condition").Updates(models.Task{AnswerLink: "./uploads/" + strconv.FormatUint(uint64(task.StudentId), 10) + "_" + strconv.FormatUint(uint64(task.Id), 10) + "_student.pdf", Condition: "answered"})

		} else {
			return c.JSON(fiber.Map{
				"message": "This is not your assignment.",
			})

		}
		return c.JSON(fiber.Map{
			"message": "success",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "You are not a student",
		})
	}

}

func RateTask(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
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

		var task models.Task
		database.DB.Where("id = ?", data["id"]).First(&task)
		if TeacherSearch(claims.Issuer) == task.TeacherId {
			database.DB.Model(&task).Select("mark", "condition").Updates(models.Task{Mark: data["mark"], Condition: "marked"})

		} else {
			return c.JSON(fiber.Map{
				"message": "This is not your assignment.",
			})

		}
		return c.JSON(fiber.Map{
			"message": "success",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "You are not a teacher",
		})
	}
}

func ShowOneTask(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	var task models.Task
	database.DB.Where("id = ?", data["id"]).First(&task)
	err = c.JSON(task)
	if err != nil {
		return err
	}
	return nil
}

func ShowFileTask(c *fiber.Ctx) error {
	path := c.Path()
	path = path[18:]
	log.Println(path)

	return c.Download(path)
}

func ShowTeachersClasses(c *fiber.Ctx) error {
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
		var groups []models.Group
		database.DB.Table("groups").Where("teacher_id = ?", TeacherSearch(claims.Issuer)).Scan(&groups)
		return c.JSON(groups)
	} else {
		return c.JSON(fiber.Map{
			"message": "You are not a teacher",
		})
	}
}

func ShowTaskByClass(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
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
		var tasks []models.Task
		database.DB.Table("tasks").Where("teacher_id = ?", TeacherSearch(claims.Issuer)).Scan(&tasks)
		var finTasks []models.Task
		GId, err := strconv.ParseUint(data["group_id"], 10, 32)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Error",
			})
		}
		for i := 0; i < len(tasks); i++ {
			if StudentClass(tasks[i].StudentId) == uint(GId) && tasks[i].Subject == data["subject"] {
				finTasks = append(finTasks, tasks[i])
			}

		}
		return c.JSON(finTasks)
	} else {
		return c.JSON(fiber.Map{
			"message": "You are not a teacher",
		})
	}
}

func ShowTaskBySubject(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
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

	if StudentSearch(claims.Issuer) != 0 {
		var tasks []models.Task
		database.DB.Table("tasks").Where("student_id = ?", StudentSearch(claims.Issuer)).Scan(&tasks)
		var finTasks []models.Task

		for i := 0; i < len(tasks); i++ {
			if tasks[i].Subject == data["subject"] && tasks[i].Condition == data["condition"] {
				finTasks = append(finTasks, tasks[i])
			}

		}
		return c.JSON(finTasks)
	} else if TeacherSearch(claims.Issuer) != 0 {
		var tasks []models.Task
		database.DB.Table("tasks").Where("teacher_id = ?", TeacherSearch(claims.Issuer)).Scan(&tasks)
		var finTasks []models.Task
		GId, err := strconv.ParseUint(data["group_id"], 10, 32)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Errorl",
			})
		}
		for i := 0; i < len(tasks); i++ {
			if tasks[i].Subject == data["subject"] && tasks[i].Condition == data["condition"] && StudentClass(tasks[i].StudentId) == uint(GId) {
				finTasks = append(finTasks, tasks[i])
			}

		}
		return c.JSON(finTasks)
	} else {
		return c.JSON(fiber.Map{
			"message": "You are not in system",
		})
	}
}

/*
	GId, err := strconv.ParseUint(data["group_id"], 10, 32)
			if err != nil {
				return c.JSON(fiber.Map{
					"message": "Errorl",
				})
			}
*/

func DeleteTask(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
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
		var task models.Task

		TId, err := strconv.ParseUint(data["group_id"], 10, 32)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Error",
			})
		}
		database.DB.Where("id = ?", uint(TId)).First(&task)
		database.DB.Delete(models.Task{}, task.Id)

		return c.JSON(fiber.Map{
			"message": "succes",
		})

	} else {
		return c.JSON(fiber.Map{
			"message": "You are not a teacher",
		})
	}
}
