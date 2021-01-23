package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person //type diatas
	first  string
	ltk    bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "Boboboy",
		ltk:   true,
	}
	p1 := person{
		first: "Putu",
		last:  "Gandi",
		age:   26,
	}

	p2 := person{
		first: "citra",
		last:  "arisuta",
		age:   25,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println()
	fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p1.first, p2.first)
}
