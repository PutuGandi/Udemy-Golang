package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Putu",
		last:  "Gandi",
		age:   26,
	}

	p2 := person{
		first: "Putu",
		last:  "Arosuta",
		age:   25,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p2.first)
}
