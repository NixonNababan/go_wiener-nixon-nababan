package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	// inisialisasi nilai awal min dan max
	min = *numbers[0]
	max = *numbers[0]

	// loop melalui array, membandingkan nilai setiap elemen dengan min dan max
	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}

	return
}
func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	// pass alamat variabel ke fungsi getMinMax menggunakan operator &
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Printf("%d is min\n", min)
	fmt.Printf("%d is max\n", max)
}
