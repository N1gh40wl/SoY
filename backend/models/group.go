package models

type Group struct {
	Id        uint   `json:"id" gorm:"unique"`
	Name      string `json:"name" gorm:"unique"`
	TeacherId uint   `json:"teacherId"`
}
