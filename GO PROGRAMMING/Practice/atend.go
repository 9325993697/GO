
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file for appending
	file, err := os.OpenFile("filename.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Ask the user for input to append
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to append: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Append the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Text appended successfully!")
}