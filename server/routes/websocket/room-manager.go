package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// RoomManager room mamaber datatype
type RoomManager struct {
	Rooms map[string]*Room
}

// NewRoomManager room manager factory
func NewRoomManager() *RoomManager {
	rm := &RoomManager{
		Rooms: make(map[string]*Room),
	}
	return rm
}

// AddRoom adds a room
func (rm *RoomManager) AddRoom(room *Room) {
	log.Println("adding room", room.Route)
	if room == nil {
		panic("room can't be nil")
	}
	// add OnQuit callback
	room.OnQuit(rm.OnQuitRoom)
	go room.Run()
	rm.Rooms[room.Route] = room
}

// OnQuitRoom OnQuitRoom gets called when a room quits
func (rm *RoomManager) OnQuitRoom(r *Room) {
	if _, ok := rm.Rooms[r.Route]; !ok {
		return
	}
	delete(rm.Rooms, r.Route)
}

// HandleConnection handles websocket request
func (rm *RoomManager) HandleConnection(c *gin.Context, roomName string) {
	if rm == nil {
		log.Fatal("ServeWs: room manager can't be nil")
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic("Can't upgrade")
	}
	client := NewClient(conn)
	if u, exists := c.Get("user"); exists {
		client.User = u.(*datatype.User)
	}

	room, ok := rm.Rooms[roomName]
	if !ok {
		// user trying to connect to a non-existent room
		return
	}

	client.Room = room
	room.Enter(client)
	go client.write()
	go client.read()
}
