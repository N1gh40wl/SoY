package models

type Task struct {
	Id          uint   `json:"id" gorm:"unique"`
	Condition   string `json:"condition"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TaskLink    string `json:"taskLink"`
	AnswerLink  string `json:"answer_link"`
	Mark        string `json:"mark"`
	Subject     string `json:"subject"`
	TeacherId   uint   `json:"teacherId"`
	StudentId   uint   `json:"studentId"`
}
