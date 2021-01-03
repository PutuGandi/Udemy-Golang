package main

import (
	"fmt"
)

func main() {
	s := "Hello World"

	// convert to decimal number
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// convert to hexa
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U \n", s[i])
	}
	fmt.Println("")

	// convert to hexa
	for i, v := range s {
		fmt.Printf("at index %d we have hex %#x \n", i, v)
	}
}
