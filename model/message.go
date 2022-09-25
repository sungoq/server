package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string
	Body      interface{}
	Timestamp uint
	Attempts  uint
}

func NewMessage(body interface{}) Message {
	return Message{
		ID:        uuid.NewString(),
		Body:      body,
		Timestamp: uint(time.Now().Unix()),
	}
}

func (m *Message) ToJSON() []byte {
	messageJson, _ := json.Marshal(m)
	return messageJson
}
