package main

// Package fmt digunakan untuk melakukan print ke console, sedangkan package math digunakan untuk melakukan operasi matematika.
import (
	"fmt"
	"math"
)

// Variabel jumlahD1 dan jumlahD2 diinisialisasi dengan nilai 0. Variabel matrix didefinisikan sebagai sebuah array dua dimensi dengan ukuran 3x3 dan diisi dengan beberapa nilai integer.
func main() {
	jumlahD1 := 0
	jumlahD2 := 0
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},

		// melakukan iterasi pada array matrix dengan menggunakan dua loop for. Pada setiap iterasi, nilai pada diagonal utama dihitung dan ditambahkan ke variabel jumlahD1, sedangkan nilai pada diagonal kedua dihitung dan ditambahkan ke variabel jumlahD2.
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				jumlahD1 += matrix[i][j]
			}
			if i+j == 2 {
				jumlahD2 += matrix[i][j]
			}
		}
	}
	fmt.Println(jumlahD1)
	fmt.Println(jumlahD2)
	fmt.Println("jumlah diagonal", int(math.Abs(float64(jumlahD1-jumlahD2))))
}
