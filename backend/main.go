package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TaskFlow Backend is running 🚀",
		})
	})

	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"tasks": []string{"Task 1", "Task 2"},
		})
	})

	r.Run(":8080")
}