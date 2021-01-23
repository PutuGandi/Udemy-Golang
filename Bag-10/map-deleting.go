package main

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"Gandi":   15000000,
		"Putu":    25000000,
		"Arisuta": 200000,
	}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println("--------------")
	delete(x, "Gandi")
	for i, v := range x {
		fmt.Println(i, v)
	}

}
