package main

import "fmt"

func Compact(ptr *[]string) int {
	slice := *ptr // Dereference the pointer to get the slice
	size := len(slice)
	compactIndex := 0

	// Iterate over the slice and compact it
	for i := 0; i < size; i++ {
		if slice[i] != "" { // Check if the element is non-zero
			slice[compactIndex] = slice[i]
			compactIndex++
		}
	}

	// Update the size of the slice after compacting
	*ptr = slice[:compactIndex]

	return compactIndex
}

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
