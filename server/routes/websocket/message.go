package websocket

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Action Action data type
type Action struct {
	Client  *Client
	Message *Message
}

// MessagePayload Message Payload datatype
type MessagePayload map[string]interface{}

// Message message datatype
type Message struct {
	//Room    string         `json:"-"`
	Message string         `json:"m"`
	Payload MessagePayload `json:"p"`
}

// NewMessage message factory
func NewMessage(msg string) *Message {
	return &Message{
		Message: msg,
		Payload: map[string]interface{}{},
	}
}

// Set sets a value in the  message's payload
func (m *Message) Set(key string, val interface{}) {
	m.Payload[key] = val
}

// Marshal marshals the message
func (m *Message) Marshal() []byte {
	b, _ := json.Marshal(m)
	return b
}

// Unmarshal unmarshals the message
func (m *Message) Unmarshal(data []byte) {
	json.Unmarshal(data, m)
}

// GetInt64 try to get an int from payload
func (m *Message) GetInt64(key string) (int64, bool) {
	if raw, exists := m.Payload[key]; exists {
		switch value := raw.(type) {
		case int32:
			return int64(value), true
		case int64:
			return value, true
		case float64:
			return int64(value), true
		}
	}
	return 0, false
}

// GetFloat64 try to get a float64 from payload
func (m *Message) GetFloat64(key string) (float64, bool) {
	if raw, exists := m.Payload[key]; exists {
		switch value := raw.(type) {
		case float64:
			return value, true
		}
	}
	return 0, false
}

// GetString try to get a string from payload
func (m *Message) GetString(key string) (string, bool) {
	if raw, exists := m.Payload[key]; exists {
		switch value := raw.(type) {
		case string:
			return value, true
		case int:
			return strconv.Itoa(value), true
		case float64:
			return fmt.Sprintf("%f", value), true
		case float32:
			return fmt.Sprintf("%f", value), true
		}
	}
	return "", false
}

// GetBool try to get a bool from payload
func (m *Message) GetBool(key string) (bool, bool) {
	if raw, exists := m.Payload[key]; exists {
		switch value := raw.(type) {
		case bool:
			return value, true
		}
	}
	return false, false
}
