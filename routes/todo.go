package routes

import (
	"github.com/Dhanus3133/gorm-gin-todo/handlers"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.RouterGroup) {
	todo := router.Group("/todos")
	{
		todo.POST("", handlers.CreateTodo)
		todo.GET("", handlers.GetTodos)
		todo.GET("/:id", handlers.GetTodoByID)
		todo.PUT("/:id", handlers.UpdateTodo)
		todo.DELETE("/:id", handlers.DeleteTodo)
	}
}
