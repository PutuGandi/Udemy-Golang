package main

import (
	"fmt"
)

func main() {
	x := []int{0, 11, 22, 33, 44}
	x = append(x, 55, 66, 77)
	fmt.Printf("%v \n\n", x)

	y := []int{88, 99, 1010}
	x = append(x, y...)
	fmt.Println(x)
}
