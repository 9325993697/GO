package main

import "fmt"

func multiplicationTable(num int) {
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}

func main() {
    var num int
    fmt.Print("Enter a number to print its multiplication table: ")
    fmt.Scanln(&num)
    multiplicationTable(num)
}
