package main

import (
	"fmt"
)

func main() {
	x := []int{2, 45, 67, 90}
	// i = index dan v = value
	for i, v := range x {
		fmt.Println(i, v)
	}
}
