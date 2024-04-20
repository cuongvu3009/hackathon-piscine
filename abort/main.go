package main

import (
	"fmt"
	"sort"
)

func Abort(a, b, c, d, e int) int {
	// Create a slice of integers from the arguments
	numbers := []int{a, b, c, d, e}

	// Sort the slice
	sort.Ints(numbers)

	// Calculate the middle index
	middleIndex := len(numbers) / 2

	// Check if the length of the slice is odd or even
	if len(numbers)%2 == 0 {
		// If the length is even, return the average of the two middle values
		return (numbers[middleIndex-1] + numbers[middleIndex]) / 2
	} else {
		// If the length is odd, return the middle value
		return numbers[middleIndex]
	}
}

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
