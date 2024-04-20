package main

import "fmt"

func Unmatch(a []int) int {
	// Create a map to store the count of each element
	counts := make(map[int]int)

	// Iterate over the slice and count occurrences of each element
	for _, num := range a {
		counts[num]++
	}

	// Iterate over the counts and find the element with an odd count
	for num, count := range counts {
		if count%2 != 0 {
			return num
		}
	}

	// If all elements have a pair, return -1
	return -1
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
