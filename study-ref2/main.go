package main

import "fmt"

type X struct {
	A int
	B string
	C float64
}

func main() {
	var x1 interface{} = &X{}
	var x2 interface{}
	if x2, ok := x1.(*X); ok {
		x2.A = 25
		x2.B = "S"
		x2.C = 5.2
	}
	fmt.Println(x1) // &{25 S 5.2}
	fmt.Println(x2) // <nil> !!?
}
