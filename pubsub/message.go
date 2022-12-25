package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	chanel    Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("Message %s", m.chanel)
}

func (m *Message) Chanel() Topic {
	return m.chanel
}

func (m *Message) SetChanel(chanel Topic) {
	m.chanel = chanel

}

func (m *Message) Data() interface{} {
	return m.data
}