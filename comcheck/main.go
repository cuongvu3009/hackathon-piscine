package main

import (
	"fmt"
	"os"
)

func main() {
	// Iterate over command-line arguments starting from index 1 (index 0 is the program name)
	for _, arg := range os.Args[1:] {
		// Check if the argument matches any of the specified strings
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return // Exit after printing "Alert!!!" once
		}
	}
}
