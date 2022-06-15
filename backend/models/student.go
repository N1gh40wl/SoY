package models

type Student struct {
	Id         uint      `json:"id" gorm:"unique"`
	Type       string    `json:"type" gorm:"->;<-:create"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Pantonymic string    `json:"pantonymic"`
	Login      string    `json:"login" gorm:"unique"`
	Password   []byte    `json:"-"`
	Timetable  Timetable `gorm:"polymorphic:Owner;"`
	GroupId    uint      `json:"group_id"`
}
