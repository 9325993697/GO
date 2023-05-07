package main

import (
	"fmt"
	"strconv"
)

func digitSquareCubeSum(num int, sqChan, cubeChan chan int) {
	squaresSum := 0
	cubesSum := 0
	digits := strconv.Itoa(num)

	for _, digit := range digits {
		digitInt, _ := strconv.Atoi(string(digit))
		squaresSum += digitInt * digitInt
		cubesSum += digitInt * digitInt * digitInt
	}

	sqChan <- squaresSum
	cubeChan <- cubesSum
}

func main() {
	num := 123

	sqChan := make(chan int)
	cubeChan := make(chan int)

	go digitSquareCubeSum(num, sqChan, cubeChan)

	squaresSum := <-sqChan
	cubesSum := <-cubeChan

	fmt.Printf("Number: %d\n", num)
	fmt.Printf("Sum of squares: %d\n", squaresSum)
	fmt.Printf("Sum of cubes: %d\n", cubesSum)
}
