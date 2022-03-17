package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

type Input struct {
	Provision string `json:"provisioner" xml:"provisioner"`
}

type Output struct {
	Deneme string `json:"provisioner" xml:"provisioner"`
}

func Enable(uri, username, password string) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: uri, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	messages := make(chan Input)

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				return
			}
			log.Printf("recv: %s", message)
			var input Input
			if err = json.Unmarshal(message, &input); err != nil {
				continue
			}
			messages <- input
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		case message := <-messages:
			//ENGINE.YAML I OVERRIDE ET GELEN PARAMETRELERE GORE
			//flow.OpenSource(message.Provision)
			fmt.Printf("INPUT FROM SERVER: %v", message)
			//RESPONSE
			output := Output{Deneme: "123"}
			outputByte, _ := json.Marshal(output)
			err := c.WriteMessage(websocket.TextMessage, outputByte)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
