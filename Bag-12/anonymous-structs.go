package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Putu",
		last:  "Gandi",
		age:   26,
	}

	fmt.Println(p1)

	fmt.Println(p1.first)
}
