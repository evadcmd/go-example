package main

import (
	"fmt"
)

func main() {
	n1 := (interface{})(nil)
	if n1 == nil {
		fmt.Println("null1")
	}
	nilptr := (*interface{})(nil)
	if nilptr == nil {
		fmt.Println("null2")
	}
	intv := (interface{})(1)
	if intv == nil {
		fmt.Println("intv")
	}
}
