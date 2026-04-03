package handlers

import (
	"net/http"

	"taskmanager/config"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
)

func CreateColumn(c *gin.Context) {
	var column models.Column

	if err := c.ShouldBindJSON(&column); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&column)
	c.JSON(http.StatusCreated, column)
}

func GetColumns(c *gin.Context) {
	var columns []models.Column
	config.DB.Find(&columns)
	c.JSON(http.StatusOK, columns)
}