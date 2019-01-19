package websocket

import "log"

// RoomHandler room handler datatype
type RoomHandler func(*Action)

// RoomCientLifecycleCallback RoomCientLifecycleCallback datatype
type RoomCientLifecycleCallback func(c *Client)

// Room room datatype
type Room struct {
	Route                   string
	clients                 map[*Client]bool
	enterChan               chan *Client
	leaveChan               chan *Client
	recvChan                chan *Action
	broadcastChan           chan *Message
	quitchan                chan bool
	handlers                map[string]RoomHandler
	internalCallbackOnEnter RoomCientLifecycleCallback
	internalCallbackOnLeave RoomCientLifecycleCallback
	onQuitCallback          func(r *Room)
	IsRunning               bool
}

// NewRoom rom factory
func NewRoom(route string) *Room {
	return &Room{
		Route:         route,
		clients:       make(map[*Client]bool),
		handlers:      make(map[string]RoomHandler),
		enterChan:     make(chan *Client),
		leaveChan:     make(chan *Client),
		recvChan:      make(chan *Action),
		broadcastChan: make(chan *Message),
		quitchan:      make(chan bool),
		IsRunning:     false,
	}
}

// Enter enter room
func (r *Room) Enter(c *Client) {
	log.Println("Room.Enter", r.Route)
	r.enterChan <- c
}

// Leave leaves room
func (r *Room) Leave(c *Client) {
	log.Println("Room.Leae", r.Route)
	r.leaveChan <- c
}

func (r *Room) handleEnter(c *Client) {
	r.clients[c] = true
	if r.internalCallbackOnEnter != nil {
		r.internalCallbackOnEnter(c)
	}
}

func (r *Room) handleLeave(c *Client) {
	if _, ok := r.clients[c]; !ok {
		return
	}
	c.Close()
	delete(r.clients, c)
	if r.internalCallbackOnLeave != nil {
		r.internalCallbackOnLeave(c)
	}
}

// SetOnEnter sets room onEnter callback
func (r *Room) SetOnEnter(f RoomCientLifecycleCallback) {
	r.internalCallbackOnEnter = f
}

// SetOnLeave sets room OnLeave callback
func (r *Room) SetOnLeave(f RoomCientLifecycleCallback) {
	r.internalCallbackOnLeave = f
}

// AddHandler adds a handler for the given msg
func (r *Room) AddHandler(message string, f RoomHandler) {
	r.handlers[message] = f
}

// Handle handles an incomming action
func (r *Room) Handle(a *Action) {
	log.Println("---------> WS MSG ", a.Message.Message, a.Message.Payload)
	r.recvChan <- a
}

// Broadcast broadcasts the given message to all room clients
func (r *Room) Broadcast(message string, p *MessagePayload) {
	go func() {
		if p == nil {
			p = &MessagePayload{}
		}
		msg := NewMessage(message)
		msg.Payload = *p
		r.broadcastChan <- msg
	}()
}

// BroadcastMsg broadcasts the given message to all room clients
func (r *Room) BroadcastMsg(msg *Message) {
	go func() {
		r.broadcastChan <- msg
	}()
}

// ClientsCount returns the number of clients connected
func (r *Room) ClientsCount() int {
	return len(r.clients)
}

// Quit quits the room
func (r *Room) Quit() {
	close(r.quitchan)
	if r.onQuitCallback != nil {
		r.onQuitCallback(r)
	}
}

// OnQuit sets a callback to be called when the room quits
func (r *Room) OnQuit(f func(r *Room)) {
	r.onQuitCallback = f
}

// Run runs room main loop
func (r *Room) Run() {
	log.Println("Running WebSockRoom!!: " + r.Route)
	r.IsRunning = true
	for {
		select {
		case client := <-r.enterChan:
			r.handleEnter(client)
			break
		case client := <-r.leaveChan:
			r.handleLeave(client)
			break
		case action := <-r.recvChan:
			if handler, ok := r.handlers[action.Message.Message]; ok {
				handler(action)
			}
			break
		case msg := <-r.broadcastChan:
			for client := range r.clients {
				client.Send(msg)
			}
			break
		case <-r.quitchan:
			r.IsRunning = false
			return
		}
	}
}
