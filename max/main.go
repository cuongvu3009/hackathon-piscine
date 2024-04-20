package main

import "fmt"

func Max(a []int) int {
	// Check if the slice is empty
	if len(a) == 0 {
		return 0
	}

	// Initialize the maximum value as the first element of the slice
	max := a[0]

	// Iterate over the slice and update the maximum value if a larger value is found
	for _, num := range a {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
