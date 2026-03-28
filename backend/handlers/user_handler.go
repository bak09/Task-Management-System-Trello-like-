package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Users = []gin.H{}
var nextUserID = 1

func CreateUser(c *gin.Context) {
	var user map[string]interface{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	user["id"] = nextUserID
	nextUserID++

	Users = append(Users, user)

	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, Users)
}
