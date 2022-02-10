package modules

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"log"
	"net/http"
)

var IOServer *socketio.Server

// Easier to get running with CORS. Thanks for help @Vindexus and @erkie
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func InitSocketIo() {
	IOServer := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	IOServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Join("vface")
		return nil
	})
	IOServer.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	IOServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go IOServer.Serve()
	defer IOServer.Close()
	http.Handle("/socket.io/", IOServer)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func SendPhoto(photo []byte) {

	IOServer.BroadcastToRoom("", "vface", "event:photo", photo)
}
