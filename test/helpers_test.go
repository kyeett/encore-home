package test

import (
	game "encore.app/game"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	testWebsocketHost = "ws://localhost:4060/game.Websocket"
	testURL           = "http://localhost:4060/"
)

func mustCreateGame(t *testing.T, c *websocket.Conn) *game.EventCreated {
	c.WriteJSON(struct{}{})
	var evt game.EventCreated
	err := c.ReadJSON(&evt)
	require.NoError(t, err)
	require.Equal(t, game.EventTypeGameCreated, evt.Type)
	return &evt
}

func mustConnectWithID(t *testing.T, clientID uuid.UUID) *websocket.Conn {
	d := &websocket.Dialer{
		HandshakeTimeout: 1 * time.Second,
	}
	c, resp, err := d.Dial(testWebsocketHost, map[string][]string{
		"x-client-id": {clientID.String()},
	})
	defer resp.Body.Close()
	require.NoError(t, err)

	// Verify client ID
	var connEvent game.EventClientConnected
	mustReceiveJSON(t, c, &connEvent)
	require.Equal(t, game.EventTypeClientConnected, connEvent.Type)
	require.Equal(t, clientID, connEvent.ClientID)

	return c
}

func mustConnect(t *testing.T) *websocket.Conn {
	d := &websocket.Dialer{
		HandshakeTimeout: 1 * time.Second,
	}
	c, resp, err := d.Dial(testWebsocketHost, nil)
	defer resp.Body.Close()
	require.NoError(t, err)

	var connEvent game.EventClientConnected
	mustReceiveJSON(t, c, &connEvent)
	require.Equal(t, game.EventTypeClientConnected, connEvent.Type)

	return c
}

func mustReceiveJSON(t *testing.T, c *websocket.Conn, v interface{}) {
	require.NoError(t, c.ReadJSON(v))
}
