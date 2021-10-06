package main

import (
	"fmt"
)

type Bitcoin int

func plusOne(x int) int {
	return x + 1
}

func bitcoinPlusOne(x Bitcoin) Bitcoin {
	return x + 1
}

func main() {
	var b1 Bitcoin = 1
	var b2 Bitcoin = 5
	sum := b1 + b2
	// It is OK to use %d verb
	fmt.Printf("%T %d\n", sum, sum) // main.Bitcoin 6

	var i int = 5
	// invalid operation: mismatched types Bitcoin and int
	// x := b1 + i
	x := b1 + Bitcoin(i)
	fmt.Printf("%T %d\n", x, x) // main.Bitcoin 6

	// cannot use b2 (variable of type Bitcoin) as int value in argument to plusOne
	// y := plusOne(b2)
	y := plusOne(int(b2))
	fmt.Printf("%T %d\n", y, y) // int 6

	// Maybe a little like inheritance conceptually
	// primitive value can be casted automatically !?
	b3 := b1 + 2                  // It is OK!?
	fmt.Printf("%T %d\n", b3, b3) // main.Bitcoin 4

	b4 := bitcoinPlusOne(0)       // It is OK!?
	fmt.Printf("%T %d\n", b4, b4) // main.Bitcoin 1
}
