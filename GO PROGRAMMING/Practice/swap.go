package main

import "fmt"

func swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    var x, y int

    fmt.Print("Enter two numbers: ")
    fmt.Scan(&x, &y)

    fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)

    swap(&x, &y)

    fmt.Printf("After swapping, x = %d and y = %d\n", x, y)
}
