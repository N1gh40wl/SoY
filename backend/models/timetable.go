package models

type Timetable struct {
	Id        uint   `json:"id" gorm:"unique"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	OwnerID   uint
	OwnerType string
}
