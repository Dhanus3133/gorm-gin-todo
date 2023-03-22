package config

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Migrate() {
	DB.AutoMigrate(&Todo{})

}
