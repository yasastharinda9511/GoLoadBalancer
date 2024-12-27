package message

import (
	"github.com/google/uuid"
)

// MessageBase defines common attributes and methods for different message types.
type Message struct {
	uid string
}

// NewMessageBase initializes the base part of a message.
func NewMessage() *Message {
	return &Message{
		uid: uuid.New().String(),
	}
}

// GetQueryParams returns the query parameters of the message.
func (base *Message) GetQueryParams() string {
	return base.uid
}

func (base *Message) GetUID() string {
	return base.uid
}
