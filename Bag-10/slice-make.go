package main

import (
	"fmt"
)

func main() {
	// make(type, length, capacity)
	x := make([]int, 5, 7)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 11
	x[4] = 55

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//menambahkan value
	x = append(x, 666)
	fmt.Println("----------------------")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 777)
	fmt.Println("---------------------1")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 88)
	fmt.Println("---------------------2")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
