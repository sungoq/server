package model

import "encoding/json"

type Message struct {
	ID      string
	Topic   string
	Payload string
}

func (m *Message) ToJSON() string {
	messageJson, _ := json.Marshal(m)
	return string(messageJson)
}
