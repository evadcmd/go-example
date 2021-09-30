package main

import (
	"fmt"
)

type X struct {
	Val int
}

type SomeItf interface {
	DoSomething() string
}

func main() {
	itf0 := (interface{})(nil)
	fmt.Println(itf0 == nil) // true
	itf1 := (interface{})(1)
	fmt.Println(itf1 == nil) // false
	itf2 := SomeItf(nil)
	fmt.Println(itf2 == nil) // true!!
	var x X
	itfX := (interface{})(x)
	fmt.Println(itfX == nil) // false!!

	itfptr2nil := (*interface{})(nil)
	fmt.Println(itfptr2nil == nil) // true!! a [interface{}] ptr pointing to 0x00
	fmt.Println(itfptr2nil)        // <nil>
	itfItfptr2nil := (interface{})(itfptr2nil)
	fmt.Println(itfItfptr2nil == nil) // false

	intptr2nil := (*int)(nil)
	fmt.Println(intptr2nil == nil) // true!! a [int] ptr pointing to 0x00
	fmt.Println(intptr2nil)        // <nil>
}
