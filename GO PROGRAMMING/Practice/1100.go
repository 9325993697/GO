package main

import "fmt"

func main() {
	evenSum := 0
	oddSum := 0

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}

	fmt.Println("Sum of even numbers between 1 to 100:", evenSum)
	fmt.Println("Sum of odd numbers between 1 to 100:", oddSum)
}
