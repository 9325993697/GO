package main

import "fmt"

func fibonacci(n int, c chan int) {
    a, b := 0, 1
    for i := 0; i < n; i++ {
        c <- a
        a, b = b, a+b
    }
    close(c)
}

func main() {
    // create a channel to hold the Fibonacci series
    c := make(chan int)

    // generate the Fibonacci series and write it to the channel
    go fibonacci(10, c)

    // read the Fibonacci series from the channel and print it
    for i := range c {
        fmt.Println(i)
    }
}
