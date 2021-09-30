package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	Name  string
	Brand string
	Age   int
	Speed int
}

func showReflect(any interface{}) {
	vc := reflect.ValueOf(any)
	fmt.Printf("reflect.Value: %v\n", vc)
	fmt.Printf("reflect.Kind: %v\n", vc.Kind())
	tc := vc.Type()
	fmt.Printf("reflect.Type: %v\n", tc)
	fmt.Println()
}

func main() {
	list := []interface{}{
		5,
		2.5,
		5 + 2i,
		[]int{6, 5, 5, 3, 6},
		[5]int{6, 5, 5, 3, 6},
		map[int]string{
			1: "one",
			2: "two",
			3: "three",
		},
		Car{
			Name:  "afx-0",
			Brand: "Sugo",
			Age:   10,
			Speed: 650,
		},
	}

	for _, obj := range list {
		showReflect(obj)
	}

	// reflect.Value: 5
	// reflect.Kind: int
	// reflect.Type: int

	// reflect.Value: 2.5
	// reflect.Kind: float64
	// reflect.Type: float64

	// reflect.Value: (5+2i)
	// reflect.Kind: complex128
	// reflect.Type: complex128

	// reflect.Value: [6 5 5 3 6]
	// reflect.Kind: slice
	// reflect.Type: []int

	// reflect.Value: [6 5 5 3 6]
	// reflect.Kind: array
	// reflect.Type: [5]int

	// reflect.Value: map[1:one 2:two 3:three]
	// reflect.Kind: map
	// reflect.Type: map[int]string

	// reflect.Value: {afx-0 Sugo 10 650}
	// reflect.Kind: struct
	// reflect.Type: main.Car
}
