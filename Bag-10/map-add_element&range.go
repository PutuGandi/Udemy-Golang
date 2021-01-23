package main

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"Gandi": 15000000,
		"Putu":  25000000,
	}

	fmt.Println(x["Gandi"])
	x["Gede"] = 123
	fmt.Println(x)

	y := map[string]int{
		"Wijaya":  115000000,
		"Arisuta": 225000000,
	}

	fmt.Println("------------------------------------------")
	for i, v := range y {
		x[i] = v
		fmt.Println(x)
	}
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println(x)
	fmt.Println(x["Wijaya"])
}
