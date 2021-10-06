package main

import (
	"fmt"
	"reflect"
)

func sum(ints []int) int {
	fmt.Println(reflect.TypeOf(ints)) // []int
	var sum int
	for _, i := range ints {
		sum += i
	}
	return sum
}

func main() {
	arr := [...]int{2, 4, 6}

	slice := arr[:] // from array to slice
	fmt.Println(sum(slice))

	// cannot convert slice (variable of type []int) to [3]int
	// arrp := ([3]int)(slice)

	arrp := (*[3]int)(slice) // from slice to array pointer
	for i, v := range arrp {
		fmt.Printf("index: %d  value: %v\n", i, v)
	}
	for i, v := range *arrp {
		fmt.Printf("index: %d  value: %v\n", i, v)
	}
}
