package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GameBoard struct {
	GameID int64    `json:"game_id"`
	Board  []string `json:"board"`
}

var (
	currentGames       = []GameBoard{}
	gameID       int64 = 0
	emptyBoard         = []string{"---", "---", "---"}
)

func getGameBoardByID(c *gin.Context) {
	id := c.Param("id")

	for _, g := range currentGames {
		if string(g.GameID) == id {
			c.IndentedJSON(http.StatusOK, g)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Game not found with given ID"})
}

func createGame(c *gin.Context) {
	var newGame GameBoard
	gameID += 1
	newGame.GameID = gameID
	newGame.Board = emptyBoard
	currentGames = append(currentGames, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

func main() {
	router := gin.Default()
	router.POST("/game/new", createGame)
	router.GET("/game/:id", getGameBoardByID)
	router.Run("localhost:8080")
}
