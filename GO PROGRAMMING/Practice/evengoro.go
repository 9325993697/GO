package main

import "fmt"

func even(c chan int, nums []int) {
    for _, num := range nums {
        if num%2 == 0 {
            c <- num
        }
    }
    close(c)
}

func odd(c chan int, nums []int) {
    for _, num := range nums {
        if num%2 != 0 {
            c <- num
        }
    }
    close(c)
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evenC := make(chan int)
    oddC := make(chan int)

    go even(evenC, nums)
    go odd(oddC, nums)

    fmt.Println("Even Numbers:")
    for num := range evenC {
        fmt.Println(num)
    }

    fmt.Println("\nOdd Numbers:")
    for num := range oddC {
        fmt.Println(num)
    }
}
