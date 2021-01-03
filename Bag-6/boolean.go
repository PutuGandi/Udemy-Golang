package main

import "fmt"

// mendeklarasikan variable x ialah bertipe boolean
var x bool

func main() {
	fmt.Println(x)
	// memasukkan nilai x menjadi True
	x = true
	fmt.Println(x)
	// mendeklarisikan y = true dengan short delcaration (:=)
	// kalau hanya (=) maka error, karena belum di deklarasikan
	y := true
	// menggunakan operator (==) untuk membandingkan
	fmt.Println(x == y)
}