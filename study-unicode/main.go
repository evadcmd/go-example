package main

import "fmt"

func main() {
	fmt.Println("\u4e16\u754c")

	for i, r := range "世界" {
		fmt.Print(i)
		fmt.Printf(" %x\n", r) // rune's value is code point of Unicode32
	}
}
