package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Input angka = ")
	fmt.Scan(&number)

	fmt.Println("Berikut faktor dari", number, "yaitu ;")
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
		}
	}

}
