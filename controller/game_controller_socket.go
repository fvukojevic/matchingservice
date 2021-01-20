package controller

import (
	"fmt"
	"github.com/fvukojevic/matchingservice/domain"
	"github.com/fvukojevic/matchingservice/utils"
	"os/exec"
)


/**
Adds a player to the room; returns room data;
*/
func JoinGameSocket(username string) (*domain.Game, *utils.RestErr) {
	var user domain.User

	out, _ := exec.Command("uuidgen").Output()
	user.Uuid = string(out)
	user.Name = username

	for _, player := range domain.UsersSlice {
		if player.Name == user.Name {
			restErr := utils.NewBadRequestError("username already exists in lobby")
			return nil, restErr
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

	return currentGame, nil
}

/**
Removes a player from the room; returns room data;
*/
func LeaveGameSocket(username string, gameId string) (*domain.Game, *utils.RestErr) {
	var user domain.User

	user.Name = username
	user.GameId = gameId

	game, err := domain.GetGameByName(user.GameId)
	if err != nil {
		return nil, utils.NewInternalServerError(fmt.Sprintf("could not find game by id %s", gameId))
	}
	game.RemovePlayerFromGame(user.Name)
	domain.RemovePlayer(user.Name)
	game.PlayerCount = len(game.Players)

	return game, nil
}