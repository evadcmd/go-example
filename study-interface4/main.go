package main

import "fmt"

type HasName interface {
	Name() string
}

func sayHello(hasName HasName) {
	fmt.Printf("I am %s, Hello\n", hasName.Name())
}

type TATA struct {
	N string
}

func (t *TATA) Name() string {
	// To deal with condition which t is nil
	// if t == nil {...}
	return t.N
}

func main() {
	t := TATA{N: "tata"}
	sayHello(&t)
	// it is OK to send nil as an interface input
	// panic: runtime error: invalid memory address or nil pointer dereference
	sayHello(nil)
}
