package domain

import (
	"fmt"
	"github.com/fvukojevic/matchingservice/utils"
)

var (
	GamesMap = map[int]*Game{}
)

const (
	MaxPlayers = 4
	StatusWaiting = "Waiting"
	StatusStarted = "Started"
)

type Game struct {
	Name        string           `json:"name"`
	Status      string           `json:"status"`
	Players     []User           `json:"players"`
	PlayerCount int              `json:"player_count"`
}

/**
  Returns the current game that we are trying to start
*/
func GetCurrentGame() *Game {
	for id, game := range GamesMap {
		if GamesMap[id].PlayerCount < MaxPlayers {
			return game
		}
	}

	nextId := len(GamesMap) + 1
	game := initGame(nextId)
	GamesMap[nextId] = game

	return game
}

/**
This function returns a game from map, by the game's name
*/
func GetGameByName(name string) (*Game, *utils.RestErr) {
	for id, game := range GamesMap {
		if GamesMap[id].Name == name {
			return game, nil
		}
	}

	return nil, utils.NewInternalServerError("Game not found")
}

func (game *Game) RemovePlayerFromGame(username string) {
	players := game.Players

	for i := range players {
		if players[i].Name == username {
			players = append(players[:i], players[i+1:]...)
			break
		}
	}
	game.Status = StatusWaiting
	game.Players = players
	return
}

/**
  Creates a new game instance and prepares sockets for listening on events from the client
*/
func initGame(nextId int) *Game {
	concatenated := fmt.Sprint("Game ", nextId)
	game := Game{
		Name:        concatenated,
		Players:     []User{},
		Status:      StatusWaiting,
		PlayerCount: 0,
	}
	return &game
}
