package lib

import "github.com/evadcmd/lib/source"

type Person struct {
	Name string
	Age  int
	source.Phone
}
