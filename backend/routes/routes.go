package routes

import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)
	app.Delete("api/delete", controllers.DeleteAccount)
	app.Get("api/timetable", controllers.TimetableShow)
	app.Post("api/createTask", controllers.CreateTask)
	app.Post("api/answerTask", controllers.AnswerTask)
	app.Post("api/downloadFile", controllers.DonwloadFile)
	app.Get("api/showFileTask/*", controllers.ShowFileTask)
	app.Post("api/showOneTask", controllers.ShowOneTask)
	app.Post("api/rateTask", controllers.RateTask)
	app.Get("api/showTeachersClasses", controllers.ShowTeachersClasses)
	app.Post("api/showTaskByClass", controllers.ShowTaskByClass)
	app.Post("api/showTaskBySubject", controllers.ShowTaskBySubject)
	app.Delete("api/deleteTask", controllers.DeleteTask)
	app.Get("api/group", controllers.TeacherGroupShow)

}
