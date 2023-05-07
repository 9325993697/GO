

package main

import (
	"fmt"
	"strconv"

)

func main() {
	var num int
	fmt.Println("Enter any number")
	fmt.Scanln(&num)
	if len(strconv.Itoa(num)) == 1 {
		fmt.Println("The number is a single digit number")
	}else{
		fmt.Println("The Number is not single digit number")
	}
}