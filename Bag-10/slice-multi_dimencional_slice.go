package main

import (
	"fmt"
)

func main() {
	x := []string{"Putu", "Gandi", "Wijaya"}
	fmt.Println(x)

	y := []string{"DevOps", "Engineer"}
	fmt.Println(y)

	z := [][]string{x, y}
	fmt.Println(z)
}
