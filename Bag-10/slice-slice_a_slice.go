package main

import (
	"fmt"
)

func main() {
	x := []int{0, 1, 2, 3, 4, 35, 46, 57, 68}
	// dimulai dari index 2 sampai 4
	fmt.Println(x[2:5])
	// dimulai dari index 1 sampai 5
	fmt.Println(x[1:6])

	for i, v := range x {
		fmt.Println(i, v)
	}

	// code dibawah harus <=7 kalau lebih akan error kalau kurang masih bisa
	for i := 0; i <= 8; i++ {
		fmt.Printf("ini index %v dengan value %v \n", i, x[i])
	}
}
