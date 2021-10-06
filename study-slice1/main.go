package main

import (
	"fmt"
	"reflect"
)

func sum(ints ...int) int {
	fmt.Println(reflect.TypeOf(ints)) // []int
	var sum int
	for _, i := range ints {
		sum += i
	}
	return sum
}

func main() {
	// cannot use ([]int literal) (value of type []int) as int value in argument to sum
	// fmt.Println(sum([]int{1, 2, 3}))
	fmt.Println(sum(1, 3, 5))
}
