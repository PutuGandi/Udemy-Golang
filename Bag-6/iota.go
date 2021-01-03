package main

import "fmt"

// penggunaannya cuma bisa dengan const
const (
	a = iota + 3
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
