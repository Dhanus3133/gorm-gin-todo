package db

import (
	"github.com/Dhanus3133/gorm-gin-todo/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Todo{})
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		panic(err)
	}
	db.Close()
}
