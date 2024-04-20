package main

import "fmt"

func ActiveBits(n int) int {
	count := 0

	// Iterate over each bit of the integer
	for n != 0 {
		// Check if the least significant bit is 1
		if n&1 == 1 {
			count++
		}

		// Shift the number to the right by one bit
		n >>= 1
	}

	return count
}

func main() {
	fmt.Println(ActiveBits(7))
}
