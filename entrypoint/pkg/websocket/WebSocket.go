package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Input struct {
	Deneme string `json:"deneme" xml:"deneme"`
}

func Enable(uri string, username string, password string) {
	/*setupRoutes()
	tlsEnabled, _ := strconv.ParseBool()
	if tlsEnabled {
		fmt.Println("Websocket Routes Enabled (:443/ws)")
		log.Fatal(http.ListenAndServeTLS(":443",
			*params["certPath"],
			*params["keyPath"],
			nil))
	} else {
		fmt.Println("Websocket Routes Enabled (:80/ws)")
		log.Fatal(http.ListenAndServe(":80", nil))
	}*/

}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func wsEndpoint(writer http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client succesfully Connected...")
	defer func() { conn.Close() }()
	reader(conn)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		input := Input{}
		if err := json.Unmarshal(p, &input); err != nil {
			log.Printf("error decoding request: %v", err)
			return
		}
		fmt.Println(input)
		if err := conn.WriteMessage(messageType, []byte("Data alındı okay")); err != nil {
			log.Println(err)
			return
		}
	}
}
