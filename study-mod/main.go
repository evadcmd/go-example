package main

import (
	"fmt"

	"github.com/evadcmd/internal"
	"github.com/evadcmd/lib"
	"github.com/evadcmd/lib/source"
)

func main() {
	p := lib.Person{
		Name:  "TATA",
		Phone: source.Phone{Num: "1234567"},
	}
	fmt.Println(p)
	a := lib.Animal{
		Name: "TANUKI",
		Age:  2,
	}
	fmt.Println(a)
	superman := internal.SuperMan{Person: p}
	superman.Fly()
}
