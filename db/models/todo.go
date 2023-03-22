package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}

type UpdateTodoInput struct {
	Title     string
	Completed bool
}
