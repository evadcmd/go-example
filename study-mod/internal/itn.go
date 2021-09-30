package internal

import (
	"fmt"

	"github.com/evadcmd/lib"
)

type SuperMan struct {
	lib.Person
}

func (su *SuperMan) Fly() {
	fmt.Println("I can fly")
}
