package main

import "fmt"

func main() {
    fmt.Println("Halo ini Golang")
    func_pertama()
    for i:=1;i<15;i++ {
        if i%2 == 0 {
            fmt.Println("Genap",32,"sdsd") //ngga bisa pakai ''
        }
    }
}

func func_pertama() {
    fmt.Println("Halo ini function Golang")
}