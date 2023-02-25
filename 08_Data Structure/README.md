# (8) Data Structure
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

## Prioritas 1

### Soal 1
```Go
package main

import "fmt"

// Fungsi ArrayMerge menerima dua parameter slice, arrayA dan arrayB, dan kemudian membuat slice baru (mergedArray) dengan ukuran arrayA + arrayB menggunakan fungsi make. Kemudian, elemen arrayA disalin ke mergedArray menggunakan fungsi copy.
func ArrayMerge(arrayA, arrayB []string) []string {
	mergedArray := make([]string, len(arrayA)+len(arrayB))
	copy(mergedArray, arrayA)

	// melakukan loop pada setiap elemen arrayB dan memeriksa apakah elemen tersebut sudah ada di dalam arrayA. Jika elemen tidak ditemukan dalam arrayA, maka elemen tersebut ditambahkan ke mergedArray menggunakan fungsi append.
	for _, value := range arrayB {
		found := false
		for _, element := range arrayA {
			if element == value {
				found = true
				break
			}
		}
		if !found {
			mergedArray = append(mergedArray, value)
		}
	}

	return mergedArray
}

func main() {

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

}
```

### Soal 2
```Go
package main

import "fmt"

//  setiap string pada slice akan diloop dan dicocokkan dengan setiap string pada slice. Jika string yang sedang dicari ditemukan, maka variabel count akan diincrement
func Mapping(slice []string) map[string]int {
	result := make(map[string]int)
	for _, s := range slice {
		count := 0
		for _, r := range slice {
			if s == r {
				count++
			}
		}
		// Setelah selesai loop kedua, variabel count akan menjadi jumlah kemunculan string tersebut pada slice, dan dimasukkan ke dalam map yang akan dikembalikan oleh fungsi Mapping.
		result[s] = count
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))

	fmt.Println(Mapping([]string{}))
}
```

### Soal 3
```Go
package main

import (
	"fmt"
	"strconv"
)

// Pada setiap perulangan, kita menghitung berapa kali karakter pada indeks i muncul pada string input angka.
func munculSekali(angka string) []int {
	var result []int
	for i := 0; i < len(angka); i++ {
		count := 0
		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] {
				count++
			}
		}

		//angka akan di-convert dari string menjadi integer menggunakan fungsi Atoi dari package strconv, lalu ditambahkan ke dalam slice result menggunakan fungsi append.
		if count == 1 {
			num, _ := strconv.Atoi(string(angka[i]))
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // Output :[4]
	fmt.Println(munculSekali("76523752"))   // Output :[6 3]
	fmt.Println(munculSekali("12345"))      // Output :[1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // Output :[]
	fmt.Println(munculSekali("0872504"))    // Output :[8 7 2 5 4]
}
```

## Soal Prioritas 2

### Soal 1
```Go
package main

import (
	"fmt"
	"strconv"
)

// Pada setiap perulangan, kita menghitung berapa kali karakter pada indeks i muncul pada string input angka.
func munculSekali(angka string) []int {
	var result []int
	for i := 0; i < len(angka); i++ {
		count := 0
		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] {
				count++
			}
		}

		//angka akan di-convert dari string menjadi integer menggunakan fungsi Atoi dari package strconv, lalu ditambahkan ke dalam slice result menggunakan fungsi append.
		if count == 1 {
			num, _ := strconv.Atoi(string(angka[i]))
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // Output :[4]
	fmt.Println(munculSekali("76523752"))   // Output :[6 3]
	fmt.Println(munculSekali("12345"))      // Output :[1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // Output :[]
	fmt.Println(munculSekali("0872504"))    // Output :[8 7 2 5 4]
}
```

## Soal Explorasi
### Soal 1
```Go
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
```