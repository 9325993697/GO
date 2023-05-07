package main

import (
    "fmt"
)

// isPalindrome function checks whether the number is a palindrome or not
func isPalindrome(num int) bool {
    var reversed int = 0
    var remainder int

    // reverse the number
    temp := num
    for {
        remainder = temp % 10
        reversed = reversed*10 + remainder
        temp /= 10

        if temp == 0 {
            break
        }
    }

    // compare the original and reversed numbers
    if num == reversed {
        return true
    } else {
        return false
    }
}

func main() {
    var num int
    fmt.Println("Enter a number: ")
    fmt.Scanf("%d", &num)

    if isPalindrome(num) {
        fmt.Printf("%d is a palindrome number.\n", num)
    } else {
        fmt.Printf("%d is not a palindrome number.\n", num)
    }
}
