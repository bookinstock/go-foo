package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	subscribers []chan string
	mu          sync.Mutex
}

func (ps *PubSub) Subscribe() <-chan string {
	ch := make(chan string, 10) // Increased buffer size to prevent blocking
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.subscribers = append(ps.subscribers, ch)
	return ch
}

func (ps *PubSub) Publish(msg string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for _, ch := range ps.subscribers {
		select {
		case ch <- msg: // Try to send without blocking
		default: // Skip if channel is full
		}
	}
	fmt.Println("Published message:", msg)
}

func RunPubSub() {
	ps := &PubSub{}

	// Subscriber 1
	go func() {
		sub := ps.Subscribe()
		for msg := range sub {
			fmt.Println("Subscriber 1:", msg)
		}
	}()

	// Subscriber 2
	go func() {
		sub := ps.Subscribe()
		for msg := range sub {
			fmt.Println("Subscriber 2:", msg)
		}
	}()

	// Give subscribers time to start
	time.Sleep(100 * time.Millisecond)

	// Publish messages
	ps.Publish("Hello")
	ps.Publish("World")

	time.Sleep(1 * time.Second)
}
