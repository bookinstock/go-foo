package concurrency

import (
	"fmt"
)

func priorityFanIn(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		defer close(out)
		for ch1 != nil || ch2 != nil {
			// First try to read from ch1
			if ch1 != nil {
				select {
				case msg, ok := <-ch1:
					if !ok {
						ch1 = nil
						continue
					}
					out <- msg
					continue
				default:
				}
			}

			// Only try ch2 if nothing available on ch1
			if ch2 != nil {
				select {
				case msg, ok := <-ch2:
					if !ok {
						ch2 = nil
						continue
					}
					out <- msg
				default:
				}
			}
		}
	}()
	return out
}

func RunPriorityFanIn() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 15; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	// time.Sleep(1 * time.Second)

	for msg := range priorityFanIn(ch1, ch2) {
		fmt.Println("Received:", msg)
	}
}
