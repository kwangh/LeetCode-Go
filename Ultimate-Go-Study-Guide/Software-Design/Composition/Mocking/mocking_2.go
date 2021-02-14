package main

import (
	"fmt"
)

type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

type mock struct{}

func (m *mock) Publish(key string, v interface{}) error {
	fmt.Println("Mock PubSub: Publish")
	return nil
}

func (m *mock) Subscribe(key string) error {
	fmt.Println("Mock PubSub: Subscribe")
	return nil
}

func main() {
	pubs := []publisher{
		New("localhost"),
		&mock{},
	}

	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
