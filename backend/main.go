package main

import (
	"net/http"

	"taskmanager/config"
	"taskmanager/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TaskFlow Backend is running 🚀",
		})
	})

	// users
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)

	// boards
	r.POST("/boards", handlers.CreateBoard)
	r.GET("/boards", handlers.GetBoards)

	// columns
	r.POST("/columns", handlers.CreateColumn)
	r.GET("/columns", handlers.GetColumns)

	// tasks
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTaskByID)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.Run(":8080")
}
