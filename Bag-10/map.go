package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Gandi": 25,
		"Citra": 24,
	}
	fmt.Println(m["Gandi"])
	// key yang dipanggil tidak ada di map maka akan diberi value 0
	fmt.Println(m["Putu"])

	v, ok := m["Gandi"]
	fmt.Println(v)
	fmt.Println(ok)

	p, tok := m["Putu"]
	fmt.Println(p)
	fmt.Println(tok)

	if vv, ook := m["Citra"]; ook {
		fmt.Println("This is the if print", vv)
	}
}
