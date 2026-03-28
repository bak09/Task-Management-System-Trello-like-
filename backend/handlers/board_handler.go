package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Boards = []gin.H{}
var nextBoardID = 1

func CreateBoard(c *gin.Context) {
	var board map[string]interface{}

	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	board["id"] = nextBoardID
	nextBoardID++

	Boards = append(Boards, board)

	c.JSON(http.StatusCreated, board)
}

func GetBoards(c *gin.Context) {
	c.JSON(http.StatusOK, Boards)
}
