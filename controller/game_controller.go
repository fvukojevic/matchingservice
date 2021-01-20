package controller

import (
	"github.com/fvukojevic/matchingservice/domain"
	"github.com/fvukojevic/matchingservice/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

/**
	Returns all players from the slice
 */
func Session(c *gin.Context) {
	c.JSON(http.StatusOK, domain.GamesMap)
}


/**
	Adds a player to the room; returns room data;
 */
func Join(c *gin.Context) {
	var user domain.User

	if len(domain.UsersSlice) >= 100 {
		restErr := utils.NewBadRequestError("Lobby full, wait for someone to leave")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := utils.NewBadRequestError("could not bind given data, expecting username")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	if user.Name == "" {
		restErr := utils.NewBadRequestError("username for game not provided")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	out, _ := exec.Command("uuidgen").Output()
	user.Uuid = string(out)

	for _, player := range domain.UsersSlice {
		if player.Name == user.Name {
			restErr := utils.NewBadRequestError("username already exists in lobby")
			c.JSON(http.StatusBadRequest, restErr)
			return
		}
	}
	domain.UsersSlice = append(domain.UsersSlice, user)
	currentGame := domain.GetCurrentGame()
	currentGame.Players = append(currentGame.Players, user)
	currentGame.PlayerCount = len(currentGame.Players)
	if currentGame.PlayerCount == 4 {
		currentGame.Status = domain.StatusStarted
	}
	user.GameId = currentGame.Name

	c.JSON(http.StatusOK, currentGame)
	return
}

/**
	Removes a player from the room; returns room data;
 */
func Leave(c *gin.Context) {
	var user domain.User
	if len(domain.UsersSlice) == 0 {
		restErr := utils.NewBadRequestError("Lobby empty, nobody to leave")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := utils.NewBadRequestError("could not bind given data, expecting username and game_id")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	if user.Name == "" || user.GameId == "" {
		restErr := utils.NewBadRequestError("username or game_id not provided")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	game, err := domain.GetGameByName(user.GameId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	game.RemovePlayerFromGame(user.Name)
	domain.RemovePlayer(user.Name)
	game.PlayerCount = len(game.Players)

	c.JSON(http.StatusOK, game)
	return
}