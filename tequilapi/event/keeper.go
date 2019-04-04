package event

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mysteriumnetwork/node/core/connection"
)

type EventKeeper struct {
	clients map[*websocket.Conn]struct{}
	lock    sync.Mutex
}

func NewEventKeeper() *EventKeeper {
	return &EventKeeper{
		clients: make(map[*websocket.Conn]struct{}),
	}
}

func (ek *EventKeeper) AddClient(c *websocket.Conn) {
	ek.lock.Lock()
	defer ek.lock.Unlock()
	ek.clients[c] = struct{}{}
}

func (ek *EventKeeper) ConsumeConnectionStateEvent(state connection.Status) {
	ek.lock.Lock()
	defer ek.lock.Unlock()
	for k := range ek.clients {
		k.SetWriteDeadline(time.Now().Add(time.Millisecond * 100))
		err := k.WriteJSON(state)
		if err != nil {
			// TODO: Logging
			ek.cleanupClient(k)
			continue
		}
	}
}

func (ek *EventKeeper) cleanupClient(c *websocket.Conn) {
	defer c.Close()
	delete(ek.clients, c)
}

type Client struct {
	Connection *websocket.Conn
}
