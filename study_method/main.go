package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func (c *Cat) AgePlusOne() {
	c.Age++
}

/* method AgePlusOne already declared for type Cat struct{Name string; Age int} */
// func (c Cat) AgePlusOne() {
//     c.Age++
// }

// go has no static method!
func (c Cat) Of(name string, age int) *Cat {
	return &Cat{
		Name: name,
		Age:  age,
	}
}

func main() {
	/* not static method */
	// not enough arguments in call to Cat.Of
	// c := Cat.Of("TATA", 3)

	// agly but only for testing
	c := Cat{}.Of("TATA", 3)
	c.AgePlusOne()
	fmt.Println(c)
}
