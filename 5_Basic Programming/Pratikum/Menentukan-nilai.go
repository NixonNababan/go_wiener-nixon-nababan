package main

// Wiener Nixon Nababan
import (
	"fmt"
)

func main() {
	var value float64
	fmt.Print("masukkan nilai = ")
	fmt.Scanln(&value)

	grade := "Nilai Invalid"
	if value >= 0 && value <= 100 {
		switch {
		case value >= 80:
			grade = "A"
		case value >= 65:
			grade = "B"
		case value >= 50:
			grade = "C"
		case value >= 35:
			grade = "D"
		default:
			grade = "E"
		}
	}

	fmt.Printf("Nilai: %s\n", grade)
}
