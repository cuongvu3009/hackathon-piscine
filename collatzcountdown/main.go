package main

import "fmt"

func CollatzCountdown(start int) int {
	// Return -1 if start is 0 or negative
	if start <= 0 {
		return -1
	}

	steps := 0

	// Perform the Collatz sequence until reaching 1
	for start != 1 {
		if start%2 == 0 {
			start /= 2
		} else {
			start = 3*start + 1
		}
		steps++
	}

	return steps
}

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}
