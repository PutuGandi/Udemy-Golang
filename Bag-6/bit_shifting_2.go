package main

import "fmt"

// func main() {
// 	kb := 1024
// 	mb := 1024 * kb
// 	gb := 1024 * mb

// 	fmt.Printf("%d\t\t%b\n", kb, kb)
// 	fmt.Printf("%d\t\t%b\n", mb, mb)
// 	fmt.Printf("%d\t%b\n", gb, gb)
// }

//menggunakan iota dan konsep bit shifting
const (
	_  = iota // dimulai dengan iota ialah 0
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
)

func main() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
