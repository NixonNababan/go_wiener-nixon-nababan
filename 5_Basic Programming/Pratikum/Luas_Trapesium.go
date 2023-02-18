package main

import (
	"fmt"
)

func main() {
	var alasAtas, alasBawah, tinggi float64

	fmt.Print("Masukkan panjang a: ")
	fmt.Scanln(&alasAtas)

	fmt.Print("Masukkan panjang b: ")
	fmt.Scanln(&alasBawah)

	fmt.Print("Masukkan t: ")
	fmt.Scanln(&tinggi)

	luas := (alasAtas + alasBawah) * tinggi / 2

	fmt.Printf("L=: %.3f\n", luas)
}
