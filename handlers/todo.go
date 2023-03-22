package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dhanus3133/gorm-gin-todo/db"
	"github.com/Dhanus3133/gorm-gin-todo/db/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	db.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)

	db.DB.Create(&todo)

	c.JSON(http.StatusCreated, gin.H{"todo": todo})
}

func GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}
	fmt.Println(id)

	var todo models.Todo
	result := db.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	var todo models.Todo
	result := db.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input models.UpdateTodoInput
	err = c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo.Title = input.Title
	todo.Completed = input.Completed

	db.DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := db.DB.Where("id = ?", id).Delete(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
