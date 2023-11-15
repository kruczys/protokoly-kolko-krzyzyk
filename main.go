package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GameBoard struct {
	GameID string   `json:"game_id"`
	Board  []string `json:"board"`
}

var (
	currentGames       = []GameBoard{}
	gameID       int64 = 0
	emptyBoard         = []string{"---", "---", "---"}
)

func createGame(c *gin.Context) {
	var newGame GameBoard
	gameID += 1
	newGame.GameID = strconv.FormatInt(gameID, 10)
	newGame.Board = emptyBoard
	currentGames = append(currentGames, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

func getGameBoardByID(c *gin.Context) {
	id := c.Param("id")

	for _, g := range currentGames {
		if g.GameID == id {
			c.IndentedJSON(http.StatusOK, g)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Game not found with ID of %s", id)})
}

func deleteGameByID(c *gin.Context) {
	id := c.Param("id")
	return
}

func main() {
	router := gin.Default()
	router.POST("/game/new", createGame)
	router.GET("/game/:id", getGameBoardByID)
	router.Run("localhost:8080")
}
