package main

import (
	"fmt"
)

type PubSub struct {
	host string
}

func New(host string) *PubSub {
	ps := PubSub{
		host: host,
	}

	return &ps
}

func (ps *PubSub) Publish(key string, v interface{}) error {
	fmt.Println("Actual PubSub: Publish")
	return nil
}

func (ps *PubSub) Subscribe(key string) error {
	fmt.Println("Actual PubSub: Subscribe")
	return nil
}
