package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// a ptr which points to a "interface{} instance" which is actually pointing to address of nil(0x0)
	efaceptr2nil := (*interface{})(nil)
	// Type.Elem() gets the derefrences type
	efacetyp := reflect.TypeOf(efaceptr2nil).Elem()
	// New returns a Value representing a pointer to a new zero value for the specified type (here is eface). That is, the returned Value's Type is PtrTo(typ).
	efaceptrVal := reflect.New(efacetyp)
	//
	efaceVal := efaceptrVal.Elem()
	efaceVal.Set(reflect.ValueOf(Person{Name: "TATA", Age: 25}))
	// efaceVal.Set(reflect.ValueOf(25))
	any := efaceVal.Interface()
	fmt.Println(any)
}
