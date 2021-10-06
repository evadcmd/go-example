package main

import "fmt"

func main() {
	// using float as key is dangerous but OK !!
	// https://golang.org/ref/spec#Comparison_operators
	m := map[float64]string{
		1.1: "string1",
		1.2: "string2",
		1.3: "string3",
	}
	for k, v := range m {
		fmt.Printf("k: %f, v: %s\n", k, v)
	}
}
