package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time

    for i := 0; i <= 10; i++ {
        fmt.Println(i)
        time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond) // wait between 0 and 250 ms
    }
}
