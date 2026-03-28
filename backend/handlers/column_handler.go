package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Columns = []gin.H{}
var nextColumnID = 1

func CreateColumn(c *gin.Context) {
	var column map[string]interface{}

	if err := c.ShouldBindJSON(&column); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	column["id"] = nextColumnID
	nextColumnID++

	Columns = append(Columns, column)

	c.JSON(http.StatusCreated, column)
}

func GetColumns(c *gin.Context) {
	c.JSON(http.StatusOK, Columns)
}