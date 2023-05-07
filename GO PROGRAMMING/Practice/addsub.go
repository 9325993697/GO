package main

import "fmt"

func addSubtract(a, b int) (int, int) {
    return a+b, a-b
}

func main() {
    sum, diff := addSubtract(10, 5)
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", diff)
}
