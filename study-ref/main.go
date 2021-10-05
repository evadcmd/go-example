package main

import (
	"fmt"
	"strconv"
	"strings"
)

type X struct {
	A int
	B string
	C float64
}

func sum(i ...int) int {
	sum := 0
	for _, x := range i {
		sum += x
	}
	return sum
}

func concat(i ...int) string {
	var sb strings.Builder
	for _, x := range i {
		sb.WriteString(strconv.Itoa(x))
	}
	return sb.String()
}

func main() {
	x1 := X{
		A: 5,
		B: "B",
		C: 5.2,
	}
	x2 := x1 // x2 is a copy of x1
	x1.A = 10
	fmt.Println(x1) // {10 B 5.2}
	fmt.Println(x2) // {5 B 5.2}

	i1 := 7
	i2 := i1
	i1 = 17
	fmt.Println(i1) // 17
	fmt.Println(i2) // 7

	var itf1 interface{} = X{}
	var itf2 interface{} = itf1
	if impl2, ok := itf2.(X); ok {
		impl2.A = 10
		impl2.B = "C"
		fmt.Println(impl2) // {10 C 0}
	}
	fmt.Println(itf1) // {0  0}
	fmt.Println(itf2) // {0  0}

	f1 := sum
	f2 := f1
	// cannot use concat (value of type func(i ...int) string) as func(i ...int) int value in assignment
	// f1 = concat
	fmt.Println(f2(1, 2, 3, 4, 5)) // sum = 15

	// var x func(...int) int
	x := (func(...int) int)(nil)
	x = sum
	x(1)

	var fun1 interface{} = sum
	if f1, ok := fun1.(func(...int) int); ok {
		fmt.Println(f1(1, 3, 5, 7, 9)) // 25
	}

	var fun2 interface{} = fun1
	if f2, ok := fun2.(func(...int) int); ok {
		fmt.Println(f2(2, 4, 6, 8, 10)) // 30
	}
	fun2 = concat
	if f2, ok := fun2.(func(...int) string); ok {
		fmt.Println(f2(2, 4, 6, 8, 10)) // 246810
	}
	if f1, ok := fun1.(func(...int) string); ok {
		fmt.Println(f1(2, 4, 6, 8, 10))
	}
}
