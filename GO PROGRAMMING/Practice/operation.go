package main

import "fmt"

func main() {
    // Initializing a slice with some elements
    slice1 := []int{1, 2, 3, 4, 5}
    fmt.Println("Original Slice:", slice1)

    // Appending an element to the slice
    slice1 = append(slice1, 6)
    fmt.Println("Slice After Appending an Element:", slice1)

    // Removing an element from the slice
    slice1 = remove(slice1, 3)
    fmt.Println("Slice After Removing an Element:", slice1)

    // Copying a slice to another slice
    slice2 := make([]int, len(slice1))
    copy(slice2, slice1)
    fmt.Println("Copied Slice:", slice2)
}

// Function to remove an element from the slice
func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}
