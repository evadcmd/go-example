package main

import "fmt"

type F func(int) int

type Enf interface{}
type Enfptr *interface{}
type Inf interface {
	method(int) int
}

type SomeStruct struct {
	X        int
	Y        int
	Descript string
}

func main() {
	/* CANNOT assign nil to struct */
	// var s SomeStruct
	// s = nil

	// assign nil to ptr
	var intptr *int
	fmt.Println(intptr == nil) // true
	intptr = nil
	fmt.Println(intptr) // <nil>

	/* * * * * * * * * * * * *  */
	/* {_type: *int, data: nil} */
	/* * * * * * * * * * * * *  */
	wrapped := interface{}(intptr)
	fmt.Println(wrapped)        // <nil>
	fmt.Println(wrapped == nil) // false !!

	// assign nil to func
	var funcptr F
	fmt.Println(funcptr) // <nil>
	funcptr = nil
	fmt.Println(funcptr) // <nil>

	// assign nil to map
	m := make(map[int]int)
	fmt.Println(m == nil) // false
	m = nil
	fmt.Println(m)        // map[]
	fmt.Println(m == nil) // true
	// panic: assignment to entry in nil map
	// m[1] = 5

	/* CANNOT assign nil to array */
	// cannot use nil (untyped nil value) as [3]int value in assignment
	// l := [3]int{1, 2, 3}
	// l = nil

	// assign nil to slice
	l := []int{}
	fmt.Println(l == nil) // false
	l = nil
	fmt.Println(l)        // []
	fmt.Println(l == nil) // true

	// assing nil to data part of interface. ex: {_type: nil, data: nil <-assigned}
	var enf Enf
	enf = nil
	fmt.Println(enf) // <nil>

	/* note: error is a interface so we can compare it with nil */
	/* type error interface { Error() string } */
	var inf Inf
	inf = nil
	fmt.Println(inf) // <nil>

	var infptr Enfptr
	infptr = nil
	fmt.Println(infptr)             // <nil>
	fmt.Println(infptr == nil)      // true
	fmt.Println(Enfptr(nil) == nil) // true

	// assign nil to chan
	ch := make(chan<- int)
	ch = nil
	fmt.Println(ch) // <nil>

	iptr := (*int)(nil)
	fmt.Println(iptr) // <nil>
}
