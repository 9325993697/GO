package main

import (
	"fmt"
)

func main() {
	// Create a channel of type int
	ch := make(chan int)

	// Start a goroutine to send values to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Use the for range loop to read values from the channel
	for val := range ch {
		fmt.Println("Received:", val)
	}

	fmt.Println("Channel is closed")
}
