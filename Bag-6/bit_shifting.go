package main

import "fmt"

func main() {
	x := 4
	y := 120
	fmt.Printf("Binary dari %d\t\t\t%b\n", x, x)
	fmt.Printf("Binary dari %d\t\t\t%b\n", y, y)

	fmt.Println("Melakukan bit shifting pada desimal 4 dengan binary 100")
	a := x << 1 //geser 1 digit
	fmt.Printf("%d\t\t\t%b\n", a, a)
	b := x >> 1 //geser 1 digit
	fmt.Printf("%d\t\t\t%b\n", b, b)
	c := x << 10 //geser 10 digit
	fmt.Printf("%d\t\t\t%b\n", c, c)
	d := x >> 10 //geser 10 digit
	fmt.Printf("%d\t\t\t%b\n", d, d)
}
