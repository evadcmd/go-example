package main

import "fmt"

// get 2's complement of int32
func comp2(i int32) string {
	return fmt.Sprintf("%b", uint32(i))
}

func descrip(i int32) {
	u := uint32(i)
	fmt.Printf("int32 value: %d\n", i)
	fmt.Printf("2'complement: %b\n", u)
	fmt.Printf("uint32 value:  %d\n", u)
	fmt.Printf("negative uint32 value:  %d\n\n", -u) // get 2'complement of u
}

func main() {
	fmt.Println("1 << 2")
	descrip(1 << 2)

	fmt.Println("-1 << 2")
	descrip(-1 << 2)

	fmt.Println("-1 << 31")
	descrip(-1 << 31)
}
