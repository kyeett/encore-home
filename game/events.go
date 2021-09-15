package game

import "github.com/google/uuid"

type EventType string

const (
	EventTypeGameCreated     EventType = "game_created"
	EventTypeClientConnected EventType = "client_connected"
)

type EventCreated struct {
	Type EventType `json:"type"`
	Game Game      `json:"game"`
}

type EventClientConnected struct {
	Type     EventType `json:"type"`
	ClientID uuid.UUID `json:"client_id"`
}
