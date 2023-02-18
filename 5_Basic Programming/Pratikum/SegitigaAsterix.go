package main

import (
	"fmt"
)

func main() {
	a := 5 // Jumlah baris yang akan mau dibuat
	for i := 0; i < a; i++ {
		for j := 0; j < a-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
