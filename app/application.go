package app

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

var (
	router = gin.Default()
	server, err = socketio.NewServer(nil)
)

func StartApplication() {
	mapUrls()
	//mapSocketUrls()
	router.Run(":8080")
}
