package main

import "fmt"

func main() {
    // create a buffered channel with capacity 3
    ch := make(chan int, 3)

    // add some values to the channel
    ch <- 1
    ch <- 2
    ch <- 3

    // find the capacity of the channel
    fmt.Printf("Channel capacity: %d\n", cap(ch))

    // find the length of the channel
    fmt.Printf("Channel length: %d\n", len(ch))

    // read values from the channel
    fmt.Printf("Values from the channel: ")
    for i := 0; i < cap(ch); i++ {
        fmt.Printf("%d ", <-ch)
    }

    // find the modified length of the channel
    fmt.Printf("\nChannel length after reading values: %d\n", len(ch))
}
