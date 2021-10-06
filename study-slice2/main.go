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
	fmt.Println(sum(arr[:]))
}
