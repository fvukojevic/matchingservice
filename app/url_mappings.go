package app

import (
	"fmt"
	"github.com/fvukojevic/matchingservice/controller"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func mapUrls() {
	router.GET("/session", controller.Session)
	router.POST("/join", controller.Join)
	router.POST("/leave", controller.Leave)
}

func mapSocketUrls() {
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "join", func(s socketio.Conn, username string) {
		if username == "" {
			s.Emit("reply", "Username missing")
		}
		game, err := controller.JoinGameSocket(username)
		if err != nil {
			s.Emit("reply", err)
		}
		s.Emit("reply", game)
	})

	server.OnEvent("/", "leave", func(s socketio.Conn, data map[string]string) {
		if _, ok := data["username"]; !ok {
			s.Emit("reply", "username not provided")
		}
		if _, ok := data["game_id"]; !ok {
			s.Emit("reply", "game_id not provided")
		}
		game, err := controller.LeaveGameSocket(data["username"], data["game_id"])
		if err != nil {
			s.Emit("reply", err)
		}

		s.Emit("reply", game)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}