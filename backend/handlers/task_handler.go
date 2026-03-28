package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Tasks = []gin.H{}
var nextTaskID = 1

func CreateTask(c *gin.Context) {
	var task map[string]interface{}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	task["id"] = nextTaskID
	nextTaskID++

	Tasks = append(Tasks, task)

	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, Tasks)
}

func GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, task := range Tasks {
		if int(task["id"].(int)) == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, task := range Tasks {
		if int(task["id"].(int)) == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}