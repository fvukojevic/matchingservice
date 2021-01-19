package domain

import (
	socketio "github.com/googollee/go-socket.io"
)

var (
	GamesMap = map[int]Game{}
)

type Game struct {
	Name        string          `json:"name"`
	Players     []User          `json:"players"`
	PlayerCount int             `json:"player_count"`
	Socket      socketio.Server `json:"password"`
}
