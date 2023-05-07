package main

import "fmt"

// function that returns multiple values
func divideAndRemainder(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}

func main() {
    // call the function and receive the multiple return values
    quotient, remainder := divideAndRemainder(10, 3)

    fmt.Printf("Quotient: %d, Remainder: %d", quotient, remainder)
}
