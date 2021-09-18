package game

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/olahol/melody.v1"
)

var m *melody.Melody

func init() {
	m = melody.New()
	m.HandleMessage(handleMessageReceived)
	m.HandlePong(handlePong)
	m.HandleConnect(handleConnect)
}

func broadcastJSON(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return m.Broadcast(b)
}

func writeJSON(s *melody.Session, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return s.Write(b)
}

func handleMessageReceived(s *melody.Session, msg []byte) {
	//g := Game{ID: uuid.NewString(), CreatedAt: time.Now()}
	//
	//evt := EventCreated{EventTypeGameCreated, g}
	//
	//// Best effort update clients
	//_ = broadcastJSON(evt)
	m.Broadcast(msg)
}

func handlePong(s *melody.Session) {
	fmt.Println("pong received!")
}

func handleConnect(s *melody.Session) {
	fmt.Println("CONNECT :D", s.MustGet("id"))
	clientID := s.MustGet("id").(uuid.UUID)
	writeJSON(s, EventClientConnected{Type: EventTypeClientConnected, ClientID: clientID})
}
