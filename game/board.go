package game

import (
	"context"
	"encore.dev/rlog"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Game struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

//encore:api public
func Ping(ctx context.Context) error {
	rlog.Info("ping")
	return nil
}

//encore:api public raw
func Websocket(w http.ResponseWriter, r *http.Request) {
	rlog.Info("websocket called")

	var clientID uuid.UUID
	xClientID := r.Header.Get("x-client-id")
	if xClientID != "" {
		id, err := uuid.Parse(xClientID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		clientID = id
	} else {
		clientID = uuid.New()
	}

	m.HandleRequestWithKeys(w, r, map[string]interface{}{
		"id": clientID,
	})
}
