package main

import (
	"fmt"
	"strings"
)

func Join(strs []string, sep string) string {
	// Use strings.Join to concatenate the strings with the separator
	return strings.Join(strs, sep)
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
