package main

import "fmt"

type F func(int) int

type Enf interface{}
type Enfptr *interface{}
type Inf interface {
	method(int) int
}

func main() {
	// assign nil to ptr
	var intptr *int
	intptr = nil
	fmt.Println(intptr) // <nil>

	// assign nil to func
	var funcptr F
	funcptr = nil
	fmt.Println(funcptr) // <nil>

	// assign nil to map
	m := make(map[int]int)
	m = nil
	fmt.Println(m) // map[]

	/* CANNOT assign nil to array */
	// cannot use nil (untyped nil value) as [3]int value in assignment
	// l := [3]int{1, 2, 3}
	// l = nil

	// assign nil to slice
	l := []int{1, 2, 3}
	l = nil
	fmt.Println(l) // []

	var enf Enf
	enf = nil
	fmt.Println(enf) // <nil>

	var inf Inf
	inf = nil
	fmt.Println(inf) // <nil>

	var infptr Enfptr
	infptr = nil
	fmt.Println(infptr) // <nil>

	// assign nil to chan
	ch := make(chan<- int)
	ch = nil
	fmt.Println(ch) // <nil>
}
