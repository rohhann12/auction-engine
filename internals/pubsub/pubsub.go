package pubsub

import (
	"sync"
)

type Manager struct {
	topic map[string][]chan []byte
	mu    sync.RWMutex //only on go routine can write rest all can read simultaneosuly
}

var (
	instance *Manager
	once     sync.Once
)

// no idea of what does this do
func GetInstance() *Manager {
	once.Do(func() {
		instance = &Manager{
			topic: make(map[string][]chan []byte),
		}
	})
	return instance
}

func Subscribe(topic string, ch chan []byte) {}

func UnSubscribe(topic string, ch chan []byte) {}

func Publish(topic string, msg []byte) {}
