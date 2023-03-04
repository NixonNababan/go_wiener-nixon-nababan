package main

import "fmt"

// definisi struct Mobil
type Mobill struct {
	Type   string
	FuelIn float64
}

// method untuk menghitung perkiraan jarak yang bisa ditempuh
func (c *Mobill) EstimateDistance() float64 {
	fuelEfficiency := 1.5 // km per mill
	distance := fuelEfficiency * c.FuelIn
	return distance
}

func main() {
	// membuat objek mobil dengan tipe "sedan" dan bahan bakar 1000 milliliter
	Mobilku := Mobill{"sedan", 1000.0}

	// menghitung perkiraan jarak yang bisa ditempuh oleh mobil menggunakan method EstimateDistance()
	distance := Mobilku.EstimateDistance()

	// mencetak hasil perhitungan ke layar
	fmt.Println("Perkiraan jarak yang bisa ditempuh:", distance, "km")
}
