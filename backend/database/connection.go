package database

import (
	"backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn string = "host=localhost user=postgres password= dbname=gorm port=5432"
var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(&models.Admin{})
	connection.AutoMigrate(&models.Teacher{})
	connection.AutoMigrate(&models.Student{})
	connection.AutoMigrate(&models.Timetable{})
	connection.AutoMigrate(&models.Group{})
	connection.AutoMigrate(&models.Task{})
}
