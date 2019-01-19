package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	sendChan chan *Message
	Room     *Room
	User     *datatype.User
}

// NewClient Client factory
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:     conn,
		sendChan: make(chan *Message, 10),
	}
}

// Send sends a message to client
func (c *Client) Send(msg *Message) {
	c.sendChan <- msg
}

// SendWithPayload sends a message to client
func (c *Client) SendWithPayload(msg string, p *MessagePayload) {
	m := NewMessage(msg)
	m.Payload = *p
	c.sendChan <- m
}

// Broadcast Broadcasts a mesage to the room of the client
func (c *Client) Broadcast(message string, payload *MessagePayload) {
	c.Room.Broadcast(message, payload)
}

// Close closes a client
func (c *Client) Close() {
	close(c.sendChan)
}

// read incomming messages on a per-connection goroutine a one reader per connection
func (c *Client) read() {
	defer func() {
		c.Room.Leave(c)
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, rawMsg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("unexpected close error: %v", err)
				return
			}
			break
		}
		rawMsg = bytes.TrimSpace(bytes.Replace(rawMsg, newline, space, -1))
		var msg Message
		json.Unmarshal(rawMsg, &msg)
		//msg.Room = c.Room.Route
		c.Room.Handle(&Action{
			Client:  c,
			Message: &msg,
		})
	}
}

// write pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	var count uint64
	for {
		select {
		case msg, ok := <-c.sendChan:

			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Print("Client.Write: ", err)
				return
			}
			w.Write(msg.Marshal())
			count = count + 1
			//log.Printf("WS: %s <-- %s %d ", c.room, msg.Message, count)
			if err := w.Close(); err != nil {
				fmt.Print("Client.Write: ", err)
				return
			}
			log.Println("-->", msg.Message)
		case <-ticker.C: // ping socket
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("SOCKET IS GONE!")
				return
			}
		}
	}
}
