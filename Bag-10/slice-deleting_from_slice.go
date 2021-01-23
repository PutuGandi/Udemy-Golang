package main

import (
	"fmt"
)

func main() {
	x := []int{0, 11, 22, 33, 44}
	x = append(x, 55, 66, 77)

	y := []int{88, 99, 1010}
	x = append(x, y...)
	fmt.Println(x)
	x = append(x[:1], y[1:]...)
	fmt.Println(x)
}
