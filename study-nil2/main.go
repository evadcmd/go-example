package main

import "fmt"

func main() {
	/* a nil slice */
	var s []struct{}
	fmt.Println(s)        // []
	fmt.Println(s == nil) // true
	// show the len && cap of the slice
	fmt.Println(len(s)) // 0
	fmt.Println(cap(s)) // 0
	// iterates zero times
	for range s {
	}
	s = append(s, struct{}{}, struct{}{})
	fmt.Println(s) //[{}]

	// panic: runtime error: index out of range [1] with length 0
	// fmt.Println(s[1])

	/* a nil slice */
	var m map[int]float64
	fmt.Println(m)        // map[]
	fmt.Println(m == nil) // true
	// show the len of the map
	fmt.Println(len(m)) // 0
	// iterates zero times
	for k, v := range m {
		fmt.Printf("%v: %v", k, v)
	}
	// It is OK to read even the map is nil
	fmt.Println((m[5])) // 0

	m[5] = 5.2 // panic: assignment to entry in nil map
}
