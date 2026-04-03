package handlers

import (
	"net/http"

	"taskmanager/config"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
)

func CreateBoard(c *gin.Context) {
	var board models.Board

	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&board)
	c.JSON(http.StatusCreated, board)
}

func GetBoards(c *gin.Context) {
	var boards []models.Board
	config.DB.Find(&boards)
	c.JSON(http.StatusOK, boards)
}