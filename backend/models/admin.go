package models

type Admin struct {
	Id        uint      `json:"id" gorm:"unique"`
	Type      string    `json:"type" gorm:"->;<-:create"`
	Name      string    `json:"name"`
	Login     string    `json:"login" gorm:"unique"`
	Password  []byte    `json:"-"`
	Timetable Timetable `gorm:"polymorphic:Owner;"`
}
