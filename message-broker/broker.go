package broker

import (
	"log"
	"sync"
)

type memorybroker struct {
	// m is the var name
	// it is a map for key strings(topic names)
	// the values is  slice of channels, those channels hold slices of bytes
	// so each topic can have multiple channels and each channel with an array of messages

	m map[string][]chan []byte
	sync.RWMutex
}

type Broker interface {
	Close() error
	Publish(topic string, payload []byte) error
	Subscribe(topic string) (<-chan []byte, error)
	Unsubscribe(topic string, subscriber <-chan []byte) error
}

func newBroker() *memorybroker {
	return &memorybroker{
		m: make(map[string][]chan []byte),
	}
}

func (b *memorybroker) Close() error{
	b.Lock()
	b.m = make(map[string][]chan []byte)
	b.Unlock()
	return nil
}

func (b *memorybroker) Publish(topic string, payload []byte) error {
	b.RLock()
	subscribers, ok := b.m[topic]
	b.RUnlock()
	if !ok {
		log.Fatal("Topic not found")
	}

	for _, subscriber := range subscribers{
		subscriber <- payload
	}

	return nil
}

func (b *memorybroker) Subscribe(topic string) (<-chan []byte, error) {
	ch := make(chan []byte, 100)
	b.Lock()
	b.m[topic] = append(b.m[topic], ch)
	b.Unlock()
	return ch, nil
}

func (b *memorybroker) Unsubscribe(topic string, subscriber <-chan []byte) error {
	b.RLock()
	subscribers, ok := b.m[topic]
	b.RUnlock()
	if !ok {
		log.Fatal("could not find the topic")
	}
	var subs []chan []byte
	for _, s := range subscribers {
		if s == subscriber {
			continue
		}
		subs = append(subs, s)
	}

	b.Lock()
	b.m[topic] = subs
	b.Unlock()
	
	return nil
}