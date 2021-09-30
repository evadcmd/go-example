package main

import (
	"fmt"
	"reflect"
)

type obj struct{}
type Int int

func main() {
	x := 5
	y := map[int]string{
		1: "one",
	}
	var z int
	var n reflect.Value
	vx := reflect.ValueOf(x)
	vy := reflect.ValueOf(y)
	vz := reflect.ValueOf(z)

	/* IsValid */
	// Value.IsValid() is supposed to report whether the reflect.Value itself is valid, not the value it wraps (if any).
	fmt.Println(vx.IsValid())
	fmt.Println(vy.IsValid())
	fmt.Println(vz.IsValid())
	fmt.Println(n.IsValid())
	fmt.Println(reflect.Value{}.IsValid())

	/* Set */

	/* Set value to unaddressed type */
	// panic: reflect: reflect.Value.SetInt using unaddressable value
	// vz.SetInt(5)

	vvz := reflect.ValueOf(&z)

	/* Set with reflect.Value */
	// cannot use 5 (untyped int constant) as reflect.Value value in argument to vvz.Elem().Set
	// vvz.Elem().Set(5)
	vvz.Elem().Set(reflect.ValueOf(2))
	fmt.Println(z)

	/* Set* with wrong type */
	// panic: reflect: call of reflect.Value.SetBool on int Value
	// vrz.Elem().SetBool(false)

	// func (reflect.Value).SetInt(x int64)
	vz = vvz.Elem()
	vz.SetInt(5) // note that input type is type64
	fmt.Println(z)
	fmt.Println(reflect.TypeOf(z)) // still int
	// reflect.Value -> specific typed value
	// func (reflect.Value).Int() int64
	zint64 := vz.Int()
	tz := reflect.TypeOf(zint64)
	fmt.Println(tz)        // int64!!
	fmt.Println(tz.Kind()) // int64!!
	// reflect.Value -> interface{}
	iz := vz.Interface()
	fmt.Println(iz) // 5
}
