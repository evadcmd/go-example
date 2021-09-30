package main

import "fmt"

type Animal interface {
	Eat(food string)
}

// Cat def: use ref

type Cat struct {
	Name string
	Age  int
}

func (c *Cat) String() string {
	return c.Name
}

// go:noinline
func (c *Cat) Eat(food string) {
	fmt.Printf("%s eat %s meow~\n", c, food)
}

// Dog def

type Dog struct {
	Name string
	Age  int
}

func (c Dog) String() string {
	return c.Name
}

func (c Dog) Eat(food string) {
	fmt.Printf("%s eat %s bow~\n", c, food)
}

func main() {
	// go pass value instead of ref to func,
	// so if we pass a value(copy) of struct to a func, golang could not get the origin ref from it.
	// But if we pass a ref to a func, golang could get the struct back.

	/* cannot use (Cat literal) (value of type Cat) as Animal value in variable declaration: missing method Eat (Eat has pointer receiver) */
	// var c Animal = Cat{Name: "ZUZU", Age: 2}

	var cat Animal = &Cat{Name: "RURU", Age: 2}
	cat.Eat("fruit")

	var dog1 Animal = Dog{Name: "TATA", Age: 0}
	dog1.Eat("meat")

	var dog2 Animal = &Dog{Name: "CHIBI", Age: 0}
	dog2.Eat("milk")

	var any interface{} = dog2
	fmt.Println(any)
}
