package main

import (
	"fmt"
)

func main() {
	var nilai int
	fmt.Print("Masukkan angka =")
	fmt.Scanln(&nilai)

	if nilai%2 == 0 {
		fmt.Printf("Bilangan %d adalah genap  ", nilai)
	} else {
		fmt.Printf("Bilangan %d adalah ganjil ", nilai)
	}
}
