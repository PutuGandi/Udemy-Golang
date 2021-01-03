package main

import "fmt"
// menggunakan tanda kutip 2 hanya 1 line
var x = "menggunakan kutip 2"

// meggunakan tanda backticks bisa membuat string pada lebih dari 1 line
var y = `Menggunakan
backticks" `

func main() {
	//x := 90
	fmt.Print(x)
	fmt.Println(x)
	fmt.Printf("ini adalang string %s yang saya gunakan",y)


}
