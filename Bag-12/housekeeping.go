package main

import (
	"fmt"
)

var x int

type person struct {
	first string
	last  string
}

type foo int // ini merupakan alias
var y foo

const bar = 42

func main() {
	y = bar
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))
	fmt.Printf("%T\n", bar)
	fmt.Println(y)
	fmt.Println(bar)
}
