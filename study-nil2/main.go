package main

import "fmt"

func main() {
	// nil slices
	var s []struct{}

	fmt.Println(len(s)) // 0
	fmt.Println(cap(s)) // 0

	// iterates zero times
	for range s {
	}

	// panic: index out of range
	fmt.Println(s[1])
}
