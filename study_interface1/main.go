package main

import "fmt"

type Test struct {
	Content string
}

func isNil(t interface{}) bool {
	return t == nil
}

func main() {
	var t Test
	fmt.Println(t) // {}
	// cannot convert nil (untyped nil value) to struct{Content string}
	// fmt.Println(t == nil)
	fmt.Println(isNil(t)) // false

	t.Content = "hello world"
	fmt.Println(t) // {hello world}

	var ptr *Test
	fmt.Println(ptr) // <nil>

	fmt.Println(ptr == nil) // true
	fmt.Println(isNil(ptr)) // false!!

	// cannot compare nil == nil (operator == not defined for untyped nil)
	// fmt.Println(nil == nil)
	fmt.Println(isNil(nil)) // true!!
}
